package handlers

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `module.go`, `module.util.go` or `module_test.go` to
	implement your API calls, helper functions and tests. The file `module.go`
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
type ModuleAPI interface {
	List(context.Context, *request.ModuleList) (interface{}, error)
	Create(context.Context, *request.ModuleCreate) (interface{}, error)
	Read(context.Context, *request.ModuleRead) (interface{}, error)
	Update(context.Context, *request.ModuleUpdate) (interface{}, error)
	Delete(context.Context, *request.ModuleDelete) (interface{}, error)
}

// HTTP API interface
type Module struct {
	List   func(http.ResponseWriter, *http.Request)
	Create func(http.ResponseWriter, *http.Request)
	Read   func(http.ResponseWriter, *http.Request)
	Update func(http.ResponseWriter, *http.Request)
	Delete func(http.ResponseWriter, *http.Request)
}

func NewModule(mh ModuleAPI) *Module {
	return &Module{
		List: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewModuleList()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Module.List", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := mh.List(r.Context(), params); err != nil {
				logger.LogControllerError("Module.List", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Module.List", r, params)
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
			params := request.NewModuleCreate()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Module.Create", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := mh.Create(r.Context(), params); err != nil {
				logger.LogControllerError("Module.Create", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Module.Create", r, params)
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
			params := request.NewModuleRead()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Module.Read", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := mh.Read(r.Context(), params); err != nil {
				logger.LogControllerError("Module.Read", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Module.Read", r, params)
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
			params := request.NewModuleUpdate()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Module.Update", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := mh.Update(r.Context(), params); err != nil {
				logger.LogControllerError("Module.Update", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Module.Update", r, params)
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
			params := request.NewModuleDelete()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Module.Delete", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := mh.Delete(r.Context(), params); err != nil {
				logger.LogControllerError("Module.Delete", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Module.Delete", r, params)
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

func (mh *Module) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Get("/namespace/{namespaceID}/module/", mh.List)
		r.Post("/namespace/{namespaceID}/module/", mh.Create)
		r.Get("/namespace/{namespaceID}/module/{moduleID}", mh.Read)
		r.Post("/namespace/{namespaceID}/module/{moduleID}", mh.Update)
		r.Delete("/namespace/{namespaceID}/module/{moduleID}", mh.Delete)
	})
}
