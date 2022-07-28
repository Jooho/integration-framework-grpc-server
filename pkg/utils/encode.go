package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/Jooho/integration-framework-server/pkg/logger"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/runtime"
	k8sJson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

func ProtobufToJson(pb proto.Message) string {
	marshaler := protojson.MarshalOptions{
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
		Multiline: true,
	}
	j, err := marshaler.Marshal(pb)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Can't convert Protobuf to JSON: %s", err))
	}
	// the return []byte type to string
	return string(j)
}

func JsonToProtobuf(json string, pb proto.Message) {
	// Unmarshal accept []byte, converting json string to []byte
	err := protojson.Unmarshal([]byte(json), pb)
	if err != nil {
		log.Println("Can't convert from JSON to Protobuf", err)
	}
}

// K8s Typed Object to JSON 
func JsonSerializer(obj runtime.Object, yamlType bool, pretty bool) []byte {
	
	encoder := k8sJson.NewSerializerWithOptions(
		k8sJson.DefaultMetaFactory, // jsonserializer.MetaFactory
		scheme.Scheme,              // runtime.ObjectCreater
		scheme.Scheme,              // runtime.ObjectTyper
		k8sJson.SerializerOptions{
			Yaml: yamlType,
			Pretty: pretty,
			Strict: false,
		},
	)
	encoded, err := runtime.Encode(encoder, obj)
	if err != nil {
		panic(err.Error())
	}
	return encoded
}

// JSON To K8s Typed Object
func JsonDeserializer(obj []byte) runtime.Object {

	decoder := k8sJson.NewSerializerWithOptions(
		k8sJson.DefaultMetaFactory, // jsonserializer.MetaFactory
		scheme.Scheme,              // runtime.Scheme implements runtime.ObjectCreater
		scheme.Scheme,              // runtime.Scheme implements runtime.ObjectTyper
		k8sJson.SerializerOptions{
			Yaml:   false,
			Pretty: false,
			Strict: false,
		},
	)

	decoded, err := runtime.Decode(decoder, obj)
	if err != nil {
		panic(err.Error())
	}
	return decoded
}


func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}
