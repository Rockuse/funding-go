package common

import "github.com/segmentio/ksuid"

func Uuid() int {
	id := int(ksuid.New().Timestamp())
	return id
}
