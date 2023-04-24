package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func ParseYaml(file string, output map[interface{}]interface{}) {
	yfile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	err2 := yaml.Unmarshal(yfile, &output)
	if err2 != nil {
		log.Fatal(err2)
	}

}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}
