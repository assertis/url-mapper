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
	OutwardDate     time.Time `query:"outward_date,rfc3339"`
	ReturnDate      time.Time `query:"return_date,unix"`
}

func TestMappingStrings(t *testing.T) {
	var r = TestRequest{}

	values, err := url.ParseQuery("o=TBW&d=LBG&pax=1&outward_date=1482852746&return_date=2016-12-31T11:00:00Z")
	assert.Nil(t, err)

	err = mapper.Unmarshal(values, &r)
	assert.Nil(t, err)

	rtnDate, err := time.Parse(time.RFC3339, "2016-12-31T11:00:00Z")
	assert.Nil(t, err)

	assert.Equal(t, "TBW", r.Origin)
	assert.Equal(t, "LBG", r.Destination)
	assert.Equal(t, time.Unix(1482852746, 0), r.OutwardDate)
	assert.Equal(t, rtnDate, r.ReturnDate)
	assert.Equal(t, 1, r.NumOfPassengers)
}
