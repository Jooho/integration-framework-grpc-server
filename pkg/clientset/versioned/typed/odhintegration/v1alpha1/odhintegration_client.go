package v1alpha1

import (
	"net/http"

	odhintegrationv1alpha1 "github.com/Jooho/integration-framework-server/pkg/api/odhintegration/v1alpha1"

	"k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type ODHIntegrationV1Alpha1Interface interface {
	RESTClient() rest.Interface
	ODHIntegrationGetter
}

type ODHIntegrationV1Alpha1Client struct {
	restClient rest.Interface
}

func (c *ODHIntegrationV1Alpha1Client) ODHIntegration(namespace string) ODHIntegrationInterface {
	return newODHIntegration(c, namespace)
}

func NewForConfig(c *rest.Config) (*ODHIntegrationV1Alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

func setConfigDefaults(config *rest.Config) error {
	gv := odhintegrationv1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// NewForConfigAndClient creates a new ODHIntegrationV1Alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ODHIntegrationV1Alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ODHIntegrationV1Alpha1Client{client}, nil
}

// New creates a new ODHIntegrationV1Alpha1Client for the given RESTClient.
func New(c rest.Interface) *ODHIntegrationV1Alpha1Client {
	return &ODHIntegrationV1Alpha1Client{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ODHIntegrationV1Alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
