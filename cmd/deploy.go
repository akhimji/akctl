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

	//"gopkg.in/yaml.v2"

	"github.com/ghodss/yaml"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

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
