package memory

type Type int64

const (
	Byte     Type = 1
	KiloByte      = Byte * 1024
	MegaByte      = KiloByte * 1024
	GigaByte      = MegaByte * 1024
)

func (t Type) Int64() int64 {
	return int64(t)
}

func (t Type) Byte() int64 {
	return int64(t) >> 0
}

func (t Type) KiloByte() int64 {
	return int64(t) >> 10
}

func (t Type) MegaByte() int64 {
	return int64(t) >> 20
}

func (t Type) GigaByte() int64 {
	return int64(t) >> 30
}
