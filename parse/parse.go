package parse

import (
	"github.com/nikandfor/tlog"
)

type (
	ID = tlog.ID

	Labels struct {
		Span   ID          `json:"s"`
		Labels tlog.Labels `json:"L"`
	}

	Location struct {
		PC    uintptr `json:"p"`
		Entry uintptr `json:"e"`
		Name  string  `json:"n"`
		File  string  `json:"f"`
		Line  int     `json:"l"`
	}

	SpanStart struct {
		ID     ID `json:"i"`
		Parent ID `json:"p"`

		Location uintptr `json:"l"`

		Started int64 `json:"s"`
	}

	SpanFinish struct {
		ID      ID    `json:"i"`
		Elapsed int64 `json:"e"`
	}

	Message struct {
		Span     ID      `json:"s"`
		Location uintptr `json:"l"`
		Time     int64   `json:"t"`
		Text     string  `json:"m"`
	}

	Metric struct {
		Span  ID      `json:"s"`
		Name  string  `json:"n"`
		Value float64 `json:"v"`
	}

	Type rune

	Reader interface {
		Read() (interface{}, error)
	}

	LowReader interface {
		Type() (Type, error)
		Read() (interface{}, error)

		Labels() (Labels, error)
		Location() (Location, error)
		Message() (Message, error)
		SpanStart() (SpanStart, error)
		SpanFinish() (SpanFinish, error)
	}
)

const (
	TypeNone       Type = 0
	TypeLabels     Type = 'L'
	TypeLocation   Type = 'l'
	TypeMessage    Type = 'm'
	TypeSpanStart  Type = 's'
	TypeSpanFinish Type = 'f'
)

func (t Type) String() string {
	if t == 0 {
		return `0`
	}
	return string(t)
}
