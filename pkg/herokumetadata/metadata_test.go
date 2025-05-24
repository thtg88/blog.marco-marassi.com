package herokumetadata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thtg88/blog.marco-marassi.com/pkg/herokumetadata"
)

const releaseMetadataIDTestValue uint64 = 1

func TestReleaseMetadata_Version(t *testing.T) {
	releaseMetadata := herokumetadata.ReleaseMetadata{
		ID: releaseMetadataIDTestValue,
	}

	version := releaseMetadata.Version()

	assert.Equal(t, "v1", version)
}
