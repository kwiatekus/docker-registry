package utils

import (
	"context"

	"github.com/kyma-project/docker-registry/components/operator/api/v1alpha1"
	"go.uber.org/zap"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type TestUtils struct {
	Ctx    context.Context
	Logger *zap.SugaredLogger
	Client client.Client

	Namespace                string
	Name                     string
	DockerregistryDeployName string
	RegistryName             string
	UpdateSpec               v1alpha1.DockerRegistrySpec
}
