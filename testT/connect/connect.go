package connect

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/go-ini/ini"
	"jenkins_service/utils"
	"log"
)

func OpenJenkins() (Conn *gojenkins.Jenkins) {
	cfg, err := ini.Load("./conf/jenkins.ini")
	if err != nil {

		code := utils.ERROR_JENKINS_CONNECTION_FAIL
		log.Println(code)
		panic(err)
	}

	// 读取配置文件
	jenkinsUrl := cfg.Section("jenkins").Key("url").String()
	jenkinsUser := cfg.Section("jenkins").Key("user").String()
	jenkinsPass := cfg.Section("jenkins").Key("passwd").String()

	// 连接jenkins
	jenkinsConn := gojenkins.CreateJenkins(nil, jenkinsUrl, jenkinsUser, jenkinsPass)

	ctx := context.Background()
	_, err = jenkinsConn.Init(ctx)
	if err != nil {

		code := utils.ERROR_JENKINS_INIT_FAIL
		log.Println(code, err)
	} else {

		code := utils.SUCCESS
		log.Println(code)
	}

	return jenkinsConn
}
