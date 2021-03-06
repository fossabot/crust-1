package handlers

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `record.go`, `record.util.go` or `record_test.go` to
	implement your API calls, helper functions and tests. The file `record.go`
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
type RecordAPI interface {
	Report(context.Context, *request.RecordReport) (interface{}, error)
	List(context.Context, *request.RecordList) (interface{}, error)
	Create(context.Context, *request.RecordCreate) (interface{}, error)
	Read(context.Context, *request.RecordRead) (interface{}, error)
	Update(context.Context, *request.RecordUpdate) (interface{}, error)
	Delete(context.Context, *request.RecordDelete) (interface{}, error)
	Upload(context.Context, *request.RecordUpload) (interface{}, error)
}

// HTTP API interface
type Record struct {
	Report func(http.ResponseWriter, *http.Request)
	List   func(http.ResponseWriter, *http.Request)
	Create func(http.ResponseWriter, *http.Request)
	Read   func(http.ResponseWriter, *http.Request)
	Update func(http.ResponseWriter, *http.Request)
	Delete func(http.ResponseWriter, *http.Request)
	Upload func(http.ResponseWriter, *http.Request)
}

func NewRecord(rh RecordAPI) *Record {
	return &Record{
		Report: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewRecordReport()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Record.Report", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := rh.Report(r.Context(), params); err != nil {
				logger.LogControllerError("Record.Report", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Record.Report", r, params)
				switch fn := value.(type) {
				case func(http.ResponseWriter, *http.Request):
					fn(w, r)
					return
				}
				resputil.JSON(w, value)
				return
			}
		},
		List: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewRecordList()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Record.List", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := rh.List(r.Context(), params); err != nil {
				logger.LogControllerError("Record.List", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Record.List", r, params)
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
			params := request.NewRecordCreate()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Record.Create", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := rh.Create(r.Context(), params); err != nil {
				logger.LogControllerError("Record.Create", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Record.Create", r, params)
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
			params := request.NewRecordRead()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Record.Read", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := rh.Read(r.Context(), params); err != nil {
				logger.LogControllerError("Record.Read", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Record.Read", r, params)
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
			params := request.NewRecordUpdate()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Record.Update", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := rh.Update(r.Context(), params); err != nil {
				logger.LogControllerError("Record.Update", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Record.Update", r, params)
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
			params := request.NewRecordDelete()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Record.Delete", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := rh.Delete(r.Context(), params); err != nil {
				logger.LogControllerError("Record.Delete", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Record.Delete", r, params)
				switch fn := value.(type) {
				case func(http.ResponseWriter, *http.Request):
					fn(w, r)
					return
				}
				resputil.JSON(w, value)
				return
			}
		},
		Upload: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewRecordUpload()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Record.Upload", r, err, params)
				resputil.JSON(w, err)
				return
			}
			if value, err := rh.Upload(r.Context(), params); err != nil {
				logger.LogControllerError("Record.Upload", r, err, params)
				resputil.JSON(w, err)
				return
			} else {
				logger.LogControllerCall("Record.Upload", r, params)
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

func (rh *Record) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Get("/namespace/{namespaceID}/module/{moduleID}/record/report", rh.Report)
		r.Get("/namespace/{namespaceID}/module/{moduleID}/record/", rh.List)
		r.Post("/namespace/{namespaceID}/module/{moduleID}/record/", rh.Create)
		r.Get("/namespace/{namespaceID}/module/{moduleID}/record/{recordID}", rh.Read)
		r.Post("/namespace/{namespaceID}/module/{moduleID}/record/{recordID}", rh.Update)
		r.Delete("/namespace/{namespaceID}/module/{moduleID}/record/{recordID}", rh.Delete)
		r.Post("/namespace/{namespaceID}/module/{moduleID}/record/{recordID}/{fieldName}/attachment", rh.Upload)
	})
}
