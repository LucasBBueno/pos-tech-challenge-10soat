package utils

import "github.com/google/uuid"

func StringToUuid(id string) uuid.UUID {
	parsedId, err := uuid.Parse(id)
	if err != nil {

	}
	return parsedId
}
