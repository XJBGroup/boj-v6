package external

type OSSEngine struct {
	Engine
}

type Engine interface {
	Get([]byte) (ByteObject, error)
	Put([]byte, []byte) error
	Delete([]byte) error
	Close() error
}

type ByteObject interface {
	Data() []byte
	Free()
}
