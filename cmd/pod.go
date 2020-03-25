package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	//"gopkg.in/yaml.v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func getPods(clientset *kubernetes.Clientset, namespace string) {
	var readyStatus string
	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}
	fmt.Fprintf(w, "\n %s\t%s\t\t%s\t%s\t", "Name", "Status", "Ready", "Age")
	fmt.Fprintf(w, "\n %s\t%s\t\t%s\t%s\t", "----", "----", "----", "----")
	fmt.Fprintln(w)
	defer w.Flush()
	for _, pod := range pods.Items {
		readyContainers := 0
		totalContainers := len(pod.Spec.Containers)
		podCreationTime := pod.GetCreationTimestamp()
		age := time.Since(podCreationTime.Time).Round(time.Minute)

		for i := len(pod.Status.ContainerStatuses) - 1; i >= 0; i-- {
			container := pod.Status.ContainerStatuses[i]
			if container.Ready && container.State.Running != nil {
				readyContainers++
			}
		}
		readyStatus = fmt.Sprint(readyContainers, "/", totalContainers)
		fmt.Fprintf(w, "%s\t%s\t\t%s\t%s\t\n", pod.GetName(), pod.Status.Phase, readyStatus, age.String())

		//fmt.Println(pod.Status.Phase)
		//fmt.Println(apiv1.PodSucceeded)

		//fmt.Printf("[%d] %s\n", i, pod.GetName(), "		", pod.Status.Phase)
		//fmt.Println("Request CPU ==> ", pod.Spec.Containers[0].Resources.Requests.Cpu(), " Request Memory ==> ", pod.Spec.Containers[0].Resources.Requests.Memory())
		//fmt.Println("Limit CPU ==> ", pod.Spec.Containers[0].Resources.Limits.Cpu(), " Limit Memory ==> ", pod.Spec.Containers[0].Resources.Limits.Memory())
		//	fmt.Println("")
		//	fmt.Println("")
	}
}

func getPodinService(clientset *kubernetes.Clientset, name string) {
	// Passthrough Clientset and Service name
	listOptions := metav1.ListOptions{}
	//listOptions
	//get all services
	svcs, err := clientset.CoreV1().Services("").List(listOptions)
	if err != nil {
		log.Fatal(err)
	}
	// loop through service obejects
	for _, svc := range svcs.Items {
		// match name to requested name

		if svc.Name == name {
			fmt.Fprintf(os.Stdout, "service name: %v\n", svc.Name)
			fmt.Println("	|")
			//fmt.Println("	|")
			fmt.Println("	--> Namespace:", svc.GetNamespace())
			fmt.Println("	--> ClusterIP:", svc.Spec.ClusterIP)
			fmt.Println("	--> Port:", svc.Spec.Ports[0].Port)
			fmt.Println("	--> Target Ports:", svc.Spec.Ports[0].TargetPort)
			fmt.Println("	--> Proctol:", svc.Spec.Ports[0].Protocol)
			set := labels.Set(svc.Spec.Selector)
			// get labels from svc.Spec.Selector and build listOptions with this label
			listOptions := metav1.ListOptions{LabelSelector: set.AsSelector().String()}
			// fetch through alll pods that have  this label
			pods, _ := clientset.CoreV1().Pods("").List(listOptions)
			// loop through pods Items and display
			for _, pod := range pods.Items {
				fmt.Println("	|")
				fmt.Println("	|")
				fmt.Fprintf(os.Stdout, "	--> backing pod name: %v\n", pod.Name)
				fmt.Fprintf(os.Stdout, "	--> backing pod IP: %v\n", pod.Status.PodIP)
				fmt.Fprintf(os.Stdout, "	--> backing pod status: %v\n", pod.Status.Phase)
			}
		}
	}

}

func deletePod(clientset *kubernetes.Clientset, podname string, ns string) {
	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{PropagationPolicy: &deletePolicy}
	if err := clientset.CoreV1().Pods(ns).Delete(podname, &deleteOptions); err != nil {
		fmt.Println("Error Failed to Delete Pod:")
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Deleting Pod:", podname)
	}
}
