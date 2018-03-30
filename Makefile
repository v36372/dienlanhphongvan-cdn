NAME     =rsimgx
MAIN_FILE=server/main.go
GIT_TAG  =v$(shell date +"%Y-%m-%d-%H-%M")
BIN_DIR      =bin
LOG_DIR      =log
LOG_FILE     =$(LOG_DIR)/access.log
PID_API_FILE =$(BIN_DIR)/rsimgx.pid
PID_API      =$(shell cat $(PID_API_FILE))
THIS_FILE := $(lastword $(MAKEFILE_LIST))

default: build

deploy: backup build

backup:
	@echo "STEP: BACKUP"
	@echo "   1. backup: binary file"
	@[ ! -f bin/$(NAME) ] && \
		 echo "   => skip: SERVICE=bin/$(NAME) DOES NOT EXIST" || ( \
		 cp bin/$(NAME) bin/$(NAME)_$(GIT_TAG) && \
	  	 echo "   => ok: SERVICE=bin/$(NAME)_$(GIT_TAG)" )
	@echo "   2. backup: git commit"
	@git tag -f $(GIT_TAG) && \
		echo  "   => ok: TAG=$(GIT_TAG)"

build:
	@echo "STEP: BUILD"
	@echo "   1. create dir: bin" \
		&& mkdir -p bin \
		&& echo "   ==> ok"
	@echo "   2. build: $(MAIN_FILE)" \
		&& go build -ldflags=-s -o bin/$(NAME) $(MAIN_FILE) \
		&& echo "   ==> ok: SERVICE=bin/$(NAME)"

docker:
	@echo "STEP: Build docker image named: imgx"
	@docker build -t imgx .

clean:
	@echo "STEP: CLEAN"
	@echo "   1. remove dir: bin"
	@rm -rf bin \
	 	&& echo "   ==> ok"

restart: 
ifneq ("$(wildcard $(PID_API_FILE))","") 
	@$(MAKE) -f $(THIS_FILE) stop
	@$(MAKE) -f $(THIS_FILE) start
else
	@$(MAKE) -f $(THIS_FILE) start
endif

start:
	@echo "Starting..."
	@mkdir -p $(LOG_DIR)
ifneq ("$(wildcard $(PID_API_FILE))","") 
	@echo "[FAIL] A processing is running. Stop it first or restart"
else
	@pid= nohup bin/$(NAME) >> $(LOG_FILE) 2>&1 & echo "$$!" > $(PID_API_FILE)
	@echo "Done!"
endif

stop:
	@echo "Stopping..."
	@kill -9 $(PID_API)
	@rm $(PID_API_FILE)
	@echo "Done!"

status:
ifeq ("$(wildcard $(PID_API_FILE))","") 
	@echo "There is no running process"
else
	@echo "Service is running at PID=$(PID_API)"
	@echo "---"
	@ps aux | grep $(PID_API)
endif
