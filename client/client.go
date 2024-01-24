package client

import (
	"fmt"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func K8Client() *kubernetes.Clientset {
	config, err := genericclioptions.NewConfigFlags(true).ToRESTConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("clientset not running")
	}

	return clientset
}
