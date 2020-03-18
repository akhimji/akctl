package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// This program lists the pods in a cluster equivalent to
//
// kubectl get pods
//

func usage() {
	ProjectName := "AKClient"
	Version := "v0.1"
	fmt.Printf("🏔️ %s %s\n", ProjectName, Version)
	fmt.Println("Author: Aly Khimji")
	fmt.Print("\nUsage: ak [-pods|-configmap|-ingress]\n")
	fmt.Println("Options:")
	fmt.Println("    --config\tConfiguration path")
	fmt.Println("    --help\tHelp info")
}

func getPods(clientset *kubernetes.Clientset) {
	fmt.Println("")
	log.Println("All Pods")
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}

	// print pods
	for i, pod := range pods.Items {
		fmt.Printf("[%d] %s\n", i, pod.GetName())
	}
	os.Exit(0)
}

func getConfigMaps(clientset *kubernetes.Clientset) {
	fmt.Println("")
	log.Println("All ConfigMaps")
	cfmaps, err := clientset.CoreV1().ConfigMaps("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get ConfigMap:", err)
	}
	// print pods
	for i, cfmaps := range cfmaps.Items {
		fmt.Printf("[%d] %s\n", i, cfmaps.GetName())
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
		//fmt.Printf("%#v\n", v)
		if v.Spec.Rules[0].HTTP.Paths[0].Backend.ServicePort.IntVal == 80 {
			log.Println("Ingress:", i)
			//fmt.Println(v.Spec.Rules[0].HTTP.Paths[0].Backend.ServicePort.IntVal)
			//fmt.Println(v.Spec)
			//fmt.Println(v.Spec.Rules[0])
			//fmt.Println(v.Spec.Rules[0].HTTP)
			fmt.Println("Ingress ServiceName:", v.Spec.Rules[0].HTTP.Paths[0].Backend.ServiceName)
			//fmt.Println(v.Spec.Rules[0].HTTP.Paths[0].Backend.ServiceName)
			fmt.Println("Ingress ServicePort:", v.Spec.Rules[0].HTTP.Paths[0].Backend.ServicePort.IntValue())
			//fmt.Println(v.Spec.Rules[0].HTTP.Paths[0].Backend.ServicePort.IntValue())
			fmt.Println("Ingress Host:", v.Spec.Rules[0].Host)
			//fmt.Println(v.Spec.Rules[0].Host)
			fmt.Println("")
		}
	}
}

// func showIngress(clientset *kubernetes.Clientset) {
// 	list, _ := clientset.ExtensionsV1beta1().Ingresses("").List(metav1.ListOptions{}) // "" is all namespaces
// 	for _, ingress := range list.Items {
// 		for _, rule := range ingress.Spec.Rules {
// 			for _, path := range rule.HTTP.Paths {
// 				service, _ := clientset.CoreV1().Services(ingress.GetObjectMeta().GetNamespace()).Get(path.Backend.ServiceName, metav1.GetOptions{})
// 				host := rule.Host
// 				fmt.Println(host)
// 				path := path.Path
// 				fmt.Println(path)
// 				backend := ingress.spec
// 				fmt.Println(backend)
// 				destination := service.Spec.ClusterIP
// 				fmt.Println(destination)
// 			}
// 		}
// 	}
// }

func getServices(clientset *kubernetes.Clientset) {
	fmt.Println("")
	log.Println("All Services")
	fmt.Println("")
	services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
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
		fmt.Println(v)
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
			set := labels.Set(svc.Spec.Selector)
			// get labels from svc.Spec.Selector and build listOptions with this label
			listOptions := metav1.ListOptions{LabelSelector: set.AsSelector().String()}
			// fetch through alll pods that have  this label
			pods, _ := clientset.CoreV1().Pods("").List(listOptions)
			// loop through pods Items and display
			for _, pod := range pods.Items {
				fmt.Fprintf(os.Stdout, "backing pod name: %v\n", pod.Name)
			}
		}
	}

}

func startArgs(clientset *kubernetes.Clientset) {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}
	podsPtr := flag.Bool("pods", false, "Pods")
	cfmapPtr := flag.Bool("configmap", false, "Config Maps")
	ingressPtr := flag.Bool("ingress", false, "Ingresses and Details")
	servicesPtr := flag.Bool("services", false, "Services and Details")
	namespacesPtr := flag.Bool("namespaces", false, "namespaces")
	podsinservicePtr := flag.Bool("podsinservice", false, "podsinservice")
	testPtr := flag.Bool("test", false, "test")
	flag.Parse()

	if *podsPtr == true {
		getPods(clientset)
		os.Exit(0)
	} else if *cfmapPtr == true {
		getConfigMaps(clientset)
		os.Exit(0)
	} else if *ingressPtr == true {
		getIngress(clientset)
		os.Exit(0)
	} else if *servicesPtr == true {
		getServices(clientset)
		os.Exit(0)
	} else if *namespacesPtr == true {
		getNamespaces(clientset)
		os.Exit(0)
	} else if *testPtr == true {
		getTest(clientset)
		os.Exit(0)
	} else if *podsinservicePtr == true {
		svcname := os.Args[2]
		getPodinService(clientset, svcname)
		os.Exit(0)
	} else {
		fmt.Println("Try Again..")
		os.Exit(0)
	}

}

func main() {

	var ns string
	flag.StringVar(&ns, "namespace", "", "namespace")
	kubeconfig := os.Getenv("kubeconfig")
	if kubeconfig == "" {
		fmt.Println("no env var found, falling back to config file")
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
		log.Println("Using kubeconfig file: ", kubeconfig)
	}
	// Bootstrap k8s configuration from local 	Kubernetes config file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create an rest client not targeting specific API version
	clientset, err := kubernetes.NewForConfig(config)
	//fmt.Println(reflect.TypeOf(clientset))
	if err != nil {
		log.Fatal(err)
	}

	startArgs(clientset)

	//name := "istio-operator-metrics"
	//getPodinService(clientset, name)

}
