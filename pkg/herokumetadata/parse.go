package herokumetadata

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

const DefaultMetadataFilename string = "/etc/heroku/dyno"

var ErrCouldNotOpenFile = errors.New("could not open file")
var ErrCouldNotReadAllFile = errors.New("could not readall file")
var ErrCouldNotUnmarshalBytes = errors.New("could not unmarshal bytes")

// Parse open the filename provided from disk, and JSON unmarshals it in the provided metadata
// Failure to either open, read, or unmarshal the file will be returned as an error
func Parse(filename string, metadata *Metadata) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Join(ErrCouldNotOpenFile, err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return errors.Join(ErrCouldNotReadAllFile, err)
	}

	err = json.Unmarshal(bytes, metadata)
	if err != nil {
		return errors.Join(ErrCouldNotUnmarshalBytes, err)
	}

	return nil
}
