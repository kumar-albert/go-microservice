package controllers

import (
	"go-microservice/utils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	w := httptest.NewRecorder()
	var c *gin.Context = utils.GetTestGinContext(w)
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")

	os.Setenv("CONFIG_FILE", "../etc/configuration/local.yaml")

	GetUsers(c)
	assert.EqualValues(t, http.StatusOK, w.Code)
}
