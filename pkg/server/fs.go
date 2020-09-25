package server

import (
	"github.com/spf13/afero"
	"os"
)

func PrepareFileSystem(mock bool) InitializeAction {
	return func(srv *Server) error {
		return srv.applyInitializeAction([]InitializeAction{
			InitFileSystem(mock),
			InitDirectory("problemPath", srv.Cfg.PathConfig.ProblemPath),
			InitDirectory("submissionPath", srv.Cfg.PathConfig.CodePath),
		})
	}
}

func (srv *Server) DropFileSystem() error {
	return srv.applyInitializeAction([]InitializeAction{
		DropDirectory("problemPath", srv.Cfg.PathConfig.ProblemPath),
		DropDirectory("submissionPath", srv.Cfg.PathConfig.CodePath),
	})
}
func InitFileSystem(mock bool) InitializeAction {
	return func(srv *Server) error {
		if srv.Filesystem == nil {
			if mock {
				srv.Filesystem = afero.NewMemMapFs()
			} else {
				srv.Filesystem = afero.NewOsFs()
			}
		}
		srv.Module.Provide("global/filesystem", srv.Filesystem)
		return nil
	}
}

func InitDirectory(name, path string) InitializeAction {
	return func(srv *Server) error {
		return initDirectory(srv, name, path)
	}
}

func initDirectory(srv *Server, name, path string) error {

	var err error

	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		srv.Logger.Debug("create path directory error", "error", err, "creating-path", path)
		return err
	}

	srv.Logger.Info("host path", "name", name, "creating-path", path)
	return nil
}

func DropDirectory(name, path string) InitializeAction {
	return func(srv *Server) error {
		return dropDirectory(srv, name, path)
	}
}
func dropDirectory(srv *Server, _, path string) error {
	if err := os.RemoveAll(path); err != nil {
		srv.Logger.Debug("drop path directory error", "error", err, "dropping-path", path)
		return err
	}
	return nil
}
