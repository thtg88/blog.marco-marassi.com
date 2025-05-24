package etagproviders

import "fmt"

const (
	HerokuProviderName                          ETagProviderName = "heroku"
	HerokuBuildCommitEnvironmentVariableName    string           = "HEROKU_BUILD_COMMIT"
	HerokuReleaseVersionEnvironmentVariableName string           = "HEROKU_RELEASE_VERSION"
)

type HerokuProviderVariantName string

const (
	HerokuProviderEnvironmentVariablesVariantName HerokuProviderVariantName = "environment_variables"
	HerokuProviderMetadataFileVariantName         HerokuProviderVariantName = "metadata_file"
)

type HerokuProvider struct {
	buildCommit   string
	releaseVerion string
	builtFrom     HerokuProviderVariantName
}

func NewHerokuProvider(buildCommit string, releaseVerion string, builtFrom HerokuProviderVariantName) *HerokuProvider {
	return &HerokuProvider{
		buildCommit:   buildCommit,
		releaseVerion: releaseVerion,
		builtFrom:     builtFrom,
	}
}

// GetETag returns the ETag for the provider.
// This will be in the form of Git commit SHA concatenated to the release version, e.g. `de75f99ac978cb09e1e4cdb993161fa6d46e86de-v19`.
// Append the release version, as a new release may alter the functionality of the app, without pushing a new commit.
func (hp *HerokuProvider) GetETag() string {
	return fmt.Sprintf("%s-%s", hp.buildCommit, hp.releaseVerion)
}

// GetName returns the name for the current provider.
// This can be used to identify within observability tools (logs, metrics, traces) which provider is being used.
// For Heroku, as we may be able to fetch the build metadata from either a metadata file or environment variables, this will be either `heroku-metadata_file` or `heroku-environment_variables`.
func (hp *HerokuProvider) GetName() string {
	return fmt.Sprintf("%s-%s", HerokuProviderName, hp.builtFrom)
}

// IsSupported returns whether the current provider is supported or not.
// For Heroku, as we build the ETag from both the build commit and release version, it means both these need to be set to correctly use the provider.
func (hp *HerokuProvider) IsSupported() bool {
	return hp.buildCommit != "" && hp.releaseVerion != ""
}
