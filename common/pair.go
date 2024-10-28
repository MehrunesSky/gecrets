package common

type Secret struct {
	Key         string
	Value       string
	ContentType string
}

func NewSecret(key string, value string, contentType string) Secret {
	return Secret{Key: key, Value: value, ContentType: contentType}
}
