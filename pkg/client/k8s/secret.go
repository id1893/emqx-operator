package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type SecretManagers interface {
	GetSecret(namespace, name string) (*corev1.Secret, error)
	CreateSecret(object *corev1.Secret) error
	UpdateSecret(object *corev1.Secret) error
	DeleteSecret(namespace, name string) error
}

// GetSecret implement the Service.Interface
func (manager *Manager) GetSecret(namespace, name string) (*corev1.Secret, error) {
	object := &corev1.Secret{}
	err := manager.Client.Get(
		context.TODO(),
		types.NamespacedName{
			Name:      name,
			Namespace: namespace,
		}, object,
	)

	if err != nil {
		return nil, err
	}
	return object, err
}

func (manager *Manager) CreateSecret(object *corev1.Secret) error {
	err := manager.Client.Create(context.TODO(), object)
	if err != nil {
		return err
	}
	manager.Logger.WithValues(
		"kind", object.Kind,
		"apiVersion", object.APIVersion,
		"namespace", object.Namespace,
		"name", object.Name,
	).Info("Create successfully")
	return nil
}

func (manager *Manager) UpdateSecret(object *corev1.Secret) error {
	if err := manager.Client.Update(context.TODO(), object); err != nil {
		return err
	}
	manager.Logger.WithValues(
		"kind", object.Kind,
		"apiVersion", object.APIVersion,
		"namespace", object.Namespace,
		"name", object.Name,
	).Info("Update successfully")
	return nil
}

func (manager *Manager) DeleteSecret(namespace string, name string) error {
	object := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	return manager.Client.Delete(context.TODO(), object)
}
