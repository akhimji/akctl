module github.com/alyarctiq/akctl

go 1.14

replace k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90

require (
	github.com/ghodss/yaml v1.0.0
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89 // indirect
)
