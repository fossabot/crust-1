// +build integration-disabled

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/namsral/flag"

	"github.com/crusttech/crust/internal/auth"
	"github.com/crusttech/crust/internal/test"
	systemRepository "github.com/crusttech/crust/system/internal/repository"
	systemTypes "github.com/crusttech/crust/system/types"
)

type (
	jsonResponse struct {
		Error struct {
			Message string `json:"message"`
			Trace   string `json:"trace,omitempty"`
		} `json:"error"`
	}
)

func TestUsers(t *testing.T) {
	ctx := context.Background()

	// we need to set this due to using Init()
	os.Setenv("SYSTEM_DB_DSN", "crust:crust@tcp(crust-db:3306)/crust?collation=utf8mb4_general_ci")

	mountFlags("system", Flags)

	// Initialize routes and exit on failure.
	err := Init(ctx)
	test.Assert(t, err == nil, "Error initializing: %+v", err)

	jwtSecret := "test-secret"

	jwtAuth, err := auth.JWT(jwtSecret, 600)
	test.NoError(t, err, "Error initializing: %v")

	routes := Routes(ctx)

	// Send check request with invalid JWT token.
	{
		req, err := http.NewRequest("GET", "http://127.0.0.1/auth/check", nil)
		test.Assert(t, err == nil, "Error creating request: %+v", err)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  "zblj",
			"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})

		jwtAuth.Encode()

		tokenString, err := token.SignedString([]byte(jwtSecret))
		test.Assert(t, err == nil, "Error creating JWT token: %+v", err)

		req.AddCookie(&http.Cookie{
			Name:     "jwt",
			Value:    tokenString,
			Domain:   ".localhost",
			Expires:  time.Now().Add(time.Hour),
			HttpOnly: true,
			MaxAge:   50000,
			Path:     "/auth",
		})

		recorder := httptest.NewRecorder()
		routes.ServeHTTP(recorder, req)
		resp := recorder.Result()

		fmt.Println(">>> (request)")
		fmt.Println(request(req))
		fmt.Println("----")
		fmt.Println("<<< (response)")
		fmt.Println(response(resp))

		jr, err := decodeJson(resp.Body)
		test.Assert(t, err == nil, "Error decoding response body: %+v", err)
		test.Assert(t, jr.Error.Message == "failed to authorize request: signature is invalid", "Expected error 'failed to authorize request: signature is invalid' got: %+v", jr.Error.Message)
	}

	// Send "Login" request without parameters.
	{
		req, err := http.NewRequest("POST", "http://localhost/auth/login", nil)
		test.Assert(t, err == nil, "Error creating request: %+v", err)

		recorder := httptest.NewRecorder()
		routes.ServeHTTP(recorder, req)
		resp := recorder.Result()

		fmt.Println(">>> (request)")
		fmt.Println(request(req))
		fmt.Println("----")
		fmt.Println("<<< (response)")
		fmt.Println(response(resp))

		jr, err := decodeJson(resp.Body)
		test.Assert(t, err == nil, "Error decoding response body: %+v", err)
		test.Assert(t, jr.Error.Message == "missing form body", "Expected error 'missing form body' got: %+v", jr.Error.Message)
	}

	// Send "Login" request with missing user.
	{
		jsonStr := `{"username":"test123","password":"test123"}`

		req, err := http.NewRequest("POST", "http://localhost/auth/login", strings.NewReader(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		test.Assert(t, err == nil, "Error creating request: %+v", err)

		recorder := httptest.NewRecorder()
		routes.ServeHTTP(recorder, req)
		resp := recorder.Result()

		fmt.Println(">>> (request)")
		fmt.Println(request(req))
		fmt.Println("----")
		fmt.Println("<<< (response)")
		fmt.Println(response(resp))

		jr, err := decodeJson(resp.Body)
		test.Assert(t, err == nil, "Error decoding response body: %+v", err)
		test.Assert(t, jr.Error.Message == "crust.auth.repository.UserNotFound", "Expected error 'crust.auth.repository.UserNotFound' got: %+v", jr.Error.Message)
	}

	// Create user.
	user := &systemTypes.User{
		ID:       1337,
		Username: "johndoe",
	}
	{
		userAPI := systemRepository.User(context.Background(), nil)
		_, err := userAPI.Create(user)
		test.Assert(t, err == nil, "Error when inserting user: %+v", err)
	}

	// Send "Login" request with existing user.
	if false {
		form := url.Values{}
		form.Add("username", "johndoe")
		form.Add("password", "johndoe123")

		req, err := http.NewRequest("POST", "http://localhost/auth/login", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		test.Assert(t, err == nil, "Error creating request: %+v", err)

		recorder := httptest.NewRecorder()
		routes.ServeHTTP(recorder, req)
		resp := recorder.Result()

		fmt.Println(">>> (request)")
		fmt.Println(request(req))
		fmt.Println("----")
		fmt.Println("<<< (response)")
		fmt.Println(response(resp))

		c := resp.Cookies()
		test.Assert(t, len(c) == 1, "Expected 1 cookie value, got: %+v", len(c))
		test.Assert(t, c[0].Value != "", "Expected non empty jwt token, got: %+v", c[0].Value)

		type jsonResponse struct {
			Response struct {
				UserID   string `json:"userID"`
				Username string `json:"username"`
			} `json:"response"`
		}

		var jr jsonResponse
		err = json.NewDecoder(resp.Body).Decode(&jr)
		test.Assert(t, err == nil, "Error decoding response body: %+v", err)
		test.Assert(t, jr.Response.UserID != "0", "Expected userID not to be 0, got: %+v", jr.Response.UserID)
		test.Assert(t, jr.Response.Username == "johndoe", "Expected username 'johndoe', got: %+v", jr.Response.Username)

		// Check JWT token after successful login.
		req, err = http.NewRequest("GET", "http://localhost/auth/check", nil)
		test.Assert(t, err == nil, "Error creating request: %+v", err)

		routes.ServeHTTP(recorder, req)

		resp = recorder.Result()

		fmt.Println(">>> (request)")
		fmt.Println(request(req))
		fmt.Println("----")
		fmt.Println("<<< (response)")
		fmt.Println(response(resp))

		test.Assert(t, resp.StatusCode == 200, "Expected http status code 200, got: %+v", resp.StatusCode)
		test.Assert(t, len(c) == 1, "Expected 1 cookie value, got: %+v", len(c))
		test.Assert(t, c[0].Value != "", "Expected non empty jwt token, got: %+v", c[0].Value)
	}

	// Send "Login" request with existing user.
	{
		jsonStr := `{"username": "johndoe", "password": "johndoe123"}`

		req, err := http.NewRequest("POST", "http://localhost/auth/login", strings.NewReader(jsonStr))
		req.Header.Add("Content-Type", "application/json")

		test.Assert(t, err == nil, "Error creating request: %+v", err)

		recorder := httptest.NewRecorder()
		routes.ServeHTTP(recorder, req)
		resp := recorder.Result()

		fmt.Println(">>> (request)")
		fmt.Println(request(req))
		fmt.Println("----")
		fmt.Println("<<< (response)")
		fmt.Println(response(resp))

		c := resp.Cookies()
		test.Assert(t, len(c) == 1, "Expected 1 cookie value, got: %+v", len(c))
		test.Assert(t, c[0].Value != "", "Expected non empty jwt token, got: %+v", c[0].Value)

		type jsonResponse struct {
			Response struct {
				UserID   string `json:"userID"`
				Username string `json:"username"`
			} `json:"response"`
		}

		var jr jsonResponse
		err = json.NewDecoder(resp.Body).Decode(&jr)
		test.Assert(t, err == nil, "Error decoding response body: %+v", err)
		test.Assert(t, jr.Response.UserID != "0", "Expected userID not to be 0, got: %+v", jr.Response.UserID)
		test.Assert(t, jr.Response.Username == "johndoe", "Expected username 'johndoe', got: %+v", jr.Response.Username)

		// Check JWT token after successful login.
		req, err = http.NewRequest("GET", "http://localhost/auth/check", nil)
		test.Assert(t, err == nil, "Error creating request: %+v", err)

		routes.ServeHTTP(recorder, req)

		resp = recorder.Result()

		fmt.Println(">>> (request)")
		fmt.Println(request(req))
		fmt.Println("----")
		fmt.Println("<<< (response)")
		fmt.Println(response(resp))

		test.Assert(t, resp.StatusCode == 200, "Expected http status code 200, got: %+v", resp.StatusCode)
		test.Assert(t, len(c) == 1, "Expected 1 cookie value, got: %+v", len(c))
		test.Assert(t, c[0].Value != "", "Expected non empty jwt token, got: %+v", c[0].Value)
	}

	// Send "Logout" request and expect empty jwt token.
	{
		req, err := http.NewRequest("GET", "http://localhost/auth/logout", nil)
		test.Assert(t, err == nil, "Error creating request: %+v", err)

		recorder := httptest.NewRecorder()
		routes.ServeHTTP(recorder, req)
		resp := recorder.Result()

		fmt.Println(">>> (request)")
		fmt.Println(request(req))
		fmt.Println("----")
		fmt.Println("<<< (response)")
		fmt.Println(response(resp))

		c := resp.Cookies()

		test.Assert(t, resp.StatusCode == 200, "Expected http status code 200, got: %+v", resp.StatusCode)
		test.Assert(t, len(c) == 1, "Expected 1 cookie value, got: %+v", len(c))
		test.Assert(t, c[0].Value == "", "Expected empty jwt token, got: %+v", c[0].Value)
	}

	// Send check request without JWT token.
	{
		req, err := http.NewRequest("GET", "http://127.0.0.1/auth/check", nil)
		test.Assert(t, err == nil, "Error creating request: %+v", err)

		recorder := httptest.NewRecorder()
		routes.ServeHTTP(recorder, req)

		resp := recorder.Result()

		fmt.Println(">>> (request)")
		fmt.Println(request(req))
		fmt.Println("----")
		fmt.Println("<<< (response)")
		fmt.Println(response(resp))

		jr, err := decodeJson(resp.Body)
		test.Assert(t, err == nil, "Error decoding response body: %+v", err)
		test.Assert(t, jr.Error.Message == "http: named cookie not present", "Expected error 'http: named cookie not present' got: %+v", jr.Error.Message)
	}
}

func request(req *http.Request) string {
	b, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return ">>> Error: " + err.Error()
	}
	if b != nil {
		return strings.TrimSpace(string(b))
	}
	return ""
}

func response(resp *http.Response) string {
	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return "<<< Error: " + err.Error()
	}
	if b != nil {
		return strings.TrimSpace(string(b))
	}
	return ""
}

func mountFlags(prefix string, mountFlags ...func(...string)) {
	for _, mount := range mountFlags {
		mount(prefix)
	}
	flag.Parse()
}

func decodeJson(r io.Reader) (jsonResponse, error) {
	var ret jsonResponse
	err := json.NewDecoder(r).Decode(&ret)
	if err != nil {
		return jsonResponse{}, err
	}
	return ret, nil
}
