package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func validateNamespaces(configCheck *Config, clientset *kubernetes.Clientset, sceCheck *bool, azureCheck *bool, omsCheck *bool) bool {
	fmt.Println(Separator)
	fmt.Println("Checking namespaces")
	fmt.Println(Separator)
	passed := true
	if *omsCheck {
		if !checkNamespaces(&configCheck.OMS.Namespaces, clientset) {
			passed = false
		}
	}
	if *sceCheck {
		if !checkNamespaces(&configCheck.SCE.Namespaces, clientset) {
			passed = false
		}
	}
	if *azureCheck {
		if !checkNamespaces(&configCheck.Azure.Namespaces, clientset) {
			passed = false
		}
	}
	fmt.Print("\n")
	return passed
}

func checkNamespaces(namespaces *[]string, clientset *kubernetes.Clientset) bool {
	var passed = true
	for _, ns := range *namespaces {
		status := ""
		color := ""
		_, err := clientset.CoreV1().Namespaces().Get(context.TODO(), ns, metav1.GetOptions{})
		if err != nil {
			status = "Not Found"
			color = Red
			passed = false
			checkError(err)
		} else {
			status = "Found"
			color = Green
		}
		fmt.Printf("%s - %s%s %s\n", ns, color, status, Reset)
	}
	return passed
}
