package cmd

import (
	"fmt"
	"log"

	//"gopkg.in/yaml.v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getConfigMaps(clientset *kubernetes.Clientset, namespace string) {
	fmt.Println("")
	cfmaps, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get ConfigMap:", err)
	}
	for i, cfmap := range cfmaps.Items {
		fmt.Printf("[%d] %s\n", i, cfmap.GetName())
	}

}
