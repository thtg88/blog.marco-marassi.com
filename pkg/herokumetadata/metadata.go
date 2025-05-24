package herokumetadata

import "fmt"

// HerokuMetadata contains metadata about the dyno, app, and release the application is running on Heroku.
// For the purposes of this repository, we are only interested in release data.
// But the whole file (at `/etc/heroku/dyno`) will look like this if you ever need other data.
// Please note this file will only be available if the `runtime-dyno-build-metadata` Heroku Labs feature is enabled.
// This file is not really documented so it may break any time, but you can see the docs on how to enable this Labs feature at:
// https://devcenter.heroku.com/articles/dyno-metadata
//
//	{
//	  "dyno": {
//	    "id": "3313268c-2e27-45f1-939c-3ba0e6ba56bc",
//	    "name": "run.6500"
//	  },
//	  "app": {
//	    "id": "31a2b7a3-3049-42a0-a67a-6c930937b1e5",
//	    "name": ""
//	  },
//	  "release": {
//	    "id": 18,
//	    "commit": "46f55628e17e2c18d5f336a4a12247f82e7af087",
//	    "description": "Deploy 46f55628"
//	  }
//	}
type Metadata struct {
	Release ReleaseMetadata `json:"release"`
}

type ReleaseMetadata struct {
	ID          uint64 `json:"id"`
	Commit      string `json:"commit"`
	Description string `json:"description"`
}

// Version returns the ReleaseMetadata's ID preceded by `v`.
// Heroku always returns the current version in the ID as an integer.
func (rm ReleaseMetadata) Version() string {
	return fmt.Sprintf("v%d", rm.ID)
}
