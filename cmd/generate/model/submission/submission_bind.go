package submission

import (
	"github.com/Myriad-Dreamin/boj-v6/abstract/problem"
	"github.com/Myriad-Dreamin/boj-v6/abstract/submission"
	"github.com/Myriad-Dreamin/boj-v6/abstract/user"
	user_problem2 "github.com/Myriad-Dreamin/boj-v6/abstract/user_problem"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"os"
)

func NewController(m module.Module) (*Controller, error) {
	s := new(Controller)
	s.db = m.RequireImpl(new(submission.DB)).(submission.DB)
	s.problemDB = m.RequireImpl(new(problem.DB)).(problem.DB)
	s.userDB = m.RequireImpl(new(user.DB)).(user.DB)
	s.logger = m.RequireImpl(new(external.Logger)).(external.Logger)
	s.cfg = m.RequireImpl(new(*config.ServerConfig)).(*config.ServerConfig)
	s.userTriedProblemDB = m.RequireImpl(new(user_problem2.TriedDB)).(user_problem2.TriedDB)
	s.dispatcher = m.RequireImpl(new(submission.Dispatcher)).(submission.Dispatcher)
	s.problemKey = "pid"
	s.key = "sid"

	//s.inner = m.RequireImpl(new(inner_control.InnerSubmissionController)).(inner_control.InnerSubmissionController)
	return s, nil
}

func WriteToFileSystem(directory string, fullPath string, code string) (err error) {

	if _, err = os.Stat(directory); err != nil {
		if !os.IsNotExist(err) {
			return nil
		}

		// todo: 处于安全考虑，需要合理商量控制一下perm
		err = os.Mkdir(directory, 0777)
		if err != nil {
			return err
		}
	}

	if _, err = os.Stat(fullPath); err != nil {
		if !os.IsNotExist(err) {
			return nil
		}

		var f *os.File
		f, err = os.Create(fullPath)
		if err != nil {
			return err
		}
		_, err = f.WriteString(code)
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
	}
	return
}
