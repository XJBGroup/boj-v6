package server

import (
	"os"
)

func PrepareFileSystem(srv *Server) error {
	return srv.applyInitializeAction([]InitializeAction {
		InitFileSystem("problemPath", srv.Cfg.PathConfig.ProblemPath),
		InitFileSystem("submissionPath", srv.Cfg.PathConfig.CodePath),
	})
}

func (srv *Server) DropFileSystem() error {
	return srv.applyInitializeAction([]InitializeAction {
		DropFileSystem("problemPath", srv.Cfg.PathConfig.ProblemPath),
		DropFileSystem("submissionPath", srv.Cfg.PathConfig.CodePath),
	})
}

func InitFileSystem(name, path string) InitializeAction {
	return func(srv *Server) error {
		return initFileSystem(srv, name, path)
	}
}

func initFileSystem(srv *Server, name, path string) error {
	var err error

	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		srv.Logger.Debug("create path directory error", "error", err, "creating-path", path)
		return err
	}

	srv.Logger.Info("host path", "name", name, "creating-path", path)
	return nil
}


func DropFileSystem(name, path string) InitializeAction {
	return func(srv *Server) error {
		return dropFileSystem(srv, name, path)
	}
}
func dropFileSystem(srv *Server, _, path string) error {
	if err := os.RemoveAll(path); err != nil {
		srv.Logger.Debug("drop path directory error", "error", err, "dropping-path", path)
		return err
	}
	return nil
}