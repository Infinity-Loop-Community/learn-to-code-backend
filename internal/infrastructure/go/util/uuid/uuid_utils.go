package uuid

import "github.com/google/uuid"

func MustNewRandomAsString() string {
	newUUID, err := uuid.NewRandom()

	if err != nil {
		panic(err)
	}

	return newUUID.String()
}
