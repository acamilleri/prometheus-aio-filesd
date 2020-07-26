package writer

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kelseyhightower/envconfig"
)

func init() {
	Register("file", NewFile)
}

// File ...
type File struct {
	Dest string `envconfig:"FILESD_WRITER_FILE_DEST" required:"true"`
}

// NewFile File writer
func NewFile() (Writer, error) {
	var fw File

	err := envconfig.Process(fmt.Sprintf("%s_%s", DefaultEnvVarsPrefix, "file"), &fw)
	if err != nil {
		return nil, err
	}
	return &fw, nil
}

// Write ...
func (fw *File) Write(data []byte) error {
	return ioutil.WriteFile(fw.Dest, data, os.ModePerm)
}
