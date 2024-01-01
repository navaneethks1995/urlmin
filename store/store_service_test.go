package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStorage()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	originalUrl := "https://github.com/navaneethks1995/urlmin"
	userId := "as9d8a7hd9asyd9"
	minUrl := "minified_url"

	SaveUrlMapping(originalUrl, minUrl, userId)

	fetchedUrl := GetOriginalUrl(minUrl)

	assert.Equal(t, fetchedUrl, originalUrl)
}
