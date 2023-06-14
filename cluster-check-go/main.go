package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/openshift/client-go/security/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/util/homedir"
)

func usage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Printf("  %s -scc-check : checks for configured security contexts in addition to pods, service accounts and namespaces\n", os.Args[0])
}

func main() {
	configFile := flag.String("config", "", "path to config json file, if not provided will check defaults")
	omsCheck := flag.Bool("oms-check", true, "checks that oms resources are created (defaults to true)")
	sceCheck := flag.Bool("sce-check", false, "perform SCE check (defaults to false)")
	sccCheck := flag.Bool("scc-check", false, "checks that security contexts are created (defaults to false)")
	azureCheck := flag.Bool("azure-check", false, "checks that azure resources are created (defaults to false)")
	genJsonCheck := flag.Bool("gen-json", false, "generate example config json that can be used to override defaults (defaults to false)")
	// storeNum := flag.String("store-num", "d", "number of stores to create")
	var kubeconfig *string
	passed := true
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Usage = usage
	flag.Parse()

	// if *storeNum == "" {
	// 	fmt.Println("Please provide a number of stores to create")
	// 	reader := bufio.NewReader(os.Stdin)
	// 	input, _ := reader.ReadString('\n')
	// 	*storeNum = input
	// }

	if *genJsonCheck {
		var baseConfig Config
		baseConfig = baseConfig.GetConfig()
		genJson(baseConfig)
		os.Exit(0)
	}
	configCheck := loadConfig(configFile)

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	if *sccCheck {
		openshiftClient, err := versioned.NewForConfig(config)
		if err != nil {
			fmt.Println("Error creating OpenShift clientset:", err)
			os.Exit(1)
		}
		sccCheck := validateSCCs(openshiftClient.SecurityV1(), configCheck)
		if !sccCheck {
			passed = false
		}
	}

	nsCheck := validateNamespaces(configCheck, clientset, sceCheck, azureCheck, omsCheck)
	if !nsCheck {
		passed = false
	}
	saCheck := validateServiceAccounts(configCheck, clientset, sceCheck, azureCheck, omsCheck)
	if !saCheck {
		passed = false
	}
	podCheck := validatePods(configCheck, clientset, sceCheck, azureCheck, omsCheck)
	if !podCheck {
		passed = false
	}

	fmt.Print("\n" + Separator + "\n")
	if passed {
		fmt.Printf("%sAll checks passed%s\n", Green, Reset)
	} else {
		fmt.Printf("%sSome checks failed%s\n", Red, Reset)
	}
	fmt.Println(Separator)
}
