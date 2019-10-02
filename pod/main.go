package main

import (
	"flag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"os"
	"path/filepath"
)

var (
	podsSchema = schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}
)

func main() {
	klog.InitFlags(nil)

	klog.Infof("Starting test")
	//config, _ := rest.InClusterConfig()

	dynClient, errClient := dynamic.NewForConfig(GetConfig())
	if errClient != nil {
		klog.Fatalf("Error received creating client %v", errClient)
	}

	crdClient := dynClient.Resource(podsSchema)

	crd, errCrd := crdClient.Namespace("default").Get("nginx-deployment-test-b5b9d9ccf-n9jvq", metav1.GetOptions{})
	if errCrd != nil {
		klog.Fatalf("Error getting CRD %v", errCrd)
	}
	klog.Infof("Got CRD: %v", crd)
}

func GetConfig() *rest.Config {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	return config
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
