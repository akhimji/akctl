package cmd

import (
	"fmt"
	"log"

	//"gopkg.in/yaml.v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getIngress(clientset *kubernetes.Clientset) {
	fmt.Println("")
	log.Println("All Ingress")
	list, err := clientset.ExtensionsV1beta1().Ingresses("").List(metav1.ListOptions{}) // "" is all namespaces
	if err != nil {
		log.Fatalln("failed to get Ingress:", err)
	}
	for i, ingress := range list.Items {
		fmt.Printf("[%d] %s\n", i, ingress.GetName())
	}

	fmt.Println("")
	log.Println("Breakout Ingress")
	fmt.Println("")
	ingress, err := clientset.ExtensionsV1beta1().Ingresses("").List(metav1.ListOptions{}) // "" is all namespaces
	for i, v := range ingress.Items {
		fmt.Println("Ingress:", i)
		fmt.Println("Ingress ServiceName:", v.Spec.Rules[0].HTTP.Paths[0].Backend.ServiceName)
		fmt.Println("Ingress ServicePort:", v.Spec.Rules[0].HTTP.Paths[0].Backend.ServicePort.IntValue())
		fmt.Println("Ingress Host:", v.Spec.Rules[0].Host)
		fmt.Println("")
		//}
	}
}

func showIngress(clientset *kubernetes.Clientset, namespace string) {
	list, _ := clientset.ExtensionsV1beta1().Ingresses(namespace).List(metav1.ListOptions{}) // "" is all namespaces
	for i, ingress := range list.Items {
		for _, rule := range ingress.Spec.Rules {
			for _, path := range rule.HTTP.Paths {
				service, _ := clientset.CoreV1().Services(ingress.GetObjectMeta().GetNamespace()).Get(path.Backend.ServiceName, metav1.GetOptions{})
				host := rule.Host
				fmt.Println("Ingress Index:", i)
				fmt.Println("Namespace:", service.GetNamespace())
				fmt.Println("Ingress Host:", host)
				path := path.Path
				fmt.Println("Ingress Path:", path)
				//backend := ingress.Spec.Backend
				//fmt.Println("Backend:", ingress.Spec.Rules[0].HTTP.Paths[0].Backend)
				destination := service.Spec.ClusterIP
				fmt.Println("Service ClusterIP:", destination)
				fmt.Println("Ingress ServiceName:", ingress.Spec.Rules[0].HTTP.Paths[0].Backend.ServiceName)
				fmt.Println("Ingress ServicePort:", ingress.Spec.Rules[0].HTTP.Paths[0].Backend.ServicePort.IntValue())
				//fmt.Println("Ingress Host:", ingress.Spec.Rules[0].Host)
				fmt.Println("")
			}
		}
	}
}
