package mapper_test

import (
	"github.com/assertis/url-mapper/mapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalInputTypes(t *testing.T) {
	for _, testType := range []interface{}{
		0,
		0.0,
		true,
		[]interface{}{},
	} {
		assert.NotNil(t, mapper.Unmarshal(testType))
	}

	for _, testType := range []interface{}{
		"mapper",
		[]byte("mapper"),
	} {
		assert.Nil(t, mapper.Unmarshal(testType))
	}
}
