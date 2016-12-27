package mapper_test

import (
	"github.com/assertis/proxtasy/jp/atomised"
	"github.com/assertis/url-mapper"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

func TestMappingStrings(t *testing.T) {
	var str = atomised.SearchRequest{}

	values, _ := url.ParseQuery("origin=TBW&destination=LBG&adults=1&children=0&outward=1482852746")
	err := mapper.Unmarshal(values, &str)

	assert.Nil(t, err)
	assert.Equal(t, "TBW", str.Origin)
	assert.Equal(t, "LBG", str.Destination)
	assert.Equal(t, time.Unix(1482852746, 0), str.Outward)
	assert.Equal(t, time.Time{}, str.Inward)
	assert.Equal(t, int64(1), str.Adults)
	assert.Equal(t, int64(0), str.Children)
}
