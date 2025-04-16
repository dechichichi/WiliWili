# 辅助工具安装列表
# 执行 go install github.com/vektra/mockery/v2@v2.52.1
# 执行 go install github.com/cloudwego/hertz/cmd/hz@latest
# 执行 go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
# 执行 go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
#项目MODULE名
MODULE=wiliwili
# 检查 tmux 是否存在
TMUX_EXISTS := $(shell command -v tmux)
# 当前架构
PREFIX = "[Makefile]"

#目录相关
DIR=$(shell pwd)
CMD = $(DIR)/cmd
CONFIG_PATH = $(DIR)/config
IDL_PATH=${DIR}/idl
OUTPUT_PATH = $(DIR)/output
API_PATH= $(DIR)/cmd/api
APP_PATH = $(DIR)/app
# 服务名
SERVICES := gateway user video like comment chat
service = $(word 1, $@)

EnvironmentStartEnv=WILIWILI_ENVIRONMENT_STARTED
EnvironmentStartFlag=true
EtcdAddrEnv=ETCD_ADDR
EtcdAddr=127.0.0.1:2379

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

.PHONY: init
init:
	@echo "init"
	@golangci-lint run 

# 定义 mock-% 规则
.PHONY: mock-%
mock-%:
	@mockery --name $$(echo $* | awk '{print toupper(substr($$0, 1, 1)) substr($$0, 2) "UseCase"}') --dir ${DIR}/app/$*/usecase --output ${DIR}/app/$*/usecase/mocks
	@mockery --name $$(echo $* | awk '{print toupper(substr($$0, 1, 1)) substr($$0, 2) "DB"}') --dir ${DIR}/app/$*/domain/repository --output ${DIR}/app/$*/usecase/mocks
.PHONY: $(SERVICES)
$(SERVICES):
	@if [ -z "$(TMUX_EXISTS)" ]; then \
		echo "$(PREFIX) tmux is not installed. Please install tmux first."; \
		exit 1; \
	fi
	@if [ -z "$$TMUX" ]; then \
		echo "$(PREFIX) you are not in tmux, press ENTER to start tmux environment."; \
		read -r; \
		if tmux has-session -t fzuhelp 2>/dev/null; then \
			echo "$(PREFIX) Tmux session 'fzuhelp' already exists. Attaching to session and running command."; \
			tmux attach-session -t fzuhelp; \
			tmux send-keys -t fzuhelp "make $(service)" C-m; \
		else \
			echo "$(PREFIX) No tmux session found. Creating a new session."; \
			tmux new-session -s fzuhelp "make $(service); $$SHELL"; \
		fi; \
	else \
		echo "$(PREFIX) Build $(service) target..."; \
		mkdir -p output; \
		bash $(DIR)/docker/script/build.sh $(service); \
		echo "$(PREFIX) Build $(service) target completed"; \
	fi
ifndef BUILD_ONLY
	@echo "$(PREFIX) Automatic run server"
	@if tmux list-windows -F '#{window_name}' | grep -q "^wiliwili-$(service)$$"; then \
		echo "$(PREFIX) Window 'wiliwili-$(service)' already exists. Reusing the window."; \
		tmux select-window -t "wiliwili-$(service)"; \
	else \
		echo "$(PREFIX) Window 'wiliwili-$(service)' does not exist. Creating a new window."; \
		tmux new-window -n "wiliwili-$(service)"; \
		tmux split-window -h ; \
		tmux select-layout -t "wiliwili-$(service)" even-horizontal; \
	fi
	@echo "$(PREFIX) Running $(service) service in tmux..."
	@tmux send-keys -t wiliwili-$(service).0 'export SERVICE=$(service) && bash ./docker/script/entrypoint.sh' C-m
	@tmux select-pane -t wiliwili-$(service).1
endif