package tlog

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"
)

// Location is a program counter alias.
// Function name, file name and line can be obtained from it but only in the same binary where Caller of Funcentry was called.
type Location uintptr

// Trace is a stack trace.
// It's quiet the same as runtime.CallerFrames but more efficient.
type Trace []Location

// Caller returns information about the calling goroutine's stack. The argument s is the number of frames to ascend, with 0 identifying the caller of Caller.
//
// It's hacked version of runtime.Caller with no allocs.
func Caller(s int) Location {
	var pc [1]uintptr
	runtime.Callers(2+s, pc[:])
	return Location(pc[0])
}

// Funcentry returns information about the calling goroutine's stack. The argument s is the number of frames to ascend, with 0 identifying the caller of Caller.
//
// It's hacked version of runtime.Callers -> runtime.CallersFrames -> Frames.Next -> Frame.Entry with no allocs.
func Funcentry(s int) Location {
	var pc [1]uintptr
	runtime.Callers(2+s, pc[:])
	return Location(pc[0]).Entry()
}

// StackTrace returns callers stack trace.
//
// It's hacked version of runtime.Callers -> runtime.CallersFrames -> Frames.Next -> Frame.Entry with only one alloc (resulting slice).
func StackTrace(skip, n int) Trace {
	tr := make([]Location, n)
	return FillStackTrace(1+skip, tr)
}

// FillStackTrace returns callers stack trace into provided array.
//
// It's hacked version of runtime.Callers -> runtime.CallersFrames -> Frames.Next -> Frame.Entry with no allocs.
func FillStackTrace(skip int, tr Trace) Trace {
	n := runtime.Callers(2+skip, *(*[]uintptr)(unsafe.Pointer(&tr)))
	return tr[:n]
}

// String formats Location as base_name.go:line.
//
// Works only in the same binary where Caller of Funcentry was called.
func (l Location) String() string {
	_, file, line := l.NameFileLine()
	file = filepath.Base(file)

	b := []byte(file)
	i := len(b)
	b = append(b, ":        "...)

	n := 1
	for q := line; q != 0; q /= 10 {
		n++
	}

	b = b[:i+n]

	for q, j := line, n-1; j >= 1; j-- {
		b[i+j] = byte(q%10) + '0'
		q /= 10
	}

	return string(b)
}

// Format is fmt.Formatter interface implementation.
// It supports width. Precision sets line number width. '+' prints full path not base.
func (l Location) Format(s fmt.State, c rune) {
	name, file, line := l.NameFileLine()

	nn := file

	if s.Flag('#') {
		nn = name
	}

	if !s.Flag('+') {
		nn = filepath.Base(nn)
		if s.Flag('#') && !s.Flag('-') {
			p := strings.IndexByte(nn, '.')
			nn = nn[p+1:]
		}
	}

	n := 1
	for q := line; q != 0; q /= 10 {
		n++
	}

	p, ok := s.Precision()

	if ok {
		n = 1 + p
	}

	s.Write([]byte(nn))

	w, ok := s.Width()

	if ok {
		p := w - len(nn) - n
		if p > 0 {
			s.Write(spaces[:p])
		}
	}

	var b [20]byte
	copy(b[:], ":                  ")

	for q, j := line, n-1; q != 0 && j >= 1; j-- {
		b[j] = byte(q%10) + '0'
		q /= 10
	}

	s.Write(b[:n])
}

// String formats Trace as list of type_name (file.go:line)
//
// Works only in the same binary where Caller of Funcentry was called.
func (t Trace) String() string {
	var b []byte
	for _, l := range t {
		n, f, l := l.NameFileLine()
		n = path.Base(n)
		b = AppendPrintf(b, "%-60s  at %s:%d\n", n, f, l)
	}
	return string(b)
}

func (t Trace) Format(s fmt.State, c rune) {
	switch {
	case s.Flag('+'):
		for _, l := range t {
			s.Write([]byte("at "))
			l.Format(s, c)
			s.Write([]byte("\n"))
		}
	default:
		for i, l := range t {
			if i != 0 {
				s.Write([]byte(" at "))
			}
			l.Format(s, c)
		}
	}
}

func cropFilename(fn, tp string) string {
	p := strings.LastIndexByte(tp, '/')
	pp := strings.IndexByte(tp[p+1:], '.')
	tp = tp[:p+pp]

again:
	if p = strings.Index(fn, tp); p != -1 {
		return fn[p:]
	}

	p = strings.IndexByte(tp, '/')
	if p == -1 {
		return filepath.Base(fn)
	}
	tp = tp[p+1:]
	goto again
}
