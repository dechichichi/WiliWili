# 辅助工具安装列表
# 执行 go install github.com/cloudwego/hertz/cmd/hz@latest
# 执行 go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

#项目MODULE名
MODULE=wiliwili


#目录相关
# 项目根目录
DIR=$(shell pwd)
IDL_PATH=${DIR}/idl

## --------------------------------------
## 构建与调试
## --------------------------------------

# 启动必要的环境，比如 mysql redis
.PHONY: env-up
env-up:
	@ docker compose -f ./docker/docker-compose.yml up -d

# 关闭必要的环境，但不清理 data（位于 docker/data 目录中）
.PHONY: env-down
env-down:
	@ cd ./docker && docker compose down

# 基于 idl 生成相关的 go 语言描述文件
.PHONY: kitex-gen-%
kitex-gen-%:
	@ kitex -module "${MODULE}" \
		-thrift no_default_serdes \
		${IDL_PATH}/$*.thrift
	@ go mod tidy

# 生成基于 Hertz 的脚手架
.PHONY: hz-%
hz-%:
	hz update -idl ${IDL_PATH}/api/$*.thrift