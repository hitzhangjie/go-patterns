package report

import (
	"net/http"
)

// no one can overwrite this unexported global variable
//
// this variable is initialized (when package initialized) only once,
// it won't be overwritten, so it's safe to use in your code as long as
// it's concurrent-safe.
var reporter Reporter = &httpReporter{client: http.DefaultClient}

// interface could be used for decouple abstraction and implemention,
// here we use it for hide concrete type implemention.
type Reporter interface {
	Report(msg string, args ...interface{}) error
}

// this is the single global access point
func GetInstance() Reporter {
	return reporter
}

// no one can access this unexporeted concrete datatype directly,
//
// no one can call `new(reporter.httpReporter)` to create other instances.
type httpReporter struct {
	client *http.Client
}

func (r *httpReporter) Report(msg string, args ...interface{}) error {
	return nil
}
