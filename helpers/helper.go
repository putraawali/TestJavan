package helpers

import "flag"

type key string

var (
	ProjectFolder        = flag.String("folder", "./", "absolute path of project folder")
	RequestIDKey  key    = "X-Request-ID"
	RequestID     string = "X-Request-ID"
)
