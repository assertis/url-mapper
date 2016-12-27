package mapper_test

import (
	"github.com/assertis/url-mapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestRequest struct {
	Origin      string `query:"origin,regexp=^[A-Z]{3}$"`
	Destination string `query:"destination"`
}

func TestMappingStrings(t *testing.T) {
	var str = TestRequest{}

	err := mapper.Unmarshal("origin=TBW&destination=LBG", &str)

	assert.Nil(t, err)
	assert.Equal(t, "TBW", str.Origin)
	assert.Equal(t, "LBG", str.Destination)
}
