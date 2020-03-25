package cmd

import (
	"fmt"
	"log"

	//"gopkg.in/yaml.v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getServices(clientset *kubernetes.Clientset, namespace string) {
	fmt.Println("")
	log.Println("All Services")
	fmt.Println("")
	services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get Services:", err)
	}
	for i, services := range services.Items {
		fmt.Printf("[%d] %s\n", i, services.GetName())
	}

	fmt.Println("")
	log.Println("Breakout Services")
	fmt.Println("")
	for _, v := range services.Items {
		fmt.Println("")
		//fmt.Println(v)
		fmt.Println("ServiceName:", v.GetName())
		fmt.Println("Namespace:", v.GetNamespace())
		fmt.Println("ClusterIP:", v.Spec.ClusterIP)
		fmt.Println("Port:", v.Spec.Ports[0].Port)
		fmt.Println("Target Ports:", v.Spec.Ports[0].TargetPort)
		fmt.Println("Proctol:", v.Spec.Ports[0].Protocol)
	}
}
