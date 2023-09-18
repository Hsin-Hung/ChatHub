package utils

import (
	"time"

	"github.com/google/uuid"
)

func Generate_id() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func Generate_timestamp() int64 {
	return time.Now().Unix()
}
