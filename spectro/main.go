package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/spectrocloud/palette/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"os"
	"os/signal"
	"path/filepath"
)

var (
	spectroSchema = schema.GroupVersionResource{
		Group:    "cluster.spectrocloud.com",
		Version:  "v1alpha1",
		Resource: "spectroclusters",
	}
)

func main() {
	watchSpectroCluster(context.Background())
}

func watchSpectroCluster(ctx context.Context) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	klog.InitFlags(nil)

	klog.Infof("Starting test")
	//config, _ := rest.InClusterConfig()

	dynClient, errClient := dynamic.NewForConfig(GetConfig())
	if errClient != nil {
		klog.Fatalf("Error received creating client %v", errClient)
	}

	crdClient := dynClient.Resource(spectroSchema)

	watcherPredicate := func() metav1.ListOptions {
		return metav1.ListOptions{
			//LabelSelector: "mdfxr=mdfjjj",
			FieldSelector: "metadata.name=spectrocluster-sample",
		}
	}

	watcher, errCrd := crdClient.Namespace("default").Watch(watcherPredicate())
	if errCrd != nil {
		klog.Fatalf("Error getting CRD %v", errCrd)
	}


	go func() {
		for {
			select {
			case e := <-watcher.ResultChan():
				if e.Object == nil {
					return
				}

				spect, ok := e.Object.(*v1alpha1.SpectroCluster)
				if !ok {
					klog.Error("inside !ok")
					continue
				}

				klog.Info("Spectro: ", spect)
				byte, err := json.Marshal(e)
				if err != nil {
					klog.Error("Error: failed to parse ", err)
					return
				}

				klog.Infof("Event JSON : ", string(byte))

				switch e.Type {
					case watch.Modified:
						klog.Info("I am in Modified")
					default:
						klog.Info("I am in default")
					}
			case <-ctx.Done():
				watcher.Stop()
				return
			}
		}
	}()

	<-sigCh
}


func getSpectroCluster() {
	klog.InitFlags(nil)

	klog.Infof("Starting test")
	//config, _ := rest.InClusterConfig()

	dynClient, errClient := dynamic.NewForConfig(GetConfig())
	if errClient != nil {
		klog.Fatalf("Error received creating client %v", errClient)
	}

	crdClient := dynClient.Resource(spectroSchema)

	crd, errCrd := crdClient.Namespace("default").Get("spectrocluster-sample", metav1.GetOptions{})
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
