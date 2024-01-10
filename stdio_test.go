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
	"io"
	"os"
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {

	m := &sync.Mutex{}

	Convey("Stdin", t, func() {
		m.Lock()
		defer m.Unlock()
		input := []byte("one line\ntwo line\n")
		s := NewStdin(input)
		So(s, ShouldNotEqual, nil)
		original := os.Stdin
		So(s.Capture(), ShouldEqual, nil)
		So(os.Stdin, ShouldNotEqual, original)
		So(os.Stdin, ShouldEqual, s.Reader())
		So(s.Capture(), ShouldEqual, ErrAlreadyCaptured)
		data, _ := io.ReadAll(os.Stdin)
		So(data, ShouldEqual, input)
		s.Restore()
		So(os.Stdin, ShouldEqual, original)
	})

	Convey("Stdout", t, func() {
		m.Lock()
		defer m.Unlock()
		s := NewStdout()
		So(s, ShouldNotEqual, nil)
		original := os.Stdout
		checkcap := s.Capture()
		checkout := os.Stdout
		checkerr := s.Capture()
		input := "one line\ntwo line\n"
		_, _ = os.Stdout.WriteString(input)
		s.Restore()
		data := s.Data()
		So(checkcap, ShouldEqual, nil)
		So(checkout, ShouldEqual, s.Writer())
		So(checkerr, ShouldEqual, ErrAlreadyCaptured)
		So(os.Stdout, ShouldEqual, original)
		So(string(data), ShouldEqual, input)
	})

	Convey("Stderr", t, func() {
		m.Lock()
		defer m.Unlock()
		s := NewStderr()
		So(s, ShouldNotEqual, nil)
		original := os.Stderr
		checkcap := s.Capture()
		checkout := os.Stderr
		checkerr := s.Capture()
		input := "one line\ntwo line\n"
		_, _ = os.Stderr.WriteString(input)
		checkdata := s.Data()
		s.Restore()
		data := s.Data()
		So(checkcap, ShouldEqual, nil)
		So(checkout, ShouldEqual, s.Writer())
		So(checkerr, ShouldEqual, ErrAlreadyCaptured)
		So(os.Stderr, ShouldEqual, original)
		So(checkdata, ShouldEqual, data)
		So(string(data), ShouldEqual, input)
	})

}
