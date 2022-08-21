package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMapToJson(t *testing.T) {
	t.Run("ParseMapToJson --> must parse a string map", func(t *testing.T) {
		testMap := map[string]string{
			"rodolfo": "handsome",
			"skill":   "nice",
		}

		expectedJsonElement := `{"rodolfo":"handsome","skill":"nice"}`

		receivedJsonElement := ParseMapToJson(testMap)

		assert.Equal(t, expectedJsonElement, receivedJsonElement)

	})
}

func TestContainsError(t *testing.T) {
	t.Run("Must return true if error exists", func(t *testing.T) {
		err := errors.New("thats the error")

		recievedBooleanElement := ContainsError(err)

		assert.True(t, recievedBooleanElement)
	})
}
