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
		mockedError := errors.New("thats the error")

		recievedBooleanElement := ContainsError(mockedError)

		assert.True(t, recievedBooleanElement)
	})
}

func TestIsEqual(t *testing.T) {
	t.Run("must return true", func(t *testing.T) {
		mockedValue := "rodolfo"

		receivedResult := IsEqual(mockedValue, mockedValue)

		assert.True(t, receivedResult)
	})

	t.Run("must return false", func(t *testing.T) {
		mockedStringValue := "rodolfo"
		mockedIntValue := 1

		receivedResult := IsEqual(mockedStringValue, mockedIntValue)

		assert.False(t, receivedResult)
	})
}
