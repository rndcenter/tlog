package parse

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nikandfor/tlog"
)

func TestJSONReader(t *testing.T) {
	var buf bytes.Buffer
	tm := time.Date(2019, 7, 31, 18, 21, 2, 0, time.UTC)
	now := func() time.Time {
		tm = tm.Add(time.Second)
		return tm
	}

	w := tlog.NewJSONWriter(&buf)

	w.Labels(Labels{"a", "b=c"})

	w.Message(tlog.Message{
		Location: tlog.Caller(0),
		Time:     time.Duration(now().UnixNano()),
		Format:   "%v",
		Args:     []interface{}{3},
	}, tlog.Span{})

	w.SpanStarted(tlog.Span{
		ID:      1,
		Started: now(),
	}, 0, tlog.Caller(0))

	w.Message(tlog.Message{
		Location: tlog.Caller(0),
		Time:     time.Duration(now().UnixNano()),
		Format:   "%v",
		Args:     []interface{}{5},
	}, tlog.Span{ID: 1})

	w.SpanStarted(tlog.Span{
		ID:      2,
		Started: now(),
	}, 1, tlog.Caller(0))

	w.SpanFinished(tlog.Span{ID: 2}, time.Second)

	w.SpanFinished(tlog.Span{ID: 1, Flags: 5}, 2*time.Second)

	t.Logf("json\n%s", buf.Bytes())

	// read
	var res []interface{}
	var err error
	r := NewJSONReader(&buf)
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
			o = t
		case Message:
			t.Location = locs[t.Location]
			o = t
		case Span:
			t.Location = locs[t.Location]
			t.Started = t.Started.UTC()
			o = t
		}

		res = append(res, o)
	}

	if err != io.EOF {
		return
	}

	tm = time.Date(2019, 7, 31, 18, 21, 2, 0, time.UTC)

	assert.Equal(t, []interface{}{
		Labels{"a", "b=c"},
		Location{
			PC:   1,
			Name: "github.com/nikandfor/tlog/parse.TestJSONReader",
			File: "github.com/nikandfor/tlog/parse/json_reader_test.go",
			Line: 27,
		},
		Message{
			Span:     0,
			Location: 1,
			Time:     time.Duration(now().UnixNano()),
			Text:     "3",
		},
		Location{
			PC:   2,
			Name: "github.com/nikandfor/tlog/parse.TestJSONReader",
			File: "github.com/nikandfor/tlog/parse/json_reader_test.go",
			Line: 36,
		},
		Span{
			ID:       1,
			Parent:   0,
			Location: 2,
			Started:  now(),
		},
		Location{
			PC:   3,
			Name: "github.com/nikandfor/tlog/parse.TestJSONReader",
			File: "github.com/nikandfor/tlog/parse/json_reader_test.go",
			Line: 39,
		},
		Message{
			Span:     1,
			Location: 3,
			Time:     time.Duration(now().UnixNano()),
			Text:     "5",
		},
		Location{
			PC:   4,
			Name: "github.com/nikandfor/tlog/parse.TestJSONReader",
			File: "github.com/nikandfor/tlog/parse/json_reader_test.go",
			Line: 48,
		},
		Span{
			ID:       2,
			Parent:   1,
			Location: 4,
			Started:  now(),
		},
		SpanFinish{
			ID:      2,
			Elapsed: 1 * time.Second,
		},
		SpanFinish{
			ID:      1,
			Elapsed: 2 * time.Second,
			Flags:   5,
		},
	}, res)
}