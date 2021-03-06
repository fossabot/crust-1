package handlers

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `namespace.go`, `namespace.util.go` or `namespace_test.go` to
	implement your API calls, helper functions and tests. The file `namespace.go`
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
type NamespaceAPI interface {
	List(context.Context, *request.NamespaceList) (interface{}, error)
	Create(context.Context, *request.NamespaceCreate) (interface{}, error)
	Read(context.Context, *request.NamespaceRead) (interface{}, error)
	Update(context.Context, *request.NamespaceUpdate) (interface{}, error)
	Delete(context.Context, *request.NamespaceDelete) (interface{}, error)
}

// HTTP API interface
type Namespace struct {
	List   func(http.ResponseWriter, *http.Request)
	Create func(http.ResponseWriter, *http.Request)
	Read   func(http.ResponseWriter, *http.Request)
	Update func(http.ResponseWriter, *http.Request)
	Delete func(http.ResponseWriter, *http.Request)
}

func NewNamespace(nh NamespaceAPI) *Namespace {
	return &Namespace{
		List: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewNamespaceList()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Namespace.List", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := nh.List(r.Context(), params); err != nil {
				logger.LogControllerError("Namespace.List", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Namespace.List", r, params)
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
			params := request.NewNamespaceCreate()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Namespace.Create", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := nh.Create(r.Context(), params); err != nil {
				logger.LogControllerError("Namespace.Create", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Namespace.Create", r, params)
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
			params := request.NewNamespaceRead()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Namespace.Read", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := nh.Read(r.Context(), params); err != nil {
				logger.LogControllerError("Namespace.Read", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Namespace.Read", r, params)
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
			params := request.NewNamespaceUpdate()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Namespace.Update", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := nh.Update(r.Context(), params); err != nil {
				logger.LogControllerError("Namespace.Update", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Namespace.Update", r, params)
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
			params := request.NewNamespaceDelete()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Namespace.Delete", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := nh.Delete(r.Context(), params); err != nil {
				logger.LogControllerError("Namespace.Delete", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Namespace.Delete", r, params)
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

func (nh *Namespace) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Get("/namespace/", nh.List)
		r.Post("/namespace/", nh.Create)
		r.Get("/namespace/{namespaceID}", nh.Read)
		r.Post("/namespace/{namespaceID}", nh.Update)
		r.Delete("/namespace/{namespaceID}", nh.Delete)
	})
}
