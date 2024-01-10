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

var _ Stdio = (*cStdin)(nil)

type cStdin struct {
	cStdio
}

// NewStdin creates a new Stdio instance that will Capture os.Stdin and
// immediately write the `data` given so that it's waiting for immediate
// reading from the captured os.Stdin instance
func NewStdin(data []byte) (s Stdio) {
	s = &cStdin{cStdio{
		Input: data,
		B:     bytes.NewBuffer(data),
	}}
	return
}

func (s *cStdin) Capture() (err error) {
	if s.R != nil && s.W != nil {
		err = ErrAlreadyCaptured
		return
	}
	s.Original = os.Stdin
	if s.R, s.W, err = os.Pipe(); err == nil {
		os.Stdin = s.R
		_, err = s.W.Write(s.Input)
		_ = s.W.Close()
	}
	return
}

func (s *cStdin) Restore() {
	if s.Original != nil {
		os.Stdin = s.Original
	}
	_ = s.R.Close()
	_ = s.W.Close()
	return
}
