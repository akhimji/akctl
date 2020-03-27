package cmd

import (
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func buildClient(cfgFile string) (*kubernetes.Clientset, error) {
	var kubeconfig string
	if cfgFile != "" {
		kubeconfig = cfgFile
		log.Println(" ✓ Using kubeconfig file via flag: ", kubeconfig)
	} else {
		kubeconfig = os.Getenv("kubeconfig")
		if kubeconfig != "" {
			log.Println(" ✓ Using kubeconfig via OS ENV")
		} else {
			kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
			if _, err := os.Stat(kubeconfig); os.IsNotExist(err) {
				log.Println(" X kubeconfig Not Found, use --kubeconfig")
				os.Exit(1)
			} else {
				log.Println(" ✓ Using kubeconfig file via homedir: ", kubeconfig)
			}

		}

	}

	// Bootstrap k8s configuration
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return clientset, err
}
