package dependencies

import (
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"

	"github.com/thtg88/blog.marco-marassi.com/pkg/etagproviders"
	"github.com/thtg88/blog.marco-marassi.com/pkg/herokumetadata"
	"github.com/thtg88/blog.marco-marassi.com/pkg/timer"
)

const (
	defaultLogPrefix            string = "[wedding] "
	defaultPort                 string = "8080"
	defaultFileServerDir        string = "./static"
	portEnvironmentVariableName string = "PORT"
)

type Logger interface {
	Fatal(v ...any)
	Printf(format string, v ...any)
	Println(v ...any)
}

type Dependencies struct {
	ETag       string
	FileServer http.Handler
	Logger     Logger
	Port       string
	Timer      timer.Timer
}

// Initialize initializes the common dependencies required by the rest of the program.
// This initializes the logger, timer, file server, TCP port the server will listen on, and ETag header the server will emit.
// This function is to be used at app boot time.
func Initialize() *Dependencies {
	deps := &Dependencies{
		FileServer: http.FileServer(http.Dir(defaultFileServerDir)),
		Logger:     log.New(os.Stdout, defaultLogPrefix, log.LstdFlags),
		Timer:      &timer.RealTimer{},
	}

	var herokuMetadata herokumetadata.Metadata
	// We are not interested in failing if we error, just log it
	if err := herokumetadata.Parse(herokumetadata.DefaultMetadataFilename, &herokuMetadata); err != nil {
		deps.Logger.Printf("could not parse heroku metadata: %v\n", err)
	}

	etagProvider := etagproviders.BuildProvider(
		herokuMetadata,
		os.Getenv(etagproviders.HerokuBuildCommitEnvironmentVariableName),
		os.Getenv(etagproviders.HerokuReleaseVersionEnvironmentVariableName),
		uuid.New(),
	)
	deps.Logger.Printf("etag provider is %s", etagProvider.GetName())

	deps.ETag = etagProvider.GetETag()
	deps.Logger.Printf("using etag %s\n", deps.ETag)

	deps.Port = os.Getenv(portEnvironmentVariableName)
	if deps.Port == "" {
		deps.Logger.Println("port not provided, defaulting")
		deps.Port = defaultPort
	}

	return deps
}
