package problem

import (
	"archive/zip"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (svc Service) ProblemFSMkdir(c controller.MContext) {
	var req api.ProblemFSMkdirRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}

	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), req.Path)

	if err := os.MkdirAll(path, 0755); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeFSExecError,
			ErrorS: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.SerializeProblemFSMkdirReply(types.CodeOK))
}

func (svc Service) ProblemFSZipWrite(c controller.MContext) {
	var req api.ProblemFSZipWriteRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}
	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), req.Path)

	file, err := c.FormFile("upload")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeUploadFileError,
			ErrorS: err.Error(),
		})
		return
	}

	zipName := filepath.Join(path, file.Filename)
	if err = c.SaveUploadedFile(file, zipName); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	}

	r, err := zip.OpenReader(zipName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	}

	defer func() {
		err := r.Close()
		if err != nil {
			svc.logger.Debug("error occurs", "error", err)
		}
		err = os.Remove(zipName)
		if err != nil {
			svc.logger.Debug("error occurs", "error", err)
			return
		}
	}()

	for _, file := range r.File {
		rc, err := file.Open()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
				Code:   types.CodeFSExecError,
				ErrorS: err.Error(),
			})
			return
		}
		filename := filepath.Join(path, file.Name)
		err = os.MkdirAll(filepath.Dir(filename), 0770)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
				Code:   types.CodeFSExecError,
				ErrorS: err.Error(),
			})
			_ = rc.Close()
			return
		}
		w, err := os.Create(filename)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
				Code:   types.CodeFSExecError,
				ErrorS: err.Error(),
			})
			_ = rc.Close()
			return
		}
		_, err = io.Copy(w, rc)
		_ = rc.Close()
		_ = w.Close()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
				Code:   types.CodeFSExecError,
				ErrorS: err.Error(),
			})
			return
		}

	}

	c.JSON(http.StatusOK, api.SerializeProblemFSZipWriteReply(types.CodeOK))
}

func (svc Service) ProblemFSZipRead(c controller.MContext) {
	var req api.ProblemFSZipReadRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}
	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), req.Path)

	c.Header("Content-type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=\""+filepath.Base(path)+".zip\"")

	c.Stream(func(w io.Writer) bool {
		compressed := zip.NewWriter(w)

		err := filepath.Walk(path, func(path string, info os.FileInfo, upstreamErr error) (err error) {
			if upstreamErr != nil {
				return upstreamErr
			}
			fileReader, err := os.Open(path)
			if err != nil {
				return err
			}
			defer func() {
				if err != nil {
					err2 := fileReader.Close()
					if err2 != nil {
						svc.logger.Error("close file error", "path", path, "error", err2.Error())
					}
				} else {
					err = fileReader.Close()
				}

			}()
			compressedFileWriter, err := compressed.Create(path)
			if err != nil {
				return err
			}
			_, err = io.Copy(compressedFileWriter, fileReader)

			return err
		})

		if err != nil {
			if _, ok := err.(*serial.ErrorSerializer); ok {
				c.AbortWithStatusJSON(http.StatusOK, err)
			} else {
				c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
					Code:   types.CodeFSExecError,
					ErrorS: err.Error(),
				})
			}
			return true
		}

		err = compressed.Close()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
				Code:   types.CodeFSExecError,
				ErrorS: err.Error(),
			})
		}
		return false
	})

	//c.JSON(http.StatusOK, api.SerializeProblemFSZipReadReply(types.CodeOK))
}

func (svc Service) ProblemFSLS(c controller.MContext) {
	var req api.ProblemFSLSRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}

	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), req.Path)

	if files, err := ioutil.ReadDir(path); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	} else {
		innerResults := make([]api.ProblemFSLSInnerReply, 0, len(files))
		for _, stat := range files {
			innerResults = append(innerResults,
				api.SerializeProblemFSLSInnerReply(
					stat.Name(), stat.Size(), stat.IsDir(), stat.ModTime()))
		}

		c.JSON(http.StatusOK, api.SerializeProblemFSLSReply(types.CodeOK, innerResults))
	}
}

func (svc Service) ProblemFSWrites(c controller.MContext) {
	var req api.ProblemFSWritesRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}
	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), req.Path)

	if err := os.MkdirAll(path, 0770); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeUploadFileError,
			ErrorS: err.Error(),
		})
		return
	}
	files := form.File["upload"]
	for _, file := range files {
		if err = c.SaveUploadedFile(file, path+file.Filename); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
				Code:   types.CodeFSExecError,
				ErrorS: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, api.SerializeProblemFSWritesReply(types.CodeOK))
}

func (svc Service) ProblemFSRemoveAll(c controller.MContext) {
	var req api.ProblemFSRemoveAllRequest
	id, ok := svc.BindProblemFSRequest(c, &req)
	if !ok {
		return
	}

	path := filepath.Join(svc.cfg.PathConfig.ProblemPath, strconv.Itoa(int(id)), req.Path)

	if err := os.RemoveAll(path); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeStatError,
			ErrorS: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snippet.ResponseOK)
}
