package main

import (
	"context"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func validatePods(configCheck *Config, clientset *kubernetes.Clientset, sceCheck *bool, azureCheck *bool, omsCheck *bool) bool {
	fmt.Println(Separator)
	fmt.Println("Checking for pods")
	fmt.Println(Separator)
	passed := true

	if *omsCheck {
		if !checkPods(&configCheck.OMS.Pods, clientset) {
			passed = false
		}
	}
	if *sceCheck {
		if !checkPods(&configCheck.SCE.Pods, clientset) {
			passed = false
		}
	}
	if *azureCheck {
		if !checkPods(&configCheck.Azure.Pods, clientset) {
			passed = false
		}
	}
	fmt.Print("\n")
	return passed
}

func checkPods(pods *[]NamespacedEntry, clientset *kubernetes.Clientset) bool {
	var passed = true
	for _, podTest := range *pods {
		status := ""
		color := ""
		pods, err := clientset.CoreV1().Pods(podTest.Namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			status = "Not Found"
			color = Red
			passed = false
			checkError(err)
		} else {
			found := false
			for _, pod := range pods.Items {
				if strings.Contains(pod.Name, podTest.Name) {
					found = true
					break
				}
			}

			if found {
				status = "Found"
				color = Green
			} else {
				status = "Not Found"
				color = Red
				passed = false
			}
		}
		fmt.Printf("%s/*%s* - %s%s %s\n", podTest.Namespace, podTest.Name, color, status, Reset)
	}
	return passed
}
