package etagproviders_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thtg88/blog.marco-marassi.com/pkg/etagproviders"
)

const (
	herokuBuildCommitTestValue    = "build-commit"
	herokuReleaseVersionTestValue = "release-version"
)

func TestHerokuProvider_GetETag(t *testing.T) {
	t.Parallel()

	provider := etagproviders.NewHerokuProvider(
		herokuBuildCommitTestValue,
		herokuReleaseVersionTestValue,
		etagproviders.HerokuProviderEnvironmentVariablesVariantName,
	)

	etag := provider.GetETag()

	assert.Equal(t, "build-commit-release-version", etag)
}

func TestHerokuProvider_GetName(t *testing.T) {
	t.Parallel()

	provider := etagproviders.NewHerokuProvider(
		herokuBuildCommitTestValue,
		herokuReleaseVersionTestValue,
		etagproviders.HerokuProviderEnvironmentVariablesVariantName,
	)

	name := provider.GetName()

	assert.Equal(t, etagproviders.HerokuProviderName, name)
}

func TestHerokuProvider_IsSupported(t *testing.T) {
	type test struct {
		description         string
		buildCommit         string
		releaseVersion      string
		expectedIsSupported bool
	}

	tests := []test{
		{
			description:         "supported provider",
			buildCommit:         herokuBuildCommitTestValue,
			releaseVersion:      herokuReleaseVersionTestValue,
			expectedIsSupported: true,
		},
		{
			description:         "not supported without a build commit",
			buildCommit:         "",
			releaseVersion:      herokuReleaseVersionTestValue,
			expectedIsSupported: false,
		},
		{
			description:         "not supported without a release version",
			buildCommit:         herokuBuildCommitTestValue,
			releaseVersion:      "",
			expectedIsSupported: false,
		},
		{
			description:         "not supported without a build commit nor a release version",
			buildCommit:         "",
			releaseVersion:      "",
			expectedIsSupported: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			provider := etagproviders.NewHerokuProvider(
				tc.buildCommit,
				tc.releaseVersion,
				etagproviders.HerokuProviderEnvironmentVariablesVariantName,
			)

			isSupported := provider.IsSupported()

			assert.Equal(t, tc.expectedIsSupported, isSupported)
		})
	}

}
