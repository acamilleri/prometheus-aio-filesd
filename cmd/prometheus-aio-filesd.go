package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/acamilleri/prometheus-aio-filesd/internal/adapter"
	"github.com/acamilleri/prometheus-aio-filesd/internal/formatter"
	_ "github.com/acamilleri/prometheus-aio-filesd/internal/provider"
	provider "github.com/acamilleri/prometheus-aio-filesd/internal/provider/core"
	"github.com/acamilleri/prometheus-aio-filesd/internal/writer"
	_ "github.com/acamilleri/prometheus-aio-filesd/internal/writer"
)

var (
	version string

	providersWant = kingpin.Flag("provider", "target(s) source").
		Envar("FILESD_PROVIDER_NAME").
		Required().
		Enums(provider.ListAvailableProviders()...)

	writerWant = kingpin.Flag("writer", "output destination (options: file, stdout)").
		Envar("FILESD_WRITER_NAME").
		Required().
		Enum(writer.ListAvailableWriters()...)

	formatterWant = kingpin.Flag("formatter", "output format (options: json, yaml)").
		Envar("FILESD_FORMATTER_NAME").
		Default("json").
		Enum(formatter.ListFormatterAvailable()...)

	refreshInterval = kingpin.Flag("refresh.interval", "refresh targets every (time.Duration)").
		Default(time.Duration(time.Second * 10).String()).
		Envar("FILESD_REFRESH_INTERVAL").
		Duration()

	logLevel = kingpin.Flag("log.level", "log level").
		Envar("FILESD_LOG_LEVEL").
		Default(logrus.InfoLevel.String()).
		Enum("debug", "info", "error")

	logFormat = kingpin.Flag("log.format", "log format").
		Envar("FILESD_LOG_FORMAT").
		Default("text").
		Enum("text", "json")
)

func main() {
	kingpin.Version(version)
	kingpin.Parse()

	err := setupLogrus(*logLevel, *logFormat)
	if err != nil {
		kingpin.Fatalf("logger: %v", err)
	}

	wr, err := writer.New(*writerWant)
	if err != nil {
		kingpin.Fatalf("writer: %v", err)
	}

	ftr, err := formatter.New(*formatterWant)
	if err != nil {
		kingpin.Fatalf("formatter: %v", err)
	}

	ticker := time.NewTicker(*refreshInterval)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logrus.Info("SIGTERM detected stopping ...")
		ticker.Stop()
		os.Exit(0)
	}()

	runFn := func() {
		for _, providerWant := range *providersWant {
			p, err := provider.New(providerWant)
			if err != nil {
				logrus.Errorf("provider: %v, skipping ...", err)
				continue
			}

			a, err := adapter.New(p, wr, ftr)
			if err != nil {
				logrus.Errorf("adapter: %v, skipping ...", err)
				continue
			}

			err = a.Run()
			if err != nil {
				logrus.Errorf("adapter: %v, skipping ...", err)
			}
		}
	}

	// run the first time before ticker
	runFn()
	for {
		select {
		case <-ticker.C:
			runFn()
		}
	}
}

func setupLogrus(level, format string) error {
	err := defineLevel(level)
	if err != nil {
		return err
	}

	defineFormatter(format)
	return nil
}

func defineLevel(level string) error {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}

	logrus.SetLevel(lvl)
	return nil
}

func defineFormatter(name string) {
	var ftr logrus.Formatter

	switch name {
	case "json":
		ftr = &logrus.JSONFormatter{}
	case "text":
		ftr = &logrus.TextFormatter{}
	default:
		logrus.Warnf("no formatter with name %s found, fallback to text", name)
		ftr = &logrus.TextFormatter{}
	}

	logrus.SetFormatter(ftr)
}
