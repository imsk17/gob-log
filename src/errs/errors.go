package errs

import "errors"

var (
	ErrNoBlogFound = errors.New("unable to find the blog")
)
