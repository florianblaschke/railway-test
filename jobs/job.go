package jobs

import "strings"

type Job struct {
	Title string
}

func (j *Job) ParseTitleFromString(s string) {
	title := strings.TrimSpace(s)
	title = strings.ToLower(title)
	title = strings.ReplaceAll(title, "  ", " ")
	j.Title = title
}

func NewJob() *Job {
	return &Job{}
}
