package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Jooho/integration-framework-server/pkg/logger"
	"github.com/Jooho/integration-framework-server/pkg/utils"

	v1storage "github.com/Jooho/integration-framework-server/pkg/api/v1/storage"
	templatev1 "github.com/openshift/api/template/v1"
	templatev1client "github.com/openshift/client-go/template/clientset/versioned/typed/template/v1"
	"github.com/openshift/library-go/pkg/template/templateprocessing"

	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type storageServer struct {
	Scheme *runtime.Scheme
	v1storage.StorageServer
	clientset      *kubernetes.Clientset
	templateClient *templatev1client.TemplateV1Client
	config         *rest.Config
}

func NewStorageServer(s grpc.Server, scheme *runtime.Scheme, c *kubernetes.Clientset, r *rest.Config) {
	logger.Log.Info("Adding Storage Service to gRPC server...")

	storage := &storageServer{}
	storage.Scheme = scheme
	storage.clientset = c
	storage.config = r
	// Create an OpenShift template/v1 client.
	t, err := templatev1client.NewForConfig(r)
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Fail to get templateclient %s", err.Error()))
		panic(err)
	}

	storage.templateClient = t
	v1storage.RegisterStorageServer(&s, storage)
}

func (s *storageServer) GetStorageParams(ctx context.Context, req *v1storage.GetStorageParamsRequest) (*v1storage.GetStorageParamResponse, error) {
	logger.Log.Debug("Entry storage.go - GetApplications")

	template, err := s.templateClient.Templates("if-templates").Get(context.Background(), "storage-"+req.StorageType, metav1.GetOptions{})
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Fail to get storage-%s teamplate: '%v' ", req.StorageType, err))
		return &v1storage.GetStorageParamResponse{}, err
	}
	// jsonpb.Marshaler
	parameterBytes, err := json.Marshal(template.Parameters)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("json marshal failed: %v", err))
		return &v1storage.GetStorageParamResponse{}, err
	}

	getStorageParamResponse := &v1storage.GetStorageParamResponse{
		StorageType: req.StorageType,
		Parameters:  string(parameterBytes),
	}

	logger.Log.Debug(fmt.Sprintf("getStorageParamResponse: '%s'", getStorageParamResponse))
	return getStorageParamResponse, nil
}

func (s *storageServer) GetRenderedStorageManifest(ctx context.Context, req *v1storage.CreateStorageRequest) (*v1storage.CreateStorageResponse, error) {

	var obj runtime.Object
	var scope conversion.Scope
	stroageTemplateName := "storage-" + req.StorageType
	templateParams := req.Parameters

	storageTemplateObj, err := s.templateClient.Templates("if-templates").Get(context.TODO(), stroageTemplateName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return &v1storage.CreateStorageResponse{}, fmt.Errorf("template %q could not be found", stroageTemplateName)
		}
		return &v1storage.CreateStorageResponse{}, err
	}

	if errs := injectUserVars(templateParams, storageTemplateObj, false); errs != nil {
		return &v1storage.CreateStorageResponse{}, kerrors.NewAggregate(errs)
	}
	resultObj := storageTemplateObj

	if resultObj, err = s.Process(storageTemplateObj); err != nil {
		return &v1storage.CreateStorageResponse{}, err
	}

	runtime.Convert_runtime_RawExtension_To_runtime_Object(&resultObj.Objects[0], &obj, scope)
	jsonString := utils.ConvertK8StoJsonString(s.Scheme, obj, false, false)

	return &v1storage.CreateStorageResponse{Manifest: jsonString}, nil
}

func (s *storageServer) Process(in *templatev1.Template) (*templatev1.Template, error) {
	template := &templatev1.Template{}
	err := s.templateClient.RESTClient().Post().
		Namespace("if-templates").
		Resource("processedTemplates").
		Body(in).Do(context.TODO()).Into(template)
	return template, err
}

func getParameterByName(t *templatev1.Template, name string) *templatev1.Parameter {
	for i, param := range t.Parameters {
		if param.Name == name {
			return &(t.Parameters[i])
		}
	}
	return nil
}

func injectUserVars(values map[string]string, t *templatev1.Template, ignoreUnknownParameters bool) []error {
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