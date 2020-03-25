package cmd

import (
	"fmt"
	"log"

	//"gopkg.in/yaml.v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getNamespaces(clientset *kubernetes.Clientset) {
	fmt.Println("")
	log.Println("All Namespaces")
	fmt.Println("")
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get namespace:", err)
	}
	for i, namespace := range namespaces.Items {
		fmt.Printf("[%d] %s\n", i, namespace.GetName())
	}
	//fmt.Println(namespaces)
}
