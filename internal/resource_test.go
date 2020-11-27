package internal_test

import (
	"testing"

	"github.com/sazzer/fakepi/internal"
	"github.com/stretchr/testify/assert"
)

func TestUnknownResource(t *testing.T) {
	_, err := internal.NewResource("./testResources/unknown")

	assert.EqualError(t, err, "Failed to open resource ./testResources/unknown: open ./testResources/unknown: no such file or directory")
}

func TestBlankResource(t *testing.T) {
	_, err := internal.NewResource("./testResources/blank")

	assert.EqualError(t, err, "Resource was empty")
}

func TestNonNumericStatusResource(t *testing.T) {
	_, err := internal.NewResource("./testResources/nonnumericStatus")

	assert.EqualError(t, err, "Failed to parse status code abc: strconv.Atoi: parsing \"abc\": invalid syntax")
}

func TestOnlyStatus(t *testing.T) {
	resource, err := internal.NewResource("./testResources/onlyStatus")

	assert.NoError(t, err)
	assert.Equal(t, resource.Status, 204)
	assert.Empty(t, resource.Headers)
	assert.Empty(t, resource.Body)
}

func TestStatusBody(t *testing.T) {
	resource, err := internal.NewResource("./testResources/statusBody")

	assert.NoError(t, err)
	assert.Equal(t, resource.Status, 200)
	assert.Empty(t, resource.Headers)
	assert.Equal(t, resource.Body, []byte("Hello\n"))
}

func TestStatusHeadersBody(t *testing.T) {
	resource, err := internal.NewResource("./testResources/headersBody")

	assert.NoError(t, err)
	assert.Equal(t, resource.Status, 200)
	assert.Equal(t, resource.Headers, []internal.Header{
		{Key: "Content-Type", Value: "application/json"},
		{Key: "Cache-Control", Value: "public, max-age=3600"},
	})
	assert.Equal(t, resource.Body, []byte("{\"hello\": \"world\"}\n"))
}
