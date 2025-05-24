package etagproviders

import (
	"github.com/google/uuid"

	"github.com/thtg88/blog.marco-marassi.com/pkg/herokumetadata"
)

// BuildProvider builds an ETagProvider from a given Heroku metadata, Heroku build commit and release version, and UUID.
// The function will try to build a provider in the order of the arguments provided. The first supported provider will be returned.
// This means that if we managed to successfully parse the Heroku metadata file, that data will be used.
// Afer, we will attempt with the environment variables.
// And finally, we will fallback to a UUID strategy.
func BuildProvider(
	herokuMetadata herokumetadata.Metadata,
	herokuBuildCommit string,
	herokuReleaseVerion string,
	id uuid.UUID,
) ETagProvider {
	herokuProvider := NewHerokuProvider(
		herokuMetadata.Release.Commit,
		herokuMetadata.Release.Version(),
		HerokuProviderMetadataFileVariantName,
	)
	if herokuProvider.IsSupported() {
		return herokuProvider
	}

	herokuProvider = NewHerokuProvider(
		herokuBuildCommit,
		herokuReleaseVerion,
		HerokuProviderEnvironmentVariablesVariantName,
	)
	if herokuProvider.IsSupported() {
		return herokuProvider
	}

	return NewUUIDProvider(id)
}
