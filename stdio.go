// Copyright (c) 2024  The Go-CoreLibs Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package stdio provides utilities for mocking os.Stdin, os.Stdout and
// os.Stderr for unit testing purposes
package stdio

import (
	"bytes"
	"errors"
	"os"
)

var (
	ErrAlreadyCaptured = errors.New("already captured")
)

type Stdio interface {
	// Capture will replace the os file handle with a faked one that writes data
	// to a temporary file
	Capture() (err error)
	// Restore will replace the faked os file handle with the original one
	// captured
	Restore()
	// Data returns the contents of the temporary file
	Data() (data []byte)
	// Reader returns the *os.File instance for read operations
	Reader() (r *os.File)
	// Writer returns the *os.File instance for read operations
	Writer() (w *os.File)
}

type cStdio struct {
	Original *os.File

	Input []byte

	R *os.File
	W *os.File
	B *bytes.Buffer
}

func (s *cStdio) Data() (data []byte) {
	if s.B != nil {
		// buffer created during Restore
		data = s.B.Bytes()
	} else if s.W != nil {
		// haven't restored yet, read from file
		data, _ = os.ReadFile(s.W.Name())
	}
	return
}

func (s *cStdio) Reader() (r *os.File) {
	r = s.R
	return
}

func (s *cStdio) Writer() (r *os.File) {
	r = s.W
	return
}
