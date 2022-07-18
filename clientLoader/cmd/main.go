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
	decoder *json.Decoder
}

func NewLoader(r io.Reader) *Loader {
	return &Loader{
		decoder: json.NewDecoder(r),
	}
}

// parsePort decodes Port JSON
func (l *Loader) parsePort() (*Port, error) {
	dec := l.decoder

	if dec.More() {
		// Ignore first key
		t, err := dec.Token()
		if err != nil {
			return nil, err
		}
		fmt.Printf("2 %T: %v\n", t, t)

		if dec.More() {
			var p Port

			// decode Port
			err := dec.Decode(&p)
			if err != nil {
				return nil, err
			}
			p.Code = p.Unlocs[0]

			//fmt.Printf("%+v\n", p)
			return &p, nil
		}
	}

	return nil, nil
}

// ParseAndSendPorts inserts or updates Port in Database
func (l *Loader) ParseAndSendPorts() error {
	logCtx := logrus.WithFields(
		logrus.Fields{"component": "cmd", "function": "ParseAndSendPorts"},
	)
	dec := l.decoder

	// read open bracket
	_, err := dec.Token()
	if err != nil {
		return err
	}

	// Loop and Parse port before sending
	for dec.More() {
		port, err := l.parsePort()
		if err != nil {
			return err
		}
		// sendPortViaGRPC(port)
		logCtx.Infof("Sending Port: %v", port)
	}

	// read close bracket
	_, err = dec.Token()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	logCtx := logrus.WithFields(
		logrus.Fields{"component": "cmd", "function": "main"},
	)
	logCtx.Info("*** Starting Client Loader ***")

	file, err := os.Open("./clientLoader/input/ports.json")
	if err != nil {
		logCtx.Error(err)
		os.Exit(1)
	}

	loader := NewLoader(file)

	stopc := make(chan os.Signal, 1)
	signal.Notify(stopc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	errc := make(chan error, 1)

	go func() {
		errc <- loader.ParseAndSendPorts()
	}()

	select {
	case <-stopc:
		goto gracefulShutdown
	case err = <-errc:
		if err != nil {
			logCtx.Error(err)
			os.Exit(1)
		}
	}

gracefulShutdown:
	fmt.Println("*** Finishing Loader ***")
}
