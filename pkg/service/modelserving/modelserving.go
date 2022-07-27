package modelserving

import (
	"context"
	"fmt"
	"strings"

	v1ms "github.com/Jooho/integration-framework-server/pkg/api/v1/modelserving"
	odhitgv1alpha1client "github.com/Jooho/integration-framework-server/pkg/clientset/versioned/typed/odhintegration/v1alpha1"
	"github.com/Jooho/integration-framework-server/pkg/logger"
	templatev1client "github.com/openshift/client-go/template/clientset/versioned/typed/template/v1"
	operatorv1alpha1client "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned/typed/operators/v1alpha1"

	semver "github.com/coreos/go-semver/semver"
	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type modelServingServer struct {
	v1ms.ModelServingServer
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
	v1ms.RegisterModelServingServer(&s, modelserving)
}

func (m *modelServingServer) ListApp(ctx context.Context, req *v1ms.GetAppRequest) (*v1ms.GetAppResponse, error) {
	logger.Log.Debug("Entry modelserving.go - GetApplications")

	appList := []*v1ms.Application{}
	odhIntegrationList, err := m.odhintegrationClient.ODHIntegration(req.Namespace).List(metav1.ListOptions{})
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Fail to get ODHIntegration List: %v", err))
		return &v1ms.GetAppResponse{}, err
	}

	for _, odhIntegration := range odhIntegrationList.Items {

		app := &v1ms.Application{
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
					logger.Log.Info(fmt.Sprintf("CSV(%s) version(%s) installed on the cluster is not supported version: %s",csv.Name, csvVersion, minSupportedVersion))
					//TODO isv supportability
				}
			}
		}

		appList = append(appList, app)
		logger.Log.Debug(fmt.Sprintf("modelserving cr is added: %s", app))
	}

	return &v1ms.GetAppResponse{Applications: appList}, nil
}

func (m *modelServingServer) GetAppParams(ctx context.Context, req *v1ms.GetAppParamsRequest) (*v1ms.GetAppParamsResponse, error) {
	logger.Log.Debug("Entry modelserving.go - GetAppParams")

	return &v1ms.GetAppParamsResponse{}, nil
}

func (m *modelServingServer) GetAppCustomResource(ctx context.Context, req *v1ms.GetAppCRRequest) (*v1ms.GetAppCRResponse, error) {
	logger.Log.Debug("Entry modelserving.go - GetAppCustomResource")

	return &v1ms.GetAppCRResponse{}, nil
}
