package mapper_test

import (
	"github.com/assertis/proxtasy/jp/atomised"
	"github.com/assertis/url-mapper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMappingStrings(t *testing.T) {
	var str = atomised.SearchRequest{}

	err := mapper.Unmarshal("origin=TBW&destination=LBG&adults=1&children=0&outward=495843958", &str)

	assert.Nil(t, err)
	assert.Equal(t, "TBW", str.Origin)
	assert.Equal(t, "LBG", str.Destination)
	assert.Equal(t, time.Unix(495843958, 0), str.Outward)
	assert.Equal(t, time.Time{}, str.Inward)
	assert.Equal(t, 1, str.Adults)
	assert.Equal(t, 0, str.Children)
}
