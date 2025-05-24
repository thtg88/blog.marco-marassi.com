package etagproviders

import "github.com/google/uuid"

const UUIDProviderName ETagProviderName = "uuid"

// UUIDProvider is a UUID v4 ETag provider
type UUIDProvider struct {
	id uuid.UUID
}

func NewUUIDProvider(id uuid.UUID) *UUIDProvider {
	return &UUIDProvider{id: id}
}

func (up *UUIDProvider) GetETag() string {
	return up.id.String()
}

func (up *UUIDProvider) GetName() string {
	return string(UUIDProviderName)
}

// IsSupported returns whether the UUIDProvider is supported.
// As long as the code compiled, it means we can support using UUIDs.
func (up *UUIDProvider) IsSupported() bool {
	return true
}
