all: help

## build <os>@编译(格式：make build os=linux/darwin/windows,os为可选参数)。
.PHONY:build
build:
	# a:强制重新编译
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o goddd-server


## run@运行服务。
.PHONY:run
run:
	@go run main.go


## clean@清理编译、日志和缓存等数据。
.PHONY:clean
clean:
	@rm -rf ./goddd-server


## commit <msg>@提交Git(格式:make commit msg=备注内容,msg为可选参数)。
.PHONY:commit
message:=$(if $(msg),$(msg),"Rebuilded at $$(date '+%Y年%m月%d日 %H时%M分%S秒')")
commit:
	@echo "\033[0;34mPush to remote...\033[0m"
	@git add .
	@git commit -m $(message)
	@echo "\033[0;31m 💿 Commit完毕\033[0m"


## push <msg>@提交并推送到Git仓库(格式:make push msg=备注内容,msg为可选参数)。
.PHONY:push
push:commit
	@git push #origin master
	@echo "\033[0;31m 📡 Push完毕\033[0m"


## help@查看make帮助。
.PHONY:help
help:Makefile
	@echo "Usage:\n  make [<command> [param]]"
	@echo
	@echo "Available Commands:"
	@sed -n "s/^##//p" $< | column -t -s '@' |grep --color=auto "^[[:space:]][a-z]\+[[:space:]]"
	@echo
	@echo "For more to see https://makefiletutorial.com/"