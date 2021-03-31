package configuration

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// getOrCreateConfigSecret finds or creates the secret which houses the combined configurations of the cluster
// for eventual parsing and emitting to the Kong Admin API on the proxy instances.
func getOrCreateConfigSecret(ctx context.Context, c client.Client, targetNsn types.NamespacedName) (*corev1.Secret, bool, error) {
	secret := new(corev1.Secret)
	if err := c.Get(ctx, targetNsn, secret); err != nil {
		if errors.IsNotFound(err) {
			secret.SetName(targetNsn.Name)
			secret.SetNamespace(targetNsn.Namespace)
			if err := c.Create(ctx, secret); err != nil {
				return nil, false, err
			}
			return secret, true, nil
		} else {
			return nil, false, err
		}
	}
	if secret.Data == nil {
		secret.Data = make(map[string][]byte)
	}
	return secret, false, nil
}