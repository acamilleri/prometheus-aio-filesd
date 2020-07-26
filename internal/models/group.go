package models

// Group Object
type Group struct {
	Targets []Target `json:"targets"`
	Labels  Label    `json:"labels"`
	Source  string   `json:"-"`
}
