package util

import (
	"time"

	"github.com/segmentio/ksuid"
)

func Uuid() int {
	id := int(ksuid.New().Timestamp())
	return id
}
func CodeGenerator(tran string) (string, error) {
	id, err := ksuid.NewRandomWithTime(time.Now())
	if err != nil {
		return "error", err
	}
	return tran + "-" + id.Next().String(), nil
}
