package time

import "time"

type Type int64

const (
	Nanosecond  Type = 1
	Microsecond      = 1000 * Nanosecond
	Millisecond      = 1000 * Microsecond
	Second           = 1000 * Millisecond
	Minute           = 60 * Second
	Hour             = 60 * Minute
)

func (t Type) Int64() int64 {
	return int64(t)
}

func (t Type) Nanosecond() int64 {
	return int64(t) / int64(time.Nanosecond)
}

func (t Type) Microsecond() int64 {
	return int64(t) / int64(time.Microsecond)
}

func (t Type) Millisecond() int64 {
	return int64(t) / int64(time.Millisecond)
}

func (t Type) Second() int64 {
	return int64(t) / int64(time.Second)
}
