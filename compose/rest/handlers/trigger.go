package handlers

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `trigger.go`, `trigger.util.go` or `trigger_test.go` to
	implement your API calls, helper functions and tests. The file `trigger.go`
	is only generated the first time, and will not be overwritten if it exists.
*/

import (
	"context"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/titpetric/factory/resputil"

	"github.com/crusttech/crust/compose/rest/request"
	"github.com/crusttech/crust/internal/logger"
)

// Internal API interface
type TriggerAPI interface {
	List(context.Context, *request.TriggerList) (interface{}, error)
	Create(context.Context, *request.TriggerCreate) (interface{}, error)
	Read(context.Context, *request.TriggerRead) (interface{}, error)
	Update(context.Context, *request.TriggerUpdate) (interface{}, error)
	Delete(context.Context, *request.TriggerDelete) (interface{}, error)
}

// HTTP API interface
type Trigger struct {
	List   func(http.ResponseWriter, *http.Request)
	Create func(http.ResponseWriter, *http.Request)
	Read   func(http.ResponseWriter, *http.Request)
	Update func(http.ResponseWriter, *http.Request)
	Delete func(http.ResponseWriter, *http.Request)
}

func NewTrigger(th TriggerAPI) *Trigger {
	return &Trigger{
		List: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewTriggerList()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Trigger.List", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := th.List(r.Context(), params); err != nil {
				logger.LogControllerError("Trigger.List", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Trigger.List", r, params)
				switch fn := value.(type) {
				case func(http.ResponseWriter, *http.Request):
					fn(w, r)
					return
				}
				resputil.JSON(w, value)
				return
			}
		},
		Create: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewTriggerCreate()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Trigger.Create", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := th.Create(r.Context(), params); err != nil {
				logger.LogControllerError("Trigger.Create", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Trigger.Create", r, params)
				switch fn := value.(type) {
				case func(http.ResponseWriter, *http.Request):
					fn(w, r)
					return
				}
				resputil.JSON(w, value)
				return
			}
		},
		Read: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewTriggerRead()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Trigger.Read", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := th.Read(r.Context(), params); err != nil {
				logger.LogControllerError("Trigger.Read", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Trigger.Read", r, params)
				switch fn := value.(type) {
				case func(http.ResponseWriter, *http.Request):
					fn(w, r)
					return
				}
				resputil.JSON(w, value)
				return
			}
		},
		Update: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewTriggerUpdate()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Trigger.Update", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := th.Update(r.Context(), params); err != nil {
				logger.LogControllerError("Trigger.Update", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Trigger.Update", r, params)
				switch fn := value.(type) {
				case func(http.ResponseWriter, *http.Request):
					fn(w, r)
					return
				}
				resputil.JSON(w, value)
				return
			}
		},
		Delete: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewTriggerDelete()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Trigger.Delete", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := th.Delete(r.Context(), params); err != nil {
				logger.LogControllerError("Trigger.Delete", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Trigger.Delete", r, params)
				switch fn := value.(type) {
				case func(http.ResponseWriter, *http.Request):
					fn(w, r)
					return
				}
				resputil.JSON(w, value)
				return
			}
		},
	}
}

func (th *Trigger) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Get("/namespace/{namespaceID}/trigger/", th.List)
		r.Post("/namespace/{namespaceID}/trigger/", th.Create)
		r.Get("/namespace/{namespaceID}/trigger/{triggerID}", th.Read)
		r.Post("/namespace/{namespaceID}/trigger/{triggerID}", th.Update)
		r.Delete("/namespace/{namespaceID}/trigger/{triggerID}", th.Delete)
	})
}
