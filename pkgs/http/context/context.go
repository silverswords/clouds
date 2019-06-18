package core

import (
	"encoding/json"
	"net/http"
)

var (
	jsonContentType = []string{"application/json; charset=utf-8"}
)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

// Context -
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

// NewContext -
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request:  r,
		Response: w,
	}
}

// WriteJSON ResponseWriter
func (c *Context) WriteJSON(code int, obj interface{}) {
	c.Response.WriteHeader(code)

	header := c.Response.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = jsonContentType
	}

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	c.Response.Write(jsonBytes)
}

// BindJSON -
func (c *Context) BindJSON(obj interface{}) error {
	decoder := json.NewDecoder(c.Request.Body)

	err := decoder.Decode(obj)

	return err
}
