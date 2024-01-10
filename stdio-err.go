// Copyright (c) 2024  The Go-Curses Authors
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

package stdio

import (
	"bytes"
	"os"
)

var (
	// StderrTempPattern is the pattern used with calling os.CreateTemp to make
	// the new temp file during os.Stderr Capture operations
	StderrTempPattern = "corelibs-mock-stdio.*.err"
)

var _ Stdio = (*cStderr)(nil)

type cStderr struct {
	cStdio
}

// NewStderr creates a new Stdio instance that will Capture os.Stderr
func NewStderr() (s Stdio) {
	s = &cStderr{}
	return
}

func (s *cStderr) Capture() (err error) {
	if s.R != nil && s.W != nil {
		err = ErrAlreadyCaptured
		return
	}
	s.Original = os.Stderr
	if s.W, err = os.CreateTemp("", StderrTempPattern); err == nil {
		s.R = s.W
		os.Stderr = s.W
	}
	return
}

func (s *cStderr) Restore() {
	if s.Original != nil {
		os.Stderr = s.Original
	}
	name := s.W.Name()
	_ = s.W.Close()
	data, _ := os.ReadFile(name)
	s.B = bytes.NewBuffer(data)
	_ = os.Remove(name)
	return
}
