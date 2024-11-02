package common

import "fmt"

type SecretI interface {
	GetKey() string
	ToJson() string
	Diff(otherSecret SecretI) bool
}

type SecretIs []SecretI

func (s SecretIs) MapByKey() map[string]SecretI {
	m := make(map[string]SecretI)

	for _, secret := range s {
		m[secret.GetKey()] = secret
	}

	return m
}

type Changed struct {
	Key      string
	OldValue string
	NewValue string
	New      bool
}

func NewChanged(key string, oldValue string, newValue string, new bool) Changed {
	return Changed{Key: key, OldValue: oldValue, NewValue: newValue, New: new}
}

func (s SecretIs) GetChangedSecrets(newSecrets SecretIs) (new []Changed, changed []Changed) {
	oldSecretsM, newSecretsM := s.MapByKey(), newSecrets.MapByKey()

	for k, secret := range newSecretsM {
		o, ok := oldSecretsM[k]
		if !ok {
			fmt.Println("Secrets", secret.GetKey(), "is new with this value :", secret)
			new = append(
				new,
				NewChanged(secret.GetKey(), "", secret.ToJson(), true),
			)
		} else if o.Diff(secret) {
			changed = append(
				changed,
				NewChanged(secret.GetKey(), o.ToJson(), secret.ToJson(), false),
			)
		}

	}
	return new, changed
}
