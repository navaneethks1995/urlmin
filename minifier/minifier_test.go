package minifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const user_id = "asdasdasdasd"

func TestMinifier(t *testing.T) {
	url1 := "https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go"
	url2 := "https://gojs.net/latest/samples/leaflet.html"

	minurl1 := GenreateMinUrl(url1, user_id)
	minurl2 := GenreateMinUrl(url2, user_id)

	assert.NotEqual(t, minurl1, minurl2)
}
