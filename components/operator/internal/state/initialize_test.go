package state

import (
	"context"
	"testing"

	"github.com/kyma-project/docker-registry/components/operator/api/v1alpha1"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func Test_sFnInitialize(t *testing.T) {
	t.Run("setup and return next step sFnRegistryConfiguration", func(t *testing.T) {
		r := &reconciler{
			cfg: cfg{
				finalizer: v1alpha1.Finalizer,
			},
			k8s: k8s{
				client: fake.NewClientBuilder().Build(),
			},
		}
		s := &systemState{
			instance: v1alpha1.DockerRegistry{
				ObjectMeta: metav1.ObjectMeta{
					Finalizers: []string{
						r.cfg.finalizer,
					},
				},
			},
		}

		// setup and return buildSFnPrerequisites
		next, result, err := sFnInitialize(context.Background(), r, s)
		require.Nil(t, err)
		require.Nil(t, result)
		requireEqualFunc(t, sFnAccessConfiguration, next)

		require.Equal(t, v1alpha1.StateProcessing, s.instance.Status.State)
	})

	t.Run("setup and return next step sFnDeleteResources", func(t *testing.T) {
		r := &reconciler{
			cfg: cfg{
				finalizer: v1alpha1.Finalizer,
			},
			k8s: k8s{
				client: fake.NewClientBuilder().Build(),
			},
		}
		metaTime := metav1.Now()
		s := &systemState{
			instance: v1alpha1.DockerRegistry{
				ObjectMeta: metav1.ObjectMeta{
					Finalizers: []string{
						r.cfg.finalizer,
					},
					DeletionTimestamp: &metaTime,
				},
			},
		}

		// setup and return buildSFnDeleteResources
		next, result, err := sFnInitialize(context.Background(), r, s)
		require.Nil(t, err)
		require.Nil(t, result)
		requireEqualFunc(t, sFnDeleteResources, next)

		require.Equal(t, v1alpha1.StateProcessing, s.instance.Status.State)
	})
}
