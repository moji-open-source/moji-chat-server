package main

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestUUID(t *testing.T) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("uuid = ", uuid.String())
}
