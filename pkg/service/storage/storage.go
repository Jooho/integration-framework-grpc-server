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


func (s *storageServer) GetStorageTypes(ctx context.Context, req *v1storage.GetStorageTypesRequest) (*v1storage.GetStorageTypesResponse, error) {
	storageTypes, err := s.getStorageTypes()
	if err != nil {
		return &v1storage.GetStorageTypesResponse{}, err
	}
	return &v1storage.GetStorageTypesResponse{Types: storageTypes}, nil
}

func (s *storageServer) GetStorageParams(ctx context.Context, req *v1storage.GetStorageParamsRequest) (*v1storage.GetStorageParamResponse, error) {
	logger.Log.Debug("Entry storage.go - GetApplications")

	template, err := s.templateClient.Templates(constants.TEMPLATE_NAMESPACE).Get(context.Background(), "storage-"+req.Type, metav1.GetOptions{})
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Fail to get storage-%s teamplate: '%v' ", req.Type, err))
		return &v1storage.GetStorageParamResponse{}, err
	}
	
	parameterBytes, err := json.Marshal(template.Parameters)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("json marshal failed: %v", err))
		return &v1storage.GetStorageParamResponse{}, err
	}

	getStorageParamResponse := &v1storage.GetStorageParamResponse{
		Type:       req.Type,
		Parameters: parameterBytes,
	}
	
	logger.Log.Debug(fmt.Sprintf("getStorageParamResponse: '%s'", getStorageParamResponse))
	return getStorageParamResponse, nil
}

func (s *storageServer) ListStorage(ctx context.Context, req *v1storage.ListStorageRequest) (*v1storage.ListStorageResponse, error) {

	//Get StorageTypes: [s3,azure]
	storageTypes, err := s.getStorageTypes()
	if err != nil {
		logger.Log.Error(fmt.Sprintf("No Storage Templates exist: %v", err))
		return &v1storage.ListStorageResponse{}, err
	}
	
	//Gather Created Storage List
	var storageList map[string]*v1storage.StorageList = make(map[string]*v1storage.StorageList)
	var storageNames []string
	
	secretList, err := s.clientset.CoreV1().Secrets(req.Namespace).List(ctx, metav1.ListOptions{LabelSelector: "opendatahub.io/integration=storage"})
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Failed to get secrets in namespace(%s) ", req.Namespace))
		return &v1storage.ListStorageResponse{}, nil
	}
	for _,secret := range secretList.Items{
		logger.Log.Debug(fmt.Sprintf("Secret: %s", secret.Name))
		for _,storageType := range storageTypes{
			if secret.GetLabels()["opendatahub.io/storage-type"] == storageType{
				storageNames = append(storageNames, secret.Name)
			}
			storageList[storageType]=&v1storage.StorageList{Items: storageNames}
		}
	}
	if len(storageList) == 0{
		logger.Log.Info("No storage exist")
	}
	return &v1storage.ListStorageResponse{Storages: storageList}, nil
}

func (s *storageServer) GetRenderedStorageManifest(ctx context.Context, req *v1storage.RenderedStorageRequest) (*v1storage.RenderedStorageResponse, error) {
	logger.Log.Debug("Entry storage.go - GetRenderedStorageManifest")
	
	storageTemplateName := "storage-" + req.Type
	templateParams := req.Parameters
	
	storageTemplateObj, err := s.templateClient.Templates(constants.TEMPLATE_NAMESPACE).Get(context.TODO(), storageTemplateName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return &v1storage.RenderedStorageResponse{}, fmt.Errorf("template %q could not be found", storageTemplateName)
		}
		return &v1storage.RenderedStorageResponse{}, err
	}
	
	jsonString, err := processString(s.Scheme, templateParams, s.templateClient, storageTemplateObj)
	if err != nil {
		return &v1storage.RenderedStorageResponse{}, err
	}
	logger.Log.Debug(fmt.Sprintf("Secret :%v", []byte(jsonString)))
	
	return &v1storage.RenderedStorageResponse{Manifest: []byte(jsonString)}, nil
}

func processString(scheme *runtime.Scheme, values map[string]string, templateClient *templatev1client.TemplateV1Client, in *templatev1.Template) (string, error) {
	logger.Log.Debug("Entry storage.go - processString")
	
	var obj runtime.Object
	var scope conversion.Scope
	
	logger.Log.Debug(fmt.Sprintf("Template(%s) Parameters : %v", in.Name, values))
	
	resultObj, err := utils.Process(values, templateClient, in)
	if err != nil {
		return "", err
	}
	logger.Log.Debug(fmt.Sprintf("ProcessedTemplate: %v", resultObj))
	runtime.Convert_runtime_RawExtension_To_runtime_Object(&resultObj.Objects[0], &obj, scope)
	jsonString := utils.ConvertK8StoJsonString(scheme, obj, false, false)
	return jsonString, nil
}

func (s *storageServer) getStorageTypes() ([]string, error) {
	storageTypes := []string{}
	storageTemplates, err := s.templateClient.Templates(constants.TEMPLATE_NAMESPACE).List(context.Background(), metav1.ListOptions{LabelSelector: "integration=storage"})
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Fail to list storage templates: %v", err))
		return nil, err
	}
	for _, s := range storageTemplates.Items {
		storageTypes = append(storageTypes, s.GetLabels()["storage.type"])
	}

	logger.Log.Debug(fmt.Sprintf("Registered Storage Types: %s", storageTypes))
	return storageTypes, nil
}