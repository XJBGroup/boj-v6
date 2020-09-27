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

func (svc *Controller) ChangeProblemDescriptionRef(c controller.MContext) {
	var req = new(api.ChangeProblemDescriptionRefRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	var obj, e = svc.descDB.QueryByKey(id, req.Name)
	if snippet.MaybeSelectError(c, obj, e) {
		return
	}

	obj.Name = req.NewName
	_, err := svc.descDB.UpdateFields(obj, []string{"name"})
	if !snippet.UpdateFields(c, err) {
		return
	}

	var rollbacks []func()
	rollbacks = append(rollbacks, func() {
		obj.Name = req.Name
		_, err := svc.descDB.UpdateFields(obj, []string{"name"})
		if !snippet.UpdateFields(c, err) {
			svc.logger.Error("rollback name error")
		}
	})

	obj.Name = req.Name
	obj.Key = nil
	e = svc.descDB.LoadDesc(obj)
	if e != nil {
		snippet.DoRollback(rollbacks)
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeProblemDescLoadError,
			ErrorS: e.Error(),
		})
		return
	}

	obj.Name = req.NewName
	obj.Key = nil
	e = svc.descDB.SaveDesc(obj)
	if e != nil {
		snippet.DoRollback(rollbacks)
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeProblemDescSaveError,
			ErrorS: e.Error(),
		})
		return
	}
	rollbacks = append(rollbacks, func() {
		obj.Name = req.NewName
		obj.Key = nil
		err := svc.descDB.DeleteDesc(obj)
		if err != nil {
			svc.logger.Error("delete redundancy problem description error",
				"problem_id", obj.ProblemID, "name", obj.Name, "error", err.Error())
		}
	})

	obj.Name = req.Name
	obj.Key = nil
	e = svc.descDB.DeleteDesc(obj)
	if e != nil {
		snippet.DoRollback(rollbacks)
		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeProblemDescDeleteError,
			ErrorS: e.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snippet.ResponseOK)
}

func (svc *Controller) PostProblemDesc(c controller.MContext) {
	var req = new(api.PostProblemDescRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	var obj = new(problem_desc.ProblemDesc)
	obj.Name = req.Name
	obj.Content = []byte(req.Content)
	obj.ProblemID = id

	a, e := svc.descDB.Create(obj)
	if !snippet.CreateObj(c, svc.descDB.UnwrapError, a, e) {
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
			Code:   types.CodeProblemDescSaveError,
			ErrorS: e.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snippet.ResponseOK)
	e = svc.descDB.ReleaseDesc(obj)
	if e != nil {
		svc.logger.Error("cannot release bytes object of problem description", "error", e.Error())
	}
}

func (svc *Controller) GetProblemDesc(c controller.MContext) {
	var req = new(api.GetProblemDescRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	var obj = new(problem_desc.ProblemDesc)
	obj.Name = req.Name
	obj.ProblemID = id

	err := svc.descDB.LoadDesc(obj)
	if err != nil {
		c.JSON(http.StatusOK, serial.ErrorSerializer{
			Code:   types.CodeProblemDescLoadError,
			ErrorS: err.Error(),
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

func (svc *Controller) PutProblemDesc(c controller.MContext) {
	var req = new(api.PostProblemDescRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	var obj = new(problem_desc.ProblemDesc)
	obj.Name = req.Name
	obj.Content = []byte(req.Content)
	obj.ProblemID = id

	a, e := svc.descDB.QueryExistenceByKey(obj.ProblemID, obj.Name)
	if snippet.MaybeQueryExistenceError(c, a, e) {
		return
	}

	e = svc.descDB.SaveDesc(obj)
	if e != nil {
		c.JSON(http.StatusOK, serial.ErrorSerializer{
			Code:   types.CodeProblemDescSaveError,
			ErrorS: e.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snippet.ResponseOK)
	e = svc.descDB.ReleaseDesc(obj)
	if e != nil {
		svc.logger.Error("cannot release bytes object of problem description", "error", e.Error())
	}
}

func (svc *Controller) DeleteProblemDesc(c controller.MContext) {
	var req = new(api.DeleteProblemDescRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	var obj, err = svc.descDB.QueryByKey(id, req.Name)
	if snippet.MaybeSelectError(c, obj, err) {
		return
	}

	aff, err := svc.descDB.Delete(obj)
	if !snippet.DeleteObj(c, aff, err) {
		return
	}

	err = svc.descDB.DeleteDesc(obj)
	if err != nil {
		aff2, err2 := svc.descDB.Create(obj)
		if aff2 == 0 || err2 != nil {
			svc.logger.Error("recreate error",
				"affect", aff2, "error", snippet.ConvertErrorToString(err2))
		}

		c.AbortWithStatusJSON(http.StatusOK, &serial.ErrorSerializer{
			Code:   types.CodeProblemDescDeleteError,
			ErrorS: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snippet.ResponseOK)
}

func (svc *Controller) ListProblemDesc(c controller.MContext) {
	var req = new(api.ListProblemDescRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	ss, err := svc.descDB.Find(id, req.Page, req.PageSize)
	if snippet.MaybeSelectError(c, ss, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeListProblemDescReply(types.CodeOK,
		api.PackSerializeProblemDescData(ss)))
}

func (svc *Controller) CountProblemDesc(c controller.MContext) {
	var req = new(api.ListProblemDescRequest)
	id, ok := snippet.ParseUintAndBind(c, svc.key, req)
	if !ok {
		return
	}

	ss, err := svc.descDB.Count(id)
	if snippet.MaybeCountError(c, err) {
		return
	}

	c.JSON(http.StatusOK, api.SerializeCountProblemDescReply(types.CodeOK, ss))
}
