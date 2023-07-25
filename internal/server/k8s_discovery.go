package server

import (
	"context"
	kuberegistry "github.com/go-kratos/kratos/contrib/registry/kubernetes/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func NewK8sDiscovery(ctx context.Context) *kuberegistry.Registry {
	clientSet, err := getClientSet()
	if err != nil {
		panic(err)
	}

	r := kuberegistry.NewRegistry(clientSet)

	svrHello := &registry.ServiceInstance{
		ID:        "1",
		Name:      "hello",
		Version:   "v1.0.0",
		Endpoints: []string{"http://127.0.0.1:80"},
	}
	registry.Register(ctx)
}

func getClientSet() (*kubernetes.Clientset, error) {
	restConfig, err := rest.InClusterConfig()
	home := homedir.HomeDir()

	if err != nil {
		kubeconfig := filepath.Join(home, ".kube", "config")
		restConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}
