package lib_test

import (
	"github.com/assertis/url-mapper/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTags(t *testing.T) {
	var expected lib.TagOptions

	expected = make(map[string]string)
	expected["ab"] = ""
	assert.Equal(t, expected, lib.ParseTagsIntoMap("ab"))

	expected = make(map[string]string)
	expected["ab"] = "cd"
	assert.Equal(t, expected, lib.ParseTagsIntoMap("ab=cd"))

	expected = make(map[string]string)
	expected["ab"] = "cd"
	expected["ef"] = ""
	expected["gh"] = ""
	expected["ij"] = "kl"
	assert.Equal(t, expected, lib.ParseTagsIntoMap("ab=cd,ef,gh,ij=kl"))
}
