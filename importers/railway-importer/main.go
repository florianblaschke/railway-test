package main

import (
	"github.com/forsuxess/railway-test/jobs"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	job := jobs.NewJob()
	job.ParseTitleFromString("  Senior Software Engineer (m/f/d)  ")

	logrus.WithFields(logrus.Fields{
		"client":  "railway_test_client",
		"run_id":  1,
		"job_id":  301,
		"title":   job.Title,
		"success": true,
	}).Info("railway importer completed")
}
