package modelserving

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	msv1 "github.com/Jooho/integration-framework-server/pkg/api/modelserving/v1"
	odhitgv1alpha1client "github.com/Jooho/integration-framework-server/pkg/clientset/versioned/typed/odhintegration/v1alpha1"
	constants "github.com/Jooho/integration-framework-server/pkg/constants"
	"github.com/Jooho/integration-framework-server/pkg/logger"
	"github.com/Jooho/integration-framework-server/pkg/service/storage"
	"github.com/Jooho/integration-framework-server/pkg/utils"
	templatev1 "github.com/openshift/api/template/v1"
	templatev1client "github.com/openshift/client-go/template/clientset/versioned/typed/template/v1"
	operatorv1alpha1client "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned/typed/operators/v1alpha1"

	semver "github.com/coreos/go-semver/semver"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	// "k8s.io/apimachinery/pkg/runtime/schema"
)

type modelServingServer struct {
	msv1.ModelServingServer
	Scheme               *runtime.Scheme
	clientset            *kubernetes.Clientset
	templateClient       *templatev1client.TemplateV1Client
	odhintegrationClient *odhitgv1alpha1client.ODHIntegrationV1Alpha1Client
	operatorClient       *operatorv1alpha1client.OperatorsV1alpha1Client
	config               *rest.Config
}

func NewModelServingServer(s grpc.Server, scheme *runtime.Scheme, c *kubernetes.Clientset, r *rest.Config) {
	logger.Log.Debug("Adding Model Serving Service to gRPC server...")

	modelserving := &modelServingServer{}
	modelserving.Scheme = scheme
	modelserving.clientset = c
	modelserving.config = r

	// Create an OpenShift template/v1 client.
	t, err := templatev1client.NewForConfig(r)
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Fail to get templateclient %s", err.Error()))
	}

	modelserving.templateClient = t

	// Create an ODHIntegration/v1alpha1 client
	odhintegrationClient, err := odhitgv1alpha1client.NewForConfig(r)
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Fail to get odhintegrationClient %s", err.Error()))
	}
	modelserving.odhintegrationClient = odhintegrationClient

	// Create an openshift CSV/v1
	operatorClient, err := operatorv1alpha1client.NewForConfig(r)
	if err != nil {
		logger.Log.Fatal(fmt.Sprintf("Fail to get operatorClient %s", err.Error()))
	}
	modelserving.operatorClient = operatorClient
	msv1.RegisterModelServingServer(&s, modelserving)
}

func (m *modelServingServer) ListApp(ctx context.Context, req *emptypb.Empty) (*msv1.GetAppResponse, error) {
	logger.Log.Debug("Entry modelserving.go - GetApplications")

	appList := []*msv1.Application{}
	odhIntegrationList, err := m.odhintegrationClient.ODHIntegration(constants.TEMPLATE_NAMESPACE).List(metav1.ListOptions{})
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Fail to get ODHIntegration List: %v", err))
		return &msv1.GetAppResponse{}, err
	}

	for _, odhIntegration := range odhIntegrationList.Items {

		app := &msv1.Application{
			Name:                odhIntegration.Name,
			CsvName:             odhIntegration.Spec.CsvName,
			Provider:            odhIntegration.Spec.ProviderName,
			Template:            odhIntegration.Spec.TemplateName,
			MinSupportedVersion: odhIntegration.Spec.MinSupportedVersion,
			Installed:           true,
		}

		csvList, err := m.operatorClient.ClusterServiceVersions("default").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			logger.Log.Error(err.Error())
		}

		for _, csv := range csvList.Items {
			if strings.Contains(csv.Name, odhIntegration.Spec.CsvName) {
				csvVersion := semver.New(csv.Spec.Version.String())
				minSupportedVersion := semver.New(odhIntegration.Spec.MinSupportedVersion)
				logger.Log.Debug(fmt.Sprintf("csvName: %s csvVersion: %s, minSupportedVersion: %s", csv.Name, csvVersion, minSupportedVersion))

				if csvVersion.LessThan(*minSupportedVersion) {
					logger.Log.Info(fmt.Sprintf("CSV(%s) version(%s) installed on the cluster is not supported version: %s", csv.Name, csvVersion, minSupportedVersion))
					//TODO isv supportability
				}
			}
		}

		appList = append(appList, app)
		logger.Log.Debug(fmt.Sprintf("modelserving cr is added: %s", app))
	}

	return &msv1.GetAppResponse{Applications: appList}, nil
}

func (m *modelServingServer) GetAppParams(ctx context.Context, req *msv1.GetAppParamsRequest) (*msv1.GetAppParamsResponse, error) {
	logger.Log.Debug("Entry modelserving.go - GetAppParams")

	odhIntegration, err := m.odhintegrationClient.ODHIntegration(constants.TEMPLATE_NAMESPACE).Get(req.AppName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		logger.Log.Error(fmt.Sprintf("ODHIntegration(%s) is not found: %v", req.AppName, err))
		return &msv1.GetAppParamsResponse{}, err
	}

	template, err := m.templateClient.Templates(constants.TEMPLATE_NAMESPACE).Get(context.Background(), odhIntegration.Spec.TemplateName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		logger.Log.Error(fmt.Sprintf("Template(%s) is not found: '%v' ", odhIntegration.Spec.TemplateName, err))
		return &msv1.GetAppParamsResponse{}, err
	}

	responseParameters := []templatev1.Parameter{}

	for _, param := range template.Parameters {
		isStorageParameter := false

		for _, storageType := range storage.StorageTypes {
			if strings.Contains(param.Name, storageType) {
				logger.Log.Debug(fmt.Sprintf("Parameter(%s) is for storage(%s) so it will be skipped", param.Name, storageType))
				isStorageParameter = true
				continue
			}
		}
		if !isStorageParameter {
			logger.Log.Debug(fmt.Sprintf("Parameter(%s) is added to responseParameters", param.Name))
			responseParameters = append(responseParameters, param)
		}

	}

	parameterBytes, err := json.Marshal(responseParameters)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("json marshal failed: %v", err))
		return &msv1.GetAppParamsResponse{}, err
	}

	getAppParamResponse := &msv1.GetAppParamsResponse{
		AppName:     req.AppName,
		StorageName: req.StorageName,
		Namespace:   req.Namespace,
		Parameters:  parameterBytes,
	}
	logger.Log.Debug(fmt.Sprintf("getAppParamResponse: %s", utils.ProtobufToJson(getAppParamResponse)))

	return getAppParamResponse, nil
}

func (m *modelServingServer) GetAppCustomResource(ctx context.Context, req *msv1.GetRenderedCRRequest) (*msv1.GetRenderedCRResponse, error) {
	logger.Log.Debug("Entry modelserving.go - GetAppCustomResource")

	odhIntegration, err := m.odhintegrationClient.ODHIntegration(constants.TEMPLATE_NAMESPACE).Get(req.AppName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		logger.Log.Error(fmt.Sprintf("ODHIntegration(%s) is not found: %v", req.AppName, err))
		return &msv1.GetRenderedCRResponse{}, err
	}

	template, err := m.templateClient.Templates(constants.TEMPLATE_NAMESPACE).Get(context.Background(), odhIntegration.Spec.TemplateName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		logger.Log.Error(fmt.Sprintf("Template(%s) is not found: '%v' ", odhIntegration.Spec.TemplateName, err))
		return &msv1.GetRenderedCRResponse{}, err
	}

	secret, err := m.clientset.CoreV1().Secrets(req.Namespace).Get(context.Background(), req.StorageName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		logger.Log.Error(fmt.Sprintf("Secret(%s) in namespace(%s) is not found: '%v' ", req.StorageName, req.Namespace, err))
		return &msv1.GetRenderedCRResponse{}, err
	}

	for secretDataKey, secretDataValue := range secret.Data {
		logger.Log.Debug(fmt.Sprintf("SecretKey:%s", secretDataKey))
		logger.Log.Debug(fmt.Sprintf("SecretValue:%s", secretDataValue))
		req.Parameters[secretDataKey] = string(secretDataValue)
	}

	jsonString, err := processString(m.Scheme, req.Parameters, m.templateClient, template)
	if err != nil {
		return &msv1.GetRenderedCRResponse{}, err
	}
	return &msv1.GetRenderedCRResponse{Manifest: []byte(jsonString)}, nil

}

func processString(scheme *runtime.Scheme, values map[string]string, templateClient *templatev1client.TemplateV1Client, in *templatev1.Template) (string, error) {
	logger.Log.Debug("Entry modelserving.go - processString")
	logger.Log.Debug(fmt.Sprintf("Template(%s) Parameters : %v", in.Name, values))

	resultObj, err := utils.Process(values, templateClient, in)
	if err != nil {
		return "", err
	}

	jsonString := string(resultObj.Objects[0].Raw)
	return jsonString, nil
}
