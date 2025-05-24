package etagproviders_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/thtg88/blog.marco-marassi.com/pkg/etagproviders"
)

const uuidStringTestValue = "0e02e434-e0b4-4ac9-9d2d-329d456325e4"

// ignore the error returned value as we parse a constant value
var uuidTestValue, _ = uuid.Parse(uuidStringTestValue)

func TestUUIDProvider_GetETag(t *testing.T) {
	t.Parallel()

	provider := etagproviders.NewUUIDProvider(uuidTestValue)

	etag := provider.GetETag()

	assert.Equal(t, uuidStringTestValue, etag)
}

func TestUUIDProvider_GetName(t *testing.T) {
	t.Parallel()

	provider := etagproviders.NewUUIDProvider(uuidTestValue)

	name := provider.GetName()

	assert.Equal(t, etagproviders.UUIDProviderName, name)
}

func TestUUIDProvider_IsSupported(t *testing.T) {
	t.Parallel()

	provider := etagproviders.NewUUIDProvider(uuidTestValue)

	isSupported := provider.IsSupported()

	assert.True(t, isSupported)
}
