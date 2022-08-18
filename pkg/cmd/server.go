package cmd

import (
	"context"
	"flag"
	"fmt"

	odhintegrationv1alpha1 "github.com/Jooho/integration-framework-server/pkg/api/odhintegration/v1alpha1"
	"github.com/Jooho/integration-framework-server/pkg/helpers/legacy"
	"github.com/Jooho/integration-framework-server/pkg/logger"
	"github.com/Jooho/integration-framework-server/pkg/protocol/grpc"
	"github.com/Jooho/integration-framework-server/pkg/protocol/rest"
	"github.com/Jooho/integration-framework-server/pkg/utils"
	templatev1 "github.com/openshift/api/template/v1"
	operatorv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	apiruntime "k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var (
	scheme   = apiruntime.NewScheme()
)

func init() {
}

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP server start parameters section
	// HTTPPort is http port to listen by http server
	HTTPPort string


	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string

	//Set running place to get a right kuberentes config
	Mode string

	//Set running environment to change log 
	Env string
}

func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.Mode, "mode", "cluster", "kubernetes config path: cluster, local")
	flag.StringVar(&cfg.Env, "env", "dev", "environment type: dev, prod")
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "9000", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "8000", "http port to bind")
	flag.IntVar(&cfg.LogLevel, "log-level", 0, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "",
		"Print time format for logger e.g. 2006-01-02T15:04:05Z07:00")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		custom_error := fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
		logger.Log.Error(custom_error.Error())
		return custom_error
	}

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		custom_error := fmt.Errorf("failed to initialize logger: %v", err)
		logger.Log.Error(custom_error.Error())
		return custom_error
	}
	
	// get k8s clientset
	clientset, err := utils.GetK8SClientSet(cfg.Mode)
	if err != nil {
		custom_error := fmt.Errorf("failed to initialize a connection to kuberenetes: %v", err)
		logger.Log.Error(custom_error.Error())
		return custom_error
	}

	// get k8s restconfig
	restconfig, err := utils.GetK8SRestConfig(cfg.Mode)
	if err != nil {
		custom_error := fmt.Errorf("failed to initialize a connection to kuberenetes: %v", err)
		logger.Log.Error(custom_error.Error())
		return custom_error
	}

	//Add 3rd API Scheme
	utilruntime.Must(templatev1.Install(scheme))
	utilruntime.Must(odhintegrationv1alpha1.Install(scheme))
	utilruntime.Must(operatorv1alpha1.AddToScheme(scheme))

	legacy.InstallExternalLegacyTemplate(scheme)
	
	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, cfg.GRPCPort, scheme, clientset, restconfig)

}
