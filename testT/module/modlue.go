package module

import (
	"context"
	"log"
	"simone_service/connect"
)

// 获取jenkins job build状态
func GetJenkinsBuildStatus(jobName string, parentIDs string) (error, string) {
	var status string
	ctx := context.Background()
	jenkinsConn := connect.OpenJenkins()

	// 判断JobName和目录都不为空执行如下流程
	if parentIDs != "" && jobName != "" {

		build, err := jenkinsConn.GetJob(ctx, jobName, parentIDs)
		if err != nil {

			return err, "JOB NOT EXIST"
		} else {

			info, _ := build.GetLastBuild(ctx)
			status = info.GetResult()
			if status != "" {

				return nil, status
			}
			return nil, "RUNNING"
		}
	} else {
		// 否则走如下流程
		build, err := jenkinsConn.GetJob(ctx, jobName)
		log.Println("GetJenkinsBuildStatus err =========>", err)
		if err != nil {

			return err, "JOB NOT EXIST"
		} else {

			// 获取最新的build
			info, err := build.GetLastBuild(ctx)
			if err != nil {

				return err, "JOB NOT EXIST"
			}
			// 获取最新build的状态
			status = info.GetResult()
			if status != "" {

				return nil, status
			}
			// 这里做了处理，当Job正在运行时，status会返回空值，这里直接返回RUNNING状态
			return nil, "RUNNING"
		}
	}
}

// 获取jenkins job build日志
func GetJenkinsBuildLogs(jobName string, parentIDs string) (error, string) {
	var consoleOutput string
	ctx := context.Background()
	jenkinsConn := connect.OpenJenkins()

	// 判断JobName和目录都不为空执行如下流程
	if parentIDs != "" && jobName != "" {

		build, err := jenkinsConn.GetJob(ctx, jobName, parentIDs)
		if err != nil {

			return err, "JOB NOT EXIST"
		} else {

			// 获取最新的build
			info, _ := build.GetLastBuild(ctx)
			// 获取最新build输出内容
			consoleOutput = info.GetConsoleOutput(ctx)

			return nil, consoleOutput
		}
	} else {

		build, err := jenkinsConn.GetJob(ctx, jobName)
		if err != nil {

			return err, "JOB NOT EXIST"
		} else {

			info, err := build.GetLastBuild(ctx)
			if err != nil {

				return err, "JOB NOT EXIST"
			}
			consoleOutput = info.GetConsoleOutput(ctx)

			return nil, consoleOutput
		}
	}
}
