package utils

import (
	"log"
	"strconv"
)

func ConvertStringToUint(id string) uint {
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("Error converting string to uint: %v", err)
		return 0
	}
	return uint(uintID)
}
