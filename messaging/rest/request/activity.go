package request

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `activity.go`, `activity.util.go` or `activity_test.go` to
	implement your API calls, helper functions and tests. The file `activity.go`
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

// Activity send request parameters
type ActivitySend struct {
	ChannelID uint64 `json:",string"`
	MessageID uint64 `json:",string"`
	Kind      string
}

func NewActivitySend() *ActivitySend {
	return &ActivitySend{}
}

func (aReq *ActivitySend) Fill(r *http.Request) (err error) {
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

	if val, ok := post["channelID"]; ok {

		aReq.ChannelID = parseUInt64(val)
	}
	if val, ok := post["messageID"]; ok {

		aReq.MessageID = parseUInt64(val)
	}
	if val, ok := post["kind"]; ok {

		aReq.Kind = val
	}

	return err
}

var _ RequestFiller = NewActivitySend()
