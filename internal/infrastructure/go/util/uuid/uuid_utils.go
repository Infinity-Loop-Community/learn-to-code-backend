package uuid

import "github.com/google/uuid"

func MustNewRandomAsString() string {
	newUuid, err := uuid.NewRandom()

	if err != nil {
		panic(err)
	}

	return newUuid.String()
}
