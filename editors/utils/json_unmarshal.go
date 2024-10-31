package utils

import (
	"encoding/json"
	"github.com/MehrunesSky/gecrets/common"
	"log"
	"reflect"
)

func Unmarshal(model common.SecretI, value []byte) common.SecretI {
	t := reflect.TypeOf(model)
	copyModel := reflect.New(t.Elem()).Interface()
	err := json.Unmarshal(value, copyModel)
	if err != nil {
		log.Fatalln(err)
	}
	return reflect.Indirect(reflect.ValueOf(copyModel)).Interface().(common.SecretI)
}
