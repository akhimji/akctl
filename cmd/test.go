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
	"log"

	//"gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
