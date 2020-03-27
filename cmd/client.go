/*
Copyright © 2020 akctl aly.khimji@arctiq.ca

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
