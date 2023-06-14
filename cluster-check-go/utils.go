package main

import (
	"encoding/json"
	"fmt"
	"os"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Reset     = "\033[0m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Separator = "---------------------------------------------------------"
)

func loadConfig(configFile *string) *Config {
	var configCheck Config
	if *configFile != "" {
		data, err := os.ReadFile(*configFile)
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal(data, &configCheck)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		configCheck = configCheck.GetConfig()
	}
	return &configCheck
}

func checkError(err error) {
	statusError, _ := err.(*apierrors.StatusError)
	color := Red
	if statusError.ErrStatus.Reason == metav1.StatusReasonForbidden || statusError.ErrStatus.Reason == metav1.StatusReasonUnauthorized {
		fmt.Printf("%s%s: Make sure you are logged in%s", color, statusError.ErrStatus.Reason, Reset)
		os.Exit(0)
	}
}

func genJson(data Config) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Write the JSON byte array to a file
	file, err := os.Create("example-config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		panic(err)
	}
}
