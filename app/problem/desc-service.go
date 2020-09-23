package problem

import (
	problem_desc "github.com/Myriad-Dreamin/boj-v6/abstract/problem-desc"
	"github.com/Myriad-Dreamin/boj-v6/api"
	"github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/lib/serial"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

func (svc Service) ChangeProblemDescriptionRef(c controller.MContext) {
	panic("implement me")
}

func (svc Service) PostProblemDesc(c controller.MContext) {
	var req = new(api.PostProblemDescRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	var obj = new(problem_desc.ProblemDesc)
	obj.Name = req.Name
	obj.Content = []byte(req.Content)
	obj.ProblemID = id

	a, e := svc.descDB.Create(obj)
	if !snippet.CreateObj(c, a, e) {
		return
	}

	e = svc.descDB.SaveDesc(obj)
	if e != nil {
		aff, err := svc.descDB.Delete(obj)
		if aff != 1 || err == nil {
			svc.logger.Error("cannot clear not saved problem description", "error",
				snippet.ConvertErrorToString(err), "affect", aff)
		}
		c.JSON(http.StatusOK, serial.ErrorSerializer{
			Code:  types.CodeProblemDescSaveError,
			Error: e.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snippet.ResponseOK)
	e = svc.descDB.ReleaseDesc(obj)
	if e != nil {
		svc.logger.Error("cannot release bytes object of problem description", "error", e.Error())
	}
}

func (svc Service) GetProblemDesc(c controller.MContext) {
	var req = new(api.GetProblemDescRequest)
	id, ok := snippet.ParseUintAndBind(c, "pid", req)
	if !ok {
		return
	}

	var obj = new(problem_desc.ProblemDesc)
	obj.Name = req.Name
	obj.ProblemID = id

	err := svc.descDB.LoadDesc(obj)
	if err != nil {
		c.JSON(http.StatusOK, serial.ErrorSerializer{
			Code:  types.CodeProblemDescLoadError,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.GetProblemDescReply{
		Code: 0,
		Data: string(obj.Content),
	})

	err = svc.descDB.ReleaseDesc(obj)
	if err != nil {
		svc.logger.Error("cannot release bytes object of problem description", "error", err.Error())
	}
}

func (svc Service) PutProblemDesc(c controller.MContext) {
	panic("implement me")
}

func (svc Service) DeleteProblemDesc(c controller.MContext) {
	panic("implement me")
}
