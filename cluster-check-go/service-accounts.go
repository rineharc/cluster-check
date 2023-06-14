package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func validateServiceAccounts(configCheck *Config, clientset *kubernetes.Clientset, sceCheck *bool, azureCheck *bool, omsCheck *bool) bool {
	fmt.Println(Separator)
	fmt.Println("Checking service accounts")
	fmt.Println(Separator)
	passed := true

	if *omsCheck {
		if !checkServiceAccounts(&configCheck.OMS.ServiceAccounts, clientset) {
			passed = false
		}
	}
	if *sceCheck {
		if !checkServiceAccounts(&configCheck.SCE.ServiceAccounts, clientset) {
			passed = false
		}
	}
	if *azureCheck {
		if !checkServiceAccounts(&configCheck.Azure.ServiceAccounts, clientset) {
			passed = false
		}
	}
	fmt.Print("\n")
	return passed
}

func checkServiceAccounts(serviceAccounts *[]NamespacedEntry, clientset *kubernetes.Clientset) bool {
	var passed = true
	for _, sa := range *serviceAccounts {
		status := ""
		color := ""
		_, err := clientset.CoreV1().ServiceAccounts(sa.Namespace).Get(context.TODO(), sa.Name, metav1.GetOptions{})
		if err != nil {
			status = "Not Found"
			color = Red
			passed = false
			checkError(err)
		} else {
			status = "Found"
			color = Green
		}
		fmt.Printf("%s/%s - %s%s %s\n", sa.Namespace, sa.Name, color, status, Reset)
	}
	return passed
}
