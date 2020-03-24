/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Create and Apply Manifest",
	Long:  `Create and Apply Manifest similarly to "kubectl apply -f": `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apply called")
		kubeconfig := os.Getenv("kubeconfig")
		if kubeconfig == "" {
			fmt.Println("no env var found, falling back to config file")
			kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "kubeconfig")
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
		ns, _ := cmd.Flags().GetString("namespace")

		delete, _ := cmd.Flags().GetString("delete")
		if delete != "" {
			deleteDeployment(clientset, delete, ns)
			os.Exit(0)
		}

		file, _ := cmd.Flags().GetString("file")
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("File reading error", err)
			os.Exit(1)
		}

		deploy, _ := cmd.Flags().GetBool("deploy")
		if deploy == true {
			createDeploymentFromYaml(clientset, data, ns)
		}

	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	applyCmd.Flags().StringP("namespace", "n", "", "namespace")
	applyCmd.Flags().BoolP("deploy", "", false, "test deploy")
	applyCmd.Flags().StringP("delete", "d", "", "test deploy")
	applyCmd.Flags().StringP("file", "f", "", "file path")
}
