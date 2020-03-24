package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ghodss/yaml"
	//"gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func usage() {
	ProjectName := "akctl"
	Version := "v0.1"
	fmt.Printf("🏔️ ✓ %s %s\n", ProjectName, Version)
	fmt.Println("Author: Aly Khimji")
	fmt.Print("\nUsage: akctl  ")
	fmt.Println("Options:")
	fmt.Println("    --config\tConfiguration path")
	fmt.Println("    --help\tHelp info")
}

func getPods(clientset *kubernetes.Clientset, namespace string) {
	fmt.Println("")
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}
	for i, pod := range pods.Items {
		fmt.Printf("[%d] %s\n", i, pod.GetName())
		fmt.Println("Request CPU ==> ", pod.Spec.Containers[0].Resources.Requests.Cpu(), " Request Memory ==> ", pod.Spec.Containers[0].Resources.Requests.Memory())
		fmt.Println("Limit CPU ==> ", pod.Spec.Containers[0].Resources.Limits.Cpu(), " Limit Memory ==> ", pod.Spec.Containers[0].Resources.Limits.Memory())
		fmt.Println("")
		fmt.Println("")
	}
}

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

func getDeployment(clientset *kubernetes.Clientset, namespace string) {
	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	fmt.Printf("Listing deployments in namespace %q:\n", namespace)
	deployments, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range deployments.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
		fmt.Printf("%s", d.ObjectMeta.SelfLink)
		fmt.Println("")
	}

}

func getTest(clientset *kubernetes.Clientset) {
	fmt.Println("")
	log.Println("All Namespaces")
	fmt.Println("")
	pods, err := clientset.CoreV1().Pods("kube-system").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get pods from namespace:", err)
	}
	for _, p := range pods.Items {
		fmt.Println(p.GetName())
	}

	list, err := clientset.CoreV1().Pods("kube-system").List(metav1.ListOptions{})
	for _, l := range list.Items {
		fmt.Println("Request CPU ==> ", l.Spec.Containers[0].Resources.Requests.Cpu(), " Request Memory ==> ", l.Spec.Containers[0].Resources.Requests.Memory())
		fmt.Println("Limit CPU ==> ", l.Spec.Containers[0].Resources.Limits.Cpu(), " Limit Memory ==> ", l.Spec.Containers[0].Resources.Limits.Memory())
	}
	nodeList, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	for _, nodeList := range nodeList.Items {
		fmt.Println(nodeList.GetName())
	}

	fmt.Println("------------")
	configmaps, err := clientset.CoreV1().ConfigMaps("").List(metav1.ListOptions{})
	for _, cm := range configmaps.Items {
		fmt.Println(cm.GetName())

	}
	fmt.Println("------------")
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deploymentk
	fmt.Println("Creating deployment...")
	//result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

}
func int32Ptr(i int32) *int32 { return &i }

func createDeploymentFromYaml(clientset *kubernetes.Clientset, podAsYaml []byte, ns string) error {
	//podAsYaml []byte
	//This is received in byte format after reading it from disk.
	fmt.Println("Attempting Deployment..")
	var deployment appsv1.Deployment
	err := yaml.Unmarshal(podAsYaml, &deployment)
	if err != nil {
		fmt.Println("Error Unmarshaling:", err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	// pointer to deployment object
	result, err := deploymentsClient.Create(&deployment)
	//pod, poderr := clientset.CoreV1().Pods(ns).Create(&deployment)
	if err != nil {
		fmt.Println("Error Creating Deployment:")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	return nil
}

func deleteDeployment(clientset *kubernetes.Clientset, deployment string, ns string) {
	// build client set
	deploymentsClient := clientset.AppsV1().Deployments(ns)
	// build delete policy
	deletePolicy := metav1.DeletePropagationForeground
	// From Docs "PropagationPolicy    *DeletionPropagation"  in json format and *DeletionPropagation is pointer to metav1.DeletePropagationForeground
	//(&deletePolicy is pointer to deletePolicy)
	deleteOptions := metav1.DeleteOptions{PropagationPolicy: &deletePolicy}
	err := deploymentsClient.Delete(deployment, &deleteOptions)
	if err != nil {
		fmt.Println("Error Deleting Deployment:")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Deleted deployment.")

}
