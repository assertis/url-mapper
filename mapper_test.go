package mapper_test

import (
	"github.com/assertis/url-mapper"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

type TestRequest struct {
	Origin          string    `query:"o"`
	Destination     string    `query:"d"`
	NumOfPassengers int       `query:"pax"`
	OutwardDate     time.Time `query:"outward_date"`
	ReturnDate      time.Time `query:"return_date,omitempty"`
}

func TestMappingStrings(t *testing.T) {
	var r = TestRequest{}

	values, _ := url.ParseQuery("o=TBW&d=LBG&pax=1&outward_date=1482852746")
	err := mapper.Unmarshal(values, &r)

	assert.Nil(t, err)
	assert.Equal(t, "TBW", r.Origin)
	assert.Equal(t, "LBG", r.Destination)
	assert.Equal(t, time.Unix(1482852746, 0), r.OutwardDate)
	assert.Equal(t, time.Time{}, r.ReturnDate)
	assert.Equal(t, 1, r.NumOfPassengers)
}
