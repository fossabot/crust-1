package request

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `attachment.go`, `attachment.util.go` or `attachment_test.go` to
	implement your API calls, helper functions and tests. The file `attachment.go`
	is only generated the first time, and will not be overwritten if it exists.
*/

import (
	"io"
	"strings"

	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

var _ = chi.URLParam
var _ = multipart.FileHeader{}

// Attachment list request parameters
type AttachmentList struct {
	PageID      uint64 `json:",string"`
	ModuleID    uint64 `json:",string"`
	RecordID    uint64 `json:",string"`
	FieldName   string
	Page        uint
	PerPage     uint
	Sign        string
	UserID      uint64 `json:",string"`
	Kind        string
	NamespaceID uint64 `json:",string"`
}

func NewAttachmentList() *AttachmentList {
	return &AttachmentList{}
}

func (aReq *AttachmentList) Fill(r *http.Request) (err error) {
	if strings.ToLower(r.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(r.Body).Decode(aReq)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return errors.Wrap(err, "error parsing http request body")
		}
	}

	if err = r.ParseForm(); err != nil {
		return err
	}

	get := map[string]string{}
	post := map[string]string{}
	urlQuery := r.URL.Query()
	for name, param := range urlQuery {
		get[name] = string(param[0])
	}
	postVars := r.Form
	for name, param := range postVars {
		post[name] = string(param[0])
	}

	if val, ok := get["pageID"]; ok {

		aReq.PageID = parseUInt64(val)
	}
	if val, ok := get["moduleID"]; ok {

		aReq.ModuleID = parseUInt64(val)
	}
	if val, ok := get["recordID"]; ok {

		aReq.RecordID = parseUInt64(val)
	}
	if val, ok := get["fieldName"]; ok {

		aReq.FieldName = val
	}
	if val, ok := get["page"]; ok {

		aReq.Page = parseUint(val)
	}
	if val, ok := get["perPage"]; ok {

		aReq.PerPage = parseUint(val)
	}
	if val, ok := get["sign"]; ok {

		aReq.Sign = val
	}
	if val, ok := get["userID"]; ok {

		aReq.UserID = parseUInt64(val)
	}
	aReq.Kind = chi.URLParam(r, "kind")
	aReq.NamespaceID = parseUInt64(chi.URLParam(r, "namespaceID"))

	return err
}

var _ RequestFiller = NewAttachmentList()

// Attachment read request parameters
type AttachmentRead struct {
	AttachmentID uint64 `json:",string"`
	Kind         string
	NamespaceID  uint64 `json:",string"`
	Sign         string
	UserID       uint64 `json:",string"`
}

func NewAttachmentRead() *AttachmentRead {
	return &AttachmentRead{}
}

func (aReq *AttachmentRead) Fill(r *http.Request) (err error) {
	if strings.ToLower(r.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(r.Body).Decode(aReq)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return errors.Wrap(err, "error parsing http request body")
		}
	}

	if err = r.ParseForm(); err != nil {
		return err
	}

	get := map[string]string{}
	post := map[string]string{}
	urlQuery := r.URL.Query()
	for name, param := range urlQuery {
		get[name] = string(param[0])
	}
	postVars := r.Form
	for name, param := range postVars {
		post[name] = string(param[0])
	}

	aReq.AttachmentID = parseUInt64(chi.URLParam(r, "attachmentID"))
	aReq.Kind = chi.URLParam(r, "kind")
	aReq.NamespaceID = parseUInt64(chi.URLParam(r, "namespaceID"))
	if val, ok := get["sign"]; ok {

		aReq.Sign = val
	}
	if val, ok := get["userID"]; ok {

		aReq.UserID = parseUInt64(val)
	}

	return err
}

var _ RequestFiller = NewAttachmentRead()

// Attachment delete request parameters
type AttachmentDelete struct {
	AttachmentID uint64 `json:",string"`
	Kind         string
	NamespaceID  uint64 `json:",string"`
	Sign         string
	UserID       uint64 `json:",string"`
}

func NewAttachmentDelete() *AttachmentDelete {
	return &AttachmentDelete{}
}

func (aReq *AttachmentDelete) Fill(r *http.Request) (err error) {
	if strings.ToLower(r.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(r.Body).Decode(aReq)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return errors.Wrap(err, "error parsing http request body")
		}
	}

	if err = r.ParseForm(); err != nil {
		return err
	}

	get := map[string]string{}
	post := map[string]string{}
	urlQuery := r.URL.Query()
	for name, param := range urlQuery {
		get[name] = string(param[0])
	}
	postVars := r.Form
	for name, param := range postVars {
		post[name] = string(param[0])
	}

	aReq.AttachmentID = parseUInt64(chi.URLParam(r, "attachmentID"))
	aReq.Kind = chi.URLParam(r, "kind")
	aReq.NamespaceID = parseUInt64(chi.URLParam(r, "namespaceID"))
	if val, ok := get["sign"]; ok {

		aReq.Sign = val
	}
	if val, ok := get["userID"]; ok {

		aReq.UserID = parseUInt64(val)
	}

	return err
}

var _ RequestFiller = NewAttachmentDelete()

// Attachment original request parameters
type AttachmentOriginal struct {
	Download     bool
	Sign         string
	UserID       uint64 `json:",string"`
	AttachmentID uint64 `json:",string"`
	Name         string
	Kind         string
	NamespaceID  uint64 `json:",string"`
}

func NewAttachmentOriginal() *AttachmentOriginal {
	return &AttachmentOriginal{}
}

func (aReq *AttachmentOriginal) Fill(r *http.Request) (err error) {
	if strings.ToLower(r.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(r.Body).Decode(aReq)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return errors.Wrap(err, "error parsing http request body")
		}
	}

	if err = r.ParseForm(); err != nil {
		return err
	}

	get := map[string]string{}
	post := map[string]string{}
	urlQuery := r.URL.Query()
	for name, param := range urlQuery {
		get[name] = string(param[0])
	}
	postVars := r.Form
	for name, param := range postVars {
		post[name] = string(param[0])
	}

	if val, ok := get["download"]; ok {

		aReq.Download = parseBool(val)
	}
	if val, ok := get["sign"]; ok {

		aReq.Sign = val
	}
	if val, ok := get["userID"]; ok {

		aReq.UserID = parseUInt64(val)
	}
	aReq.AttachmentID = parseUInt64(chi.URLParam(r, "attachmentID"))
	aReq.Name = chi.URLParam(r, "name")
	aReq.Kind = chi.URLParam(r, "kind")
	aReq.NamespaceID = parseUInt64(chi.URLParam(r, "namespaceID"))

	return err
}

var _ RequestFiller = NewAttachmentOriginal()

// Attachment preview request parameters
type AttachmentPreview struct {
	AttachmentID uint64 `json:",string"`
	Ext          string
	Kind         string
	NamespaceID  uint64 `json:",string"`
	Sign         string
	UserID       uint64 `json:",string"`
}

func NewAttachmentPreview() *AttachmentPreview {
	return &AttachmentPreview{}
}

func (aReq *AttachmentPreview) Fill(r *http.Request) (err error) {
	if strings.ToLower(r.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(r.Body).Decode(aReq)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return errors.Wrap(err, "error parsing http request body")
		}
	}

	if err = r.ParseForm(); err != nil {
		return err
	}

	get := map[string]string{}
	post := map[string]string{}
	urlQuery := r.URL.Query()
	for name, param := range urlQuery {
		get[name] = string(param[0])
	}
	postVars := r.Form
	for name, param := range postVars {
		post[name] = string(param[0])
	}

	aReq.AttachmentID = parseUInt64(chi.URLParam(r, "attachmentID"))
	aReq.Ext = chi.URLParam(r, "ext")
	aReq.Kind = chi.URLParam(r, "kind")
	aReq.NamespaceID = parseUInt64(chi.URLParam(r, "namespaceID"))
	if val, ok := get["sign"]; ok {

		aReq.Sign = val
	}
	if val, ok := get["userID"]; ok {

		aReq.UserID = parseUInt64(val)
	}

	return err
}

var _ RequestFiller = NewAttachmentPreview()
