package common

type SecretI interface {
	GetKey() string
	ToJson() string
	Diff(otherSecret SecretI) bool
}
