package v1alpha1

import (
	"context"

	odhintegrationv1alpha1 "github.com/Jooho/integration-framework-server/pkg/api/odhintegration/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type ODHIntegrationGetter interface {
	ODHIntegration(namespace string) ODHIntegrationInterface
}

type odhIntegration struct {
	restClient rest.Interface
	ns         string
}

// newODHIntegration returns a Templates
func newODHIntegration(c *ODHIntegrationV1Alpha1Client, namespace string) *odhIntegration {
	return &odhIntegration{
		restClient: c.RESTClient(),
		ns:         namespace,
	}
}

type ODHIntegrationInterface interface {
	List(opts metav1.ListOptions) (*odhintegrationv1alpha1.ODHIntegrationList, error)
	Get(name string, options metav1.GetOptions) (*odhintegrationv1alpha1.ODHIntegration, error)
}

func (o *odhIntegration) List(opts metav1.ListOptions) (result *odhintegrationv1alpha1.ODHIntegrationList, err error) {
	result = &odhintegrationv1alpha1.ODHIntegrationList{}
	err = o.restClient.
		Get().
		Namespace(o.ns).
		Resource("odhintegrations").
		VersionedParams(&opts, odhintegrationv1alpha1.ParameterCodec).
		Do(context.Background()).
		Into(result)
		
	return 
}

func (o *odhIntegration) Get(name string, opts metav1.GetOptions) (*odhintegrationv1alpha1.ODHIntegration, error) {
	result := odhintegrationv1alpha1.ODHIntegration{}
	err := o.restClient.
		Get().
		Namespace(o.ns).
		Resource("odhintegrations").
		Name(name).
		VersionedParams(&opts, odhintegrationv1alpha1.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}
