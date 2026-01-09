package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/forsuxess/railway-test/jobs"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

type RickAndMorty struct {
	Characters string `json:"characters"`
	Locations  string `json:"locations"`
	Episodes   string `json:"episodes"`
}

func main() {
	job := jobs.NewJob()
	job.ParseTitleFromString("  Junior Software Engineer (m/f/d)  ")

	logrus.WithFields(logrus.Fields{
		"client":  "railway_test_client_two",
		"run_id":  1,
		"job_id":  301,
		"title":   job.Title,
		"success": true,
	}).Info("railway importer completed")
	client := http.Client(http.Client{Timeout: time.Second * 5})
	request, err := http.NewRequest("GET", "https://rickandmortyapi.com/api", nil)
	if err != nil {
		fmt.Printf("Could not make request")
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Failed executing request")
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("Failed parsing body")
	}
	rick := RickAndMorty{}
	err = json.Unmarshal(body, &rick)
	if err != nil {
		fmt.Printf("failure unmarshaling body")
	}

	logrus.WithFields(logrus.Fields{
		"client":     "ricky-morty-call",
		"run_id":     1,
		"job_id":     200,
		"characters": rick.Characters,
		"locations":  rick.Locations,
		"episodes":   rick.Episodes,
		"success":    true,
	}).Info("we have morty")

}
