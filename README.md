# akctl

akctl
```
This is a pet proeject to rebuild the kubectl project from scratch to further understand 
                        the kubernetes API and related Go framwork.

Usage:
  akctl [flags]
  akctl [command]

Available Commands:
  apply       Create and Apply Manifest
  get         get subfuction to pull data from the kubernets cluster
  help        Help about any command

Flags:
  -h, --help                help for akctl
      --kubeconfig string   config file (default is $HOME/.kube/kubeconfig)
  -t, --toggle              Help message for toggle

Use "akctl [command] --help" for more information about a command.
```

get subfunction
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

apply subfunction
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
