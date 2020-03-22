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
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
			and usage of using your command. For example:`,

	Run: func(cmd *cobra.Command, args []string) {
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

		pods, _ := cmd.Flags().GetBool("pods")
		if pods == true {
			ns, _ := cmd.Flags().GetString("namespace")
			if ns == "" {
				fmt.Println("namespace has not been diclared")
				os.Exit(1)
			} else {
				getPods(clientset, ns)
			}

		}

		fmt.Println("get called")

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//getCmd.PersistentFlags().String("n", "", "namespace")
	getCmd.Flags().StringP("namespace", "n", "", "namespace")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getCmd.Flags().BoolP("pods", "p", false, "")
}
