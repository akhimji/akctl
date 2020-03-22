package cmd

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

func startArgs(clientset *kubernetes.Clientset) {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}
	podsPtr := flag.Bool("pods", false, "pods")
	nsPtr := flag.String("n", "", "namespace")
	svcPtr := flag.String("s", "", "service")
	cfmapPtr := flag.Bool("configmaps", false, "Config Maps")
	ingressPtr := flag.Bool("ingress", false, "Ingresses and Details")
	servicesPtr := flag.Bool("services", false, "Services and Details")
	namespacesPtr := flag.Bool("namespaces", false, "namespaces")
	podsinservicePtr := flag.Bool("podsinservice", false, "podsinservice")
	testPtr := flag.Bool("test", false, "test")

	flag.Parse()

	if *podsPtr == true {
		getPods(clientset, *nsPtr)
		os.Exit(0)
	}
	if *cfmapPtr == true {
		getConfigMaps(clientset, *nsPtr)
		os.Exit(0)
	}
	if *ingressPtr == true {
		showIngress(clientset, *nsPtr)
		fmt.Println("")
		//getIngress(clientset)
		os.Exit(0)
	}
	if *servicesPtr == true {
		getServices(clientset, *nsPtr)
		os.Exit(0)
	}
	if *namespacesPtr == true {
		getNamespaces(clientset)
		os.Exit(0)
	}
	if *testPtr == true {
		getTest(clientset)
		os.Exit(0)
	}
	if *podsinservicePtr == true {
		getPodinService(clientset, *svcPtr)
		os.Exit(0)
	}
	fmt.Println("Try Again..")
	os.Exit(0)
}

func main() {

	var kc string
	flag.StringVar(&kc, "kubeconfig", "", "kubeconfig")
	//flag.Parse()

	kubeconfig := os.Getenv("kubeconfig")
	if kubeconfig == "" {
		fmt.Println("no env var found, falling back to config file")
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
		log.Println(" ✓ Using kubeconfig file: ", kubeconfig)
		fmt.Println("")
	} else {
		log.Println(" ✓ Using kubeconfig via OS ENV")
		fmt.Println("")
	}
	// Bootstrap k8s configuration
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
