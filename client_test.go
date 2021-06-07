package prismacloudcompute

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

var GenericAuthFile = []byte(`{
	"url":"localhost",
	"username":"YOUR_USERNAME",
	"password":"YOUR_PASSWORD",
	"port":8083,
	"skip_ssl_cert_verification":true
}`)

var LoginBody = `{
    "message": "login_successful",
    "token": "testJwtToken"
}`

func init() {
	log.SetFlags(0)
}

func rl() {
	log.SetOutput(os.Stderr)
}

func OkResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: 200,
		Body: ioutil.NopCloser(
			strings.NewReader(body),
		),
	}
}

func ErrorResponse(code int, body string, e interface{}) *http.Response {
	ans := &http.Response{
		StatusCode: code,
		Body: ioutil.NopCloser(
			strings.NewReader(body),
		),
	}

	if e != nil {
		errBody, _ := json.Marshal(e)
		ans.Header = map[string][]string{
			"X-Redlock-Status": []string{
				"[" + string(errBody) + "]",
			},
		}
	}

	return ans
}

func MockClient(responses []*http.Response) Client {
	c := Client{
		pcAuthFileContent: GenericAuthFile,
	}

	c.pcResponses = make([]*http.Response, 0, len(responses)+1)
	c.pcResponses = append(c.pcResponses, OkResponse(LoginBody))
	c.pcResponses = append(c.pcResponses, responses...)

	if len(responses) > 0 {
		_ = c.Initialize("creds.json")
	}

	return c
}

func TestLogin(t *testing.T) {
	c := MockClient(nil)
	err := c.Initialize("creds.json")

	if err != nil {
		t.Fail()
	}
}

func TestInitializeSetsJwt(t *testing.T) {
	c := MockClient(nil)
	_ = c.Initialize("creds.json")

	if c.JsonWebToken == "" {
		t.Fail()
	}
}

func TestInitializeReadsDefaults(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := MockClient(nil)
	_ = c.Initialize("creds.json")

	if c.Protocol == "" {
		t.Fail()
	}

	if c.Timeout == 0 {
		t.Fail()
	}

	if len(c.Logging) == 0 {
		t.Fail()
	}

	s := buf.String()
	if s == "" {
		t.Fail()
	}
}

func TestLogAction(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := MockClient(nil)
	c.Logging = map[string]bool{LogAction: true}
	c.Log(LogAction, "ok")
	s := buf.String()
	if s != "ok\n" {
		t.Fail()
	}
}

func TestLogActionDisabled(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer rl()

	c := MockClient(nil)
	c.Logging = map[string]bool{LogQuiet: true}
	c.Log(LogAction, "ok")
	s := buf.String()
	if s != "" {
		t.Fail()
	}
}
