package utils

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jooho/integration-framework-server/pkg/logger"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	config    *rest.Config
	clientset *kubernetes.Clientset
	err       error
)

func GetK8SClientSet(mode string) (*kubernetes.Clientset, error) {
	logger.Log.Debug("Entry kubernetes.go - GetK8SClientSet")

	if strings.ToLower(mode) == "local" {
		kubeconfig := filepath.Join(
			os.Getenv("HOME"), ".kube", "config",
		)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			logger.Log.Fatal(err.Error())
			return &kubernetes.Clientset{},  err
		}

	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			logger.Log.Fatal(err.Error())
			return &kubernetes.Clientset{}, err
		}
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		logger.Log.Fatal(err.Error())
		return &kubernetes.Clientset{}, err
	}
	return clientset, nil
}

func ConvertK8StoJsonString(obj runtime.Object, yamlType bool, pretty bool) string {
	gvks, _, err := scheme.Scheme.ObjectKinds(obj)
	if err != nil {
		logger.Log.Error(err.Error())
	}

	for _, gvk := range gvks {
		if len(gvk.Kind) == 0 {
			continue
		}
		if len(gvk.Version) == 0 || gvk.Version == runtime.APIVersionInternal {
			continue
		}
		obj.GetObjectKind().SetGroupVersionKind(gvk)
		break
	}

	return bytes.NewBuffer(JsonSerializer(obj, yamlType, pretty)).String()
}
