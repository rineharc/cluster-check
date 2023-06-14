package main

import (
	"context"
	"fmt"
	v1 "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func validateSCCs(securityClient v1.SecurityV1Interface, configCheck *Config) bool {
	var passed = true

	fmt.Println(Separator)
	fmt.Println("Checking security context constraints")
	fmt.Println(Separator)
	for _, scc := range configCheck.SecurityContexts {
		status := ""
		color := ""
		_, err := securityClient.SecurityContextConstraints().Get(context.TODO(), scc, metav1.GetOptions{})
		if err != nil {
			status = "Not Found"
			color = Red
			passed = false
			checkError(err)
		} else {
			status = "Found"
			color = Green
		}
		fmt.Printf("%s - %s%s %s\n", scc, color, status, Reset)
	}

	fmt.Print("\n")
	return passed
}
