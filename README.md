
![Go](https://github.com/alyarctiq/akctl/workflows/Go/badge.svg)
![build-release](https://github.com/alyarctiq/akctl/workflows/build-release/badge.svg)

# akctl

### This is a project to rebuild core functionality of the kubectl project from scratch to further understand the kubernetes framework and related Go Kubernetes Spec.
```
Usage:
  akctl [flags]
  akctl [command]

Available Commands:
  apply       Apply subfuction
  delete      Delete subfuction
  get         Get subfuction to pull data from the kubernets cluster
  help        Help about any command

Flags:
  -h, --help                help for akctl
      --kubeconfig string   config file (default is $HOME/.kube/kubeconfig)
  -t, --toggle              Help message for toggle

Use "akctl [command] --help" for more information about a command.
```

### get subfunction
```
get subfuction to pull data from the kubernets cluster

Usage:
  akctl get [flags]

Flags:
  -c, --configmap          get configmap
  -d, --deployment         get deployment
  -a, --getns              get all namespaces
  -h, --help               help for get
  -i, --ingress            get ingress
  -n, --namespace string   namespace
  -p, --pods               get pods
      --podsinsvc          get pods behind a service
  -s, --service string     service
      --services           get services
  -t, --test               test block

Global Flags:
      --kubeconfig string   config file (default is $HOME/.kube/kubeconfig)
```
Get Pods
![Get Pods](/screencap/getpods.png)

Get ConfigMap
![Get  ConfigMap](/screencap/getmc.png)

Get Ingress
![Get Ingress](/screencap/getIngress.png)

Get Service in Namespace
![Get Svc](/screencap/getsvc.png)

Get Backing Pods for Service
![Get Pods in Svc](/screencap/getpodsinsvc.png)

### apply subfunction
```
Create and Apply Manifest similarly to "kubectl apply -f":

Usage:
  akctl apply [flags]

Flags:
  -d, --delete string      test deploy
      --deploy             test deploy
  -f, --file string        file path
  -h, --help               help for apply
  -n, --namespace string   namespace

Global Flags:
      --kubeconfig string   config file (default is $HOME/.kube/kubeconfig)
```
Apply Deployment Config Yaml
![deply](/screencap/deploy.png)

### delete subfunction
```
delete subfuction similar to "kubectl delete ":

Usage:
  akctl delete [flags]

Flags:
  -d, --deployment string   delete --deployment  <name of deployment>
  -h, --help                help for delete
  -n, --namespace string    namespace
  -p, --pod string          delete --pod <name of pod>

Global Flags:
      --kubeconfig string   config file (default is $HOME/.kube/kubeconfig)
```
Delete Deployment Config 
![deply](/screencap/delete.png)
