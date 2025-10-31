package storage

import (
	"fmt"
	"os"
)

func NewStorageFromEnv() (Storage, error) {
	storageType := os.Getenv("STORAGE_TYPE")

	switch storageType {
	case "locla":
		return NewLocalStorage(
			os.Getenv("LOCAL_PATH"),
			os.Getenv("LOCAL_URL"),
		), nil

	case "s3":
		return nil, fmt.Errorf("s3 not implemented")

	default:
		return nil, fmt.Errorf("unknown storage type: %s", storageType)
	}
}
