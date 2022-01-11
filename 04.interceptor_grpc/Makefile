# Makefile

.PHONY: proto
proto:	
	#for x in **/*.proto; do protoc --go_out=plugins=grpc,paths=source_relative:. $x; done
	protoc --proto_path=. --go_out=plugins=grpc,paths=source_relative:. ./proto/hello/*.proto

# --proto_path 或 -I 指定 protoc 的工作目录
# --go_out 选项是用来指定 protoc-gen-go 插件的工作方式 和 go 代码目录架构的生成位置的模式（有 3 种模式，如 source_relative 模式），可以向 --go_out 传递多参数
# 主要的两个参数为 plugins 和 paths ，代表 生成 go 代码所使用的插件 和 生成的 go 代码的目录怎样架构
# plugins=grpc 表示会将 proto 文件中指定的服务编译为接口代码
# paths=source_relative:. 表示生成的 .pb.go 文件将放置在与输入文件相同的相对目录中

# 注意：因为，protoc-gen-go 不支持多包同时编译，如果觉得麻烦，可以执行脚本（**/* 代表递归获取当前目录下所有的文件和文件夹）：
# for x in **/*.proto; do protoc --go_out=plugins=grpc,paths=source_relative:. $x; done

