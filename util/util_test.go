package util

import (
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
