package mapper_test

import (
	"github.com/assertis/url-mapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTags(t *testing.T) {
	var expected mapper.TagOptions

	expected = make(map[string]string)
	name, opts := mapper.ParseTagsIntoMap("ab")
	assert.Equal(t, "ab", name)
	assert.Equal(t, expected, opts)

	expected = make(map[string]string)
	expected["cd"] = "ef"
	name, opts = mapper.ParseTagsIntoMap("ab,cd=ef")
	assert.Equal(t, "ab", name)
	assert.Equal(t, expected, opts)

	expected = make(map[string]string)
	expected["ef"] = ""
	expected["gh"] = ""
	expected["ij"] = "kl"
	name, opts = mapper.ParseTagsIntoMap("ab,ef,gh,ij=kl")
	assert.Equal(t, "ab", name)
	assert.Equal(t, expected, opts)
}
