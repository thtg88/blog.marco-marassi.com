package herokumetadata_test

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thtg88/blog.marco-marassi.com/pkg/herokumetadata"
)

func TestParse(t *testing.T) {
	type test struct {
		description      string
		createFile       bool
		fileContent      string
		expectedErr      error
		expectedMetadata herokumetadata.Metadata
	}

	tests := []test{
		{
			description:      "non-existent file",
			createFile:       false,
			expectedErr:      herokumetadata.ErrCouldNotOpenFile,
			expectedMetadata: herokumetadata.Metadata{},
		},
		{
			description:      "invalid json file",
			createFile:       true,
			fileContent:      "",
			expectedErr:      herokumetadata.ErrCouldNotUnmarshalBytes,
			expectedMetadata: herokumetadata.Metadata{},
		},
		{
			description:      "json file not containing metadata",
			createFile:       true,
			fileContent:      "{\"foo\": \"bar\"}",
			expectedErr:      nil,
			expectedMetadata: herokumetadata.Metadata{},
		},
		{
			description: "valid metadata file",
			createFile:  true,
			fileContent: `{
				"dyno": {
					"id": "3313268c-2e27-45f1-939c-3ba0e6ba56bc",
					"name": "run.6500"
				},
				"app": {
					"id": "31a2b7a3-3049-42a0-a67a-6c930937b1e5",
					"name": ""
				},
				"release": {
					"id": 18,
					"commit": "46f55628e17e2c18d5f336a4a12247f82e7af087",
					"description": "Deploy 46f55628"
				}
			}`,
			expectedErr: nil,
			expectedMetadata: herokumetadata.Metadata{
				Release: herokumetadata.ReleaseMetadata{
					ID:          18,
					Commit:      "46f55628e17e2c18d5f336a4a12247f82e7af087",
					Description: "Deploy 46f55628",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			var filename string
			if tc.createFile {
				var err error
				file, err := os.CreateTemp(t.TempDir(), "*")
				require.NoError(t, err)
				defer file.Close()

				filename = file.Name()

				n, err := file.WriteString(tc.fileContent)
				require.NoError(t, err)
				require.Equal(t, len(tc.fileContent), n)
			} else {
				filename = path.Join(t.TempDir(), "non-existent-file-name")
			}

			var metadata herokumetadata.Metadata
			err := herokumetadata.Parse(filename, &metadata)
			assert.ErrorIs(t, err, tc.expectedErr)
			assert.Equal(t, tc.expectedMetadata, metadata)
		})
	}
}
