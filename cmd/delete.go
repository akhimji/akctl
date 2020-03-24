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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
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
		clientset, err := kubernetes.NewForConfig(config)

		ns, _ := cmd.Flags().GetString("namespace")
		if ns == "" {
			fmt.Println("namespace has not been declared use: '-n <nanespace>")
			os.Exit(1)
		}
		deployment, _ := cmd.Flags().GetString("deployment")
		if deployment != "" {
			deleteDeployment(clientset, deployment, ns)
			os.Exit(0)
		}
		pod, _ := cmd.Flags().GetString("pod")
		if pod != "" {
			deletePod(clientset, pod, ns)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().StringP("namespace", "n", "", "namespace")
	deleteCmd.Flags().StringP("deployment", "d", "", "delete deployment  <name of deployment>")
	deleteCmd.Flags().StringP("pod", "p", "", "delete pod <name of deployment>")
}
