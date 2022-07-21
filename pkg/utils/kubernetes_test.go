package utils

import (
	"strings"
	"testing"

	"github.com/openshift/client-go/template/clientset/versioned/scheme"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Test converting K8S object configmap to JSON string
func TestConvertK8StoJsonString(t *testing.T) {
	//given
	k8sObjConfigMap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		Data: map[string]string{"foo": "bar"},
	}
	k8sObjConfigMap.Namespace = "default"
	k8sObjConfigMap.Name = "my-configmap"

	expectedJsonString := "{\"kind\":\"ConfigMap\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"my-configmap\",\"namespace\":\"default\",\"creationTimestamp\":null},\"data\":{\"foo\":\"bar\"}}"

	//when - ConvertK8StoJsonString "K8S to JSON"
	testname := "Test converting K8S object configmap to JSON string"
	t.Run(testname, func(t *testing.T) {
		jsonString := ConvertK8StoJsonString(scheme.Scheme, k8sObjConfigMap, false, false)

		//then - jsonString must be the same as expectedJsonString
		if result, _ := AreEqualJSON(jsonString, expectedJsonString); !result {
			t.Errorf("got %s, want %s", jsonString, expectedJsonString)
		}
	})
}

func TestJsonStringToK8S(t *testing.T) {
	//given
	k8sObjConfigMap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		Data: map[string]string{"foo": "bar"},
	}
	k8sObjConfigMap.Namespace = "default"
	k8sObjConfigMap.Name = "my-configmap"

	JsonBytesConfigMap := []byte{123, 34, 107, 105, 110, 100, 34, 58, 34, 67, 111, 110, 102, 105, 103, 77, 97, 112, 34, 44, 34, 97, 112, 105, 86, 101, 114, 115, 105, 111, 110, 34, 58, 34, 118, 49, 34, 44, 34, 109, 101, 116, 97, 100, 97, 116, 97, 34, 58, 123, 34, 110, 97, 109, 101, 34, 58, 34, 109, 121, 45, 99, 111, 110, 102, 105, 103, 109, 97, 112, 34, 44, 34, 110, 97, 109, 101, 115, 112, 97, 99, 101, 34, 58, 34, 100, 101, 102, 97, 117, 108, 116, 34, 44, 34, 99, 114, 101, 97, 116, 105, 111, 110, 84, 105, 109, 101, 115, 116, 97, 109, 112, 34, 58, 110, 117, 108, 108, 125, 44, 34, 100, 97, 116, 97, 34, 58, 123, 34, 102, 111, 111, 34, 58, 34, 98, 97, 114, 34, 125, 125, 10}

	//when - ConvertK8StoJsonString "K8S to JSON"
	testname := "Test converting JSON bytes to K8S object configmap "
	t.Run(testname, func(t *testing.T) {
		configmapTest := (JsonDeserializer(JsonBytesConfigMap)).(*corev1.ConfigMap)
		//then - configmapTest name must be the same as original confimap object
		if strings.Compare(k8sObjConfigMap.Name, configmapTest.Name) != 0 {
			t.Errorf("got %s, want %s", k8sObjConfigMap.Name, configmapTest.Name)
		}
	})
}
