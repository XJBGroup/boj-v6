package oss

import (
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LevelDBEngine struct {
	DB   *leveldb.DB
	wOpt *opt.WriteOptions
	rOpt *opt.ReadOptions
}

func (db *LevelDBEngine) Get(k []byte) (external.ByteObject, error) {
	return ToByteObject(db.DB.Get(k, db.rOpt))
}

func (db *LevelDBEngine) Put(k []byte, v []byte) error {
	return db.DB.Put(k, v, db.wOpt)
}

func (db *LevelDBEngine) Delete(k []byte) error {
	return db.DB.Delete(k, db.wOpt)
}

func (db *LevelDBEngine) Close() error {
	return db.DB.Close()
}

func NewLevelDB(path string, opts *opt.Options) (*external.OSSEngine, error) {
	e := new(LevelDBEngine)
	var err error
	e.DB, err = leveldb.OpenFile(path, opts)
	if err != nil {
		return nil, err
	}
	e.wOpt = &opt.WriteOptions{
		NoWriteMerge: false,
		Sync:         true,
	}
	e.rOpt = &opt.ReadOptions{
		DontFillCache: false,
		Strict:        0,
	}
	return &external.OSSEngine{Engine: e}, nil
}

type byteObject []byte

func (b byteObject) Data() []byte { return b }
func (b byteObject) Free()        {}

func ToByteObject(obj []byte, err error) (external.ByteObject, error) {
	if err != nil {
		return nil, err
	}
	return byteObject(obj), nil
}
