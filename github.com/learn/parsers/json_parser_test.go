package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSON(t *testing.T) {
	ParseJSON(`{
		"users": [
		  {
			"name": "Elliot",
			"type": "Reader",
			"age": 23,
			"social": {
			  "facebook": "https://facebook.com",
			  "twitter": "https://twitter.com"
			}
		  },
		  {
			"name": "Fraser",
			"type": "Author",
			"age": 17,
			"social": {
			  "facebook": "https://facebook.com",
			  "twitter": "https://twitter.com"
			}
		  }
		]
		}`)

	assert := assert.New(t)
	assert.Equal(1+1, 2)
}
