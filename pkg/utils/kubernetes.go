package utils

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Jooho/integration-framework-server/pkg/logger"
	"github.com/openshift/library-go/pkg/template/templateprocessing"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	constants "github.com/Jooho/integration-framework-server/pkg/constants"
	templatev1 "github.com/openshift/api/template/v1"
	templatev1client "github.com/openshift/client-go/template/clientset/versioned/typed/template/v1"
)

var (
	config    *rest.Config
	clientset *kubernetes.Clientset
	err       error
)

func GetK8SRestConfig(mode string) (*rest.Config, error) {
	logger.Log.Debug("Entry kubernetes.go - GetK8SRestConfig")

	if strings.ToLower(mode) == "local" {
		kubeconfig := filepath.Join(
			os.Getenv("HOME"), ".kube", "config",
		)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			logger.Log.Fatal(err.Error())
			return &rest.Config{}, err
		}

	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			logger.Log.Fatal(err.Error())
			return &rest.Config{}, err
		}
	}
	return config, nil
}
func GetK8SClientSet(mode string) (*kubernetes.Clientset, error) {
	logger.Log.Debug("Entry kubernetes.go - GetK8SClientSet")

	GetK8SRestConfig(mode)

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		logger.Log.Fatal(err.Error())
		return &kubernetes.Clientset{}, err
	}

	return clientset, nil
}

func ConvertK8StoJsonString(scheme *runtime.Scheme, obj runtime.Object, yamlType bool, pretty bool) string {
	gvks, _, err := scheme.ObjectKinds(obj)
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



func Process(values map[string]string,templateClient *templatev1client.TemplateV1Client, in *templatev1.Template) (*templatev1.Template, error) {
	logger.Log.Debug("Entry kubernetes.go - process")

	if errs := injectUserVars(values, in, true); errs != nil {
		return &templatev1.Template{}, err
	}
	logger.Log.Debug(fmt.Sprintf("Template(%s) after injectVariabels", in))
	template := &templatev1.Template{}
	err := templateClient.RESTClient().Post().
		Namespace(constants.TEMPLATE_NAMESPACE).
		Resource("processedTemplates").
		Body(in).Do(context.TODO()).Into(template)
	return template, err
}

// func getParameterByName(t *templatev1.Template, name string) *templatev1.Parameter {
// 	logger.Log.Debug("Entry kubernetes.go - getParameterByName")
	
// 	for i, param := range t.Parameters {
// 		if param.Name == name {
// 			return &(t.Parameters[i])
// 		}
// 	}
// 	return nil
// }

func injectUserVars(values map[string]string, t *templatev1.Template, ignoreUnknownParameters bool) []error {
	logger.Log.Debug("Entry kubernetes.go - injectUserVars")

	var errors []error
	for param, val := range values {
		v := templateprocessing.GetParameterByName(t, param)
		if v != nil {
			v.Value = val
			v.Generate = ""
		} else if !ignoreUnknownParameters {
			errors = append(errors, fmt.Errorf("unknown parameter name %q\n", param))
		}
	}
	return errors
}
