package problem_desc

import (
	"encoding/binary"
	problem_desc "github.com/Myriad-Dreamin/boj-v6/abstract/problem-desc"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

func NewDB(m module.Module) (*DBImpl, error) {
	d, e := newDB(m)
	return &DBImpl{
		db:        d,
		ossEngine: m.RequireImpl(new(*external.OSSEngine)).(*external.OSSEngine),
	}, e
}

type DBImpl struct {
	db
	ossEngine *external.OSSEngine
}

const ProblemDescPrefix = "pd:"

func CreateProblemDescKey(pid uint, name string) (k []byte) {
	k = make([]byte, 3+8+len(name))
	copy(k, ProblemDescPrefix)
	binary.BigEndian.PutUint64(k[3:11], uint64(pid))
	copy(k[11:], name)
	return
}

func (db DBImpl) NewProblemDesc(pid uint, name string, content []byte) *problem_desc.ProblemDesc {
	return &problem_desc.ProblemDesc{
		ProblemID:  pid,
		Name:       name,
		Content:    content,
		Key:        CreateProblemDescKey(pid, name),
		FreeHandle: nil,
	}
}

//func GetProblemDesc(pid uint, name string) (ProblemDesc, error) {
//	k := CreateProblemDescKey(pid, name)
//	b, err := engine.Get(k)
//	if err != nil {
//		return ProblemDesc{}, err
//	}
//	return ProblemDesc{
//		Content:    b.Data(),
//		K:          k,
//		freeHandle: b.Free,
//	}, nil
//}

func (db DBImpl) LoadDesc(a *problem_desc.ProblemDesc) (err error) {
	db.ensureKey(a)
	err = db.ReleaseDesc(a)
	if err != nil {
		return
	}
	b, err := db.ossEngine.Get(a.Key)
	if err != nil {
		return err
	}
	a.FreeHandle = b.Free
	a.Content = b.Data()
	return nil
}

func (db DBImpl) SaveDesc(a *problem_desc.ProblemDesc) error {
	db.ensureKey(a)
	return db.ossEngine.Put(a.Key, a.Content)
}

func (db DBImpl) ReleaseDesc(a *problem_desc.ProblemDesc) error {
	if a.FreeHandle != nil {
		a.FreeHandle()
		a.FreeHandle = nil
	}
	return nil
}

func (db DBImpl) DeleteDesc(a *problem_desc.ProblemDesc) error {
	db.ensureKey(a)
	return db.ossEngine.Delete(a.Key)
}

func (db DBImpl) RenameDesc(a *problem_desc.ProblemDesc, newName string) (int64, error) {
	if a.FreeHandle == nil {
		err := db.LoadDesc(a)
		if err != nil {
			return 0, err
		}
	}

	err := db.DeleteDesc(a)
	if err != nil {
		return 0, err
	}

	a.Name = newName
	a.Key = CreateProblemDescKey(a.ProblemID, a.Name)

	err = db.SaveDesc(a)
	if err != nil {
		return 0, err
	}

	// todo: update database
	return 1, nil
}

func (db DBImpl) ensureKey(a *problem_desc.ProblemDesc) {
	if len(a.Key) == 0 {
		a.Key = CreateProblemDescKey(a.ProblemID, a.Name)
	}
}

func (db DBImpl) InvalidateDescCache(a *problem_desc.ProblemDesc) error {
	err := db.ReleaseDesc(a)
	if err != nil {
		return err
	}
	a.Key = nil
	return nil
}

func (db DBImpl) QueryByKey(pid uint, pdName string) (pd *problem_desc.ProblemDesc, err error) {
	pd = new(problem_desc.ProblemDesc)
	return pd, db.GORMDBImpl.Query(pd, "pid = ? and name = ?", pid, pdName)
}

func (db DBImpl) QueryByPID(pid uint) (pd []problem_desc.ProblemDesc, err error) {
	return pd, db.GORMDBImpl.Query(&pd, "pid = ?", pid)
}
