package controller

import (
	"github.com/Anderson-Lu/gofasion/gofasion"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"jenkins_service/module"
	"jenkins_service/utils"
	"net/http"
)

func GetJobBuildStatus(c *gin.Context) {
	app := utils.Gin{C: c}
	dataByte, _ := ioutil.ReadAll(c.Request.Body)
	gf := gofasion.NewFasion(string(dataByte))

	// 解析前端json参数
	jobName := gf.Get("jobName").ValueStr()
	parentIDs := gf.Get("parentIDs").ValueStr()

	err, result := module.GetJenkinsBuildStatus(jobName, parentIDs)
	if err != nil {

		app.Response(http.StatusInternalServerError, utils.ERROR_GET_JENKINS_JOB_STATUS, nil)
		return
	}

	app.Response(http.StatusOK, utils.SUCCESS, result)
}

func GetJobBuildLogs(c *gin.Context) {
	app := utils.Gin{
		C: c}
	dataByte, _ := ioutil.ReadAll(c.Request.Body)
	gf := gofasion.NewFasion(string(dataByte))

	// 解析前端json参数
	jobName := gf.Get("jobName").ValueStr()
	parentIDs := gf.Get("parentIDs").ValueStr()

	err, result := module.GetJenkinsBuildLogs(jobName, parentIDs)
	if err != nil {

		app.Response(http.StatusInternalServerError, utils.ERROR_GET_JENKINS_JOB_LOGS, nil)
		return
	}
	app.Response(http.StatusOK, utils.SUCCESS, result)
}
