package parse

import (
	"bytes"
	"encoding/hex"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nikandfor/tlog"
)

func testReader(t *testing.T, neww func(io.Writer) tlog.Writer, newr func(io.Reader) Reader) {
	const Prefix = "github.com/nikandfor/tlog/"

	var buf bytes.Buffer
	tm := time.Date(2019, 7, 31, 18, 21, 2, 0, time.Local)
	now := func() int64 {
		tm = tm.Add(time.Second)
		return tm.UnixNano()
	}

	w := neww(&buf)

	_ = w.Labels(tlog.Labels{"a", "b=c"}, ID{1, 2, 3, 4})

	_ = w.Message(tlog.Message{
		Location: tlog.Caller(0),
		Time:     now(),
		Format:   "%v",
		Args:     []interface{}{3},
	}, tlog.ID{})

	_ = w.SpanStarted(ID{1}, ID{}, now(), tlog.Caller(0))

	_ = w.Message(tlog.Message{
		Location: tlog.Caller(0),
		Time:     now(),
		Format:   "%v",
		Args:     []interface{}{5},
	}, ID{1})

	_ = w.SpanStarted(ID{2}, ID{1}, now(), tlog.Caller(0))

	_ = w.SpanFinished(ID{2}, time.Second.Nanoseconds())

	_ = w.SpanFinished(ID{1}, 2*time.Second.Nanoseconds())

	t.Logf("data:\n%v", hex.Dump(buf.Bytes()))

	// read
	r := newr(&buf)
	if r, ok := r.(*ProtoReader); ok {
		r.buf = r.buf[:10]
	}

	var res []interface{}
	var err error
	locs := map[uintptr]uintptr{}

	for {
		var o interface{}

		o, err = r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Errorf("Read: %v", err)
			break
		}

		switch t := o.(type) {
		case Location:
			ex, ok := locs[t.PC]
			if !ok {
				ex = uintptr(len(locs) + 1)
				locs[t.PC] = ex
			}
			t.PC = ex
			t.Entry = 0x100 + ex
			t.File = strings.TrimPrefix(t.File, Prefix) // cut prefix in case of repo is not in GOPATH or similar folder structure
			o = t
		case Message:
			t.Location = locs[t.Location]
			o = t
		case SpanStart:
			t.Location = locs[t.Location]
			o = t
		}

		res = append(res, o)
	}

	if err != io.EOF {
		return
	}

	tm = time.Date(2019, 7, 31, 18, 21, 2, 0, time.Local)

	assert.Equal(t, []interface{}{
		Labels{Labels: tlog.Labels{"a", "b=c"}, Span: ID{1, 2, 3, 4}},
		Location{
			PC:    1,
			Entry: 0x101,
			Name:  "github.com/nikandfor/tlog/parse.testReader",
			File:  "parse/proto_reader_test.go",
			Line:  31,
		},
		Message{
			Span:     ID{},
			Location: 1,
			Time:     now(),
			Text:     "3",
		},
		Location{
			PC:    2,
			Entry: 0x102,
			Name:  "github.com/nikandfor/tlog/parse.testReader",
			File:  "parse/proto_reader_test.go",
			Line:  37,
		},
		SpanStart{
			ID:       ID{1},
			Parent:   ID{},
			Location: 2,
			Started:  now(),
		},
		Location{
			PC:    3,
			Entry: 0x103,
			Name:  "github.com/nikandfor/tlog/parse.testReader",
			File:  "parse/proto_reader_test.go",
			Line:  40,
		},
		Message{
			Span:     ID{1},
			Location: 3,
			Time:     now(),
			Text:     "5",
		},
		Location{
			PC:    4,
			Entry: 0x104,
			Name:  "github.com/nikandfor/tlog/parse.testReader",
			File:  "parse/proto_reader_test.go",
			Line:  46,
		},
		SpanStart{
			ID:       ID{2},
			Parent:   ID{1},
			Location: 4,
			Started:  now(),
		},
		SpanFinish{
			ID:      ID{2},
			Elapsed: 1 * time.Second.Nanoseconds(),
		},
		SpanFinish{
			ID:      ID{1},
			Elapsed: 2 * time.Second.Nanoseconds(),
		},
	}, res)
}

func TestProtoReader(t *testing.T) {
	testReader(t,
		func(w io.Writer) tlog.Writer { return tlog.NewProtoWriter(w) },
		func(r io.Reader) Reader { return NewProtoReader(r) },
	)
}
