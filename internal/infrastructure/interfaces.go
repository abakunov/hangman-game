package infrastructure

import "io"

//go:generate mockery --name=HangmanReadCloser --dir=. --output=mocks --outpkg=mocks --case=underscore

type HangmanReadCloser interface {
	io.ReadCloser
}
