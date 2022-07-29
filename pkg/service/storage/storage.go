package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Jooho/integration-framework-server/pkg/logger"
	"github.com/Jooho/integration-framework-server/pkg/utils"

	v1storage "github.com/Jooho/integration-framework-server/pkg/api/storage/v1"
	constants "github.com/Jooho/integration-framework-server/pkg/constants"
	templatev1 "github.com/openshift/api/template/v1"
	templatev1client "github.com/openshift/client-go/template/clientset/versioned/typed/template/v1"

	"google.golang.org/grpc"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type storageServer struct {
	v1storage.StorageServer
	Scheme         *runtime.Scheme
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

	template, err := s.templateClient.Templates(constants.TEMPLATE_NAMESPACE).Get(context.Background(), "storage-"+req.StorageType, metav1.GetOptions{})
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Fail to get storage-%s teamplate: '%v' ", req.StorageType, err))
		return &v1storage.GetStorageParamResponse{}, err
	}
	
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
	logger.Log.Debug("Entry storage.go - GetRenderedStorageManifest")

	stroageTemplateName := "storage-" + req.StorageType
	templateParams := req.Parameters

	storageTemplateObj, err := s.templateClient.Templates(constants.TEMPLATE_NAMESPACE).Get(context.TODO(), stroageTemplateName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return &v1storage.CreateStorageResponse{}, fmt.Errorf("template %q could not be found", stroageTemplateName)
		}
		return &v1storage.CreateStorageResponse{}, err
	}

	jsonString,err := processString(s.Scheme,templateParams,s.templateClient,storageTemplateObj)
	if err != nil{
		return &v1storage.CreateStorageResponse{}, err
	}

	return &v1storage.CreateStorageResponse{Manifest: jsonString}, nil
}

func processString(scheme *runtime.Scheme, values map[string]string, templateClient *templatev1client.TemplateV1Client, in *templatev1.Template) (string, error) {
	logger.Log.Debug("Entry storage.go - processString")

	var obj runtime.Object
	var scope conversion.Scope

	logger.Log.Debug(fmt.Sprintf("Template(%s) Parameters : %v", in.Name,values))

	resultObj, err := utils.Process(values, templateClient, in); 
	if err != nil {
		return "", err
	}
	logger.Log.Debug(fmt.Sprintf("ProcessedTemplate: %v", resultObj))
	runtime.Convert_runtime_RawExtension_To_runtime_Object(&resultObj.Objects[0], &obj, scope)
	jsonString := utils.ConvertK8StoJsonString(scheme, obj, false, false)
	return jsonString, nil
}