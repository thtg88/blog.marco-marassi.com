package etagproviders_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/thtg88/blog.marco-marassi.com/pkg/etagproviders"
	"github.com/thtg88/blog.marco-marassi.com/pkg/herokumetadata"
)

func TestBuildProvider(t *testing.T) {
	type test struct {
		description              string
		herokuMetadata           herokumetadata.Metadata
		herokuBuildCommit        string
		herokuReleaseVerion      string
		id                       uuid.UUID
		expectedETagProviderName string
	}

	tests := []test{
		{
			description: "heroku etag provider from metadata",
			herokuMetadata: herokumetadata.Metadata{
				Release: herokumetadata.ReleaseMetadata{
					ID:     1,
					Commit: herokuBuildCommitTestValue,
				},
			},
			herokuBuildCommit:        herokuBuildCommitTestValue,
			herokuReleaseVerion:      herokuReleaseVersionTestValue,
			id:                       uuidTestValue,
			expectedETagProviderName: fmt.Sprintf("%s-%s", etagproviders.HerokuProviderName, etagproviders.HerokuProviderMetadataFileVariantName),
		},
		{
			description:              "heroku etag provider from env variables",
			herokuBuildCommit:        herokuBuildCommitTestValue,
			herokuReleaseVerion:      herokuReleaseVersionTestValue,
			id:                       uuidTestValue,
			expectedETagProviderName: fmt.Sprintf("%s-%s", etagproviders.HerokuProviderName, etagproviders.HerokuProviderEnvironmentVariablesVariantName),
		},
		{
			description:              "uuid etag provider",
			herokuBuildCommit:        "",
			herokuReleaseVerion:      herokuReleaseVersionTestValue,
			id:                       uuidTestValue,
			expectedETagProviderName: string(etagproviders.UUIDProviderName),
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			provider := etagproviders.BuildProvider(tc.herokuMetadata, tc.herokuBuildCommit, tc.herokuReleaseVerion, tc.id)

			assert.Equal(t, tc.expectedETagProviderName, provider.GetName())
		})
	}
}
