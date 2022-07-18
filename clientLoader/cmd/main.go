package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

type Port struct {
	Name        string        `json:"name"`
	City        string        `json:"city"`
	Country     string        `json:"country"`
	Alias       []interface{} `json:"alias"`
	Regions     []interface{} `json:"regions"`
	Coordinates []float64     `json:"coordinates"`
	Province    string        `json:"province,omitempty"`
	Timezone    string        `json:"timezone,omitempty"`
	Unlocs      []string      `json:"unlocs"`
	Code        string        `json:"code"`
}

type Loader struct {
	decoder json.Decoder
}

func NewLoader(r io.Reader) *Loader {
	return &Loader{
		decoder: *json.NewDecoder(r),
	}
}

// ParsePort decodes input Port JSON
func (l *Loader) ParsePort() {
}

// SendPorts insert or update Port in Database
func (l *Loader) SendPorts() {
	// Loop and Parse port before sending
}

func main() {
	logCtx := logrus.WithFields(
		logrus.Fields{"component": "cmd", "function": "main"},
	)
	logCtx.Info("Starting Client Loader")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	file, err := os.Open("./clientLoader/input/ports.json")
	if err != nil {
		logCtx.Error(err)
		os.Exit(1)
	}

	loader := NewLoader(file)
	go loader.SendPorts()

	<-c
	fmt.Println("Loader Finishing")
}
