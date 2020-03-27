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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get subfuction to pull data from the kubernets cluster",
	//Long: `	-pods

	Run: func(cmd *cobra.Command, args []string) {
		clientset, err := buildClient(cfgFile)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(1)
		}

		getns, _ := cmd.Flags().GetBool("getns")
		if getns == true {
			getNamespaces(clientset)
		}

		ns, _ := cmd.Flags().GetString("namespace")
		svc, _ := cmd.Flags().GetString("service")
		if ns == "" {
			fmt.Println("namespace has not been declared use: '-n <nanespace>")
			os.Exit(1)
		}
		pods, _ := cmd.Flags().GetBool("pods")
		if pods == true {
			getPods(clientset, ns)
		}
		configmap, _ := cmd.Flags().GetBool("configmap")
		if configmap == true {
			getConfigMaps(clientset, ns)
		}
		ingress, _ := cmd.Flags().GetBool("ingress")
		if ingress == true {
			showIngress(clientset, ns)
		}
		services, _ := cmd.Flags().GetBool("services")
		if services == true {
			getServices(clientset, ns)
		}
		podsinsvc, _ := cmd.Flags().GetBool("podsinsvc")
		if podsinsvc == true {
			getPodinService(clientset, svc)
		}
		deployment, _ := cmd.Flags().GetBool("deployment")
		if deployment == true {
			getDeployment(clientset, ns)
		}
		//deploy, _ := cmd.Flags().GetBool("deploy")
		//if deploy == true {
		//	createDeploymentFromYaml(clientset, podAsYaml, ns)
		//}
		test, _ := cmd.Flags().GetBool("test")
		if test == true {
			getTest(clientset)
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//getCmd.PersistentFlags().String("n", "", "namespace")
	getCmd.Flags().StringP("namespace", "n", "", "namespace")
	getCmd.Flags().StringP("service", "s", "", "service")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getCmd.Flags().BoolP("pods", "p", false, "get pods")
	getCmd.Flags().BoolP("configmap", "c", false, "get configmap")
	getCmd.Flags().BoolP("ingress", "i", false, "get ingress")
	getCmd.Flags().BoolP("services", "", false, "get services")
	getCmd.Flags().BoolP("podsinsvc", "", false, "get pods behind a service")
	getCmd.Flags().BoolP("getns", "a", false, "get all namespaces")
	getCmd.Flags().BoolP("deployment", "d", false, "get deployment")
	getCmd.Flags().BoolP("test", "t", false, "test block")
}
