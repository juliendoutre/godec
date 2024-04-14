package url_test

import (
	"testing"

	"github.com/juliendoutre/godec/examples/url"
	"github.com/stretchr/testify/assert"
)

func TestSchemeValidHTTPURLDecoding(t *testing.T) {
	var scheme string
	var username string
	var password string
	var host string
	var port uint
	var path string
	var query string
	var fragment string

	decoder := url.Codec(&scheme, &username, &password, &host, &port, &path, &query, &fragment)
	remainder, err := decoder.Decode([]byte("http://google.com:443/test?page=1#title"))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(remainder))
	assert.Equal(t, "http", scheme)
	assert.Equal(t, "", username)
	assert.Equal(t, "", password)
	assert.Equal(t, "google.com", host)
	assert.Equal(t, uint(443), port)
	assert.Equal(t, "/test", path)
	assert.Equal(t, "page=1", query)
	assert.Equal(t, "title", fragment)
}

func TestSchemeValidPostgresDecoding(t *testing.T) {
	var scheme string
	var username string
	var password string
	var host string
	var port uint
	var path string
	var query string
	var fragment string

	decoder := url.Codec(&scheme, &username, &password, &host, &port, &path, &query, &fragment)
	remainder, err := decoder.Decode([]byte("postgres://user:password@localhost:5432/database"))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(remainder))
	assert.Equal(t, "postgres", scheme)
	assert.Equal(t, "user", username)
	assert.Equal(t, "password", password)
	assert.Equal(t, "localhost", host)
	assert.Equal(t, uint(5432), port)
	assert.Equal(t, "/database", path)
	assert.Equal(t, "", query)
	assert.Equal(t, "", fragment)
}
