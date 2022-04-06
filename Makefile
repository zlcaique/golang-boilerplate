
ifeq ($(OS),Windows_NT)
    OS += windowns
    ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
        PROCESSOR += -amd64
    else
        ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
            PROCESSOR += -amd64
        endif
        ifeq ($(PROCESSOR_ARCHITECTURE),x86)
            PROCESSOR += -386
        endif
    endif
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        OS += linux
    endif
    ifeq ($(UNAME_S),Darwin)
        PROCESSOR += darwin
    endif
    UNAME_P := $(shell uname -p)
    ifeq ($(UNAME_P),x86_64)
        PROCESSOR += amd64
    endif
    ifneq ($(filter %86,$(UNAME_P)),)
        PROCESSOR += 386
    endif
    ifneq ($(filter arm%,$(UNAME_P)),)
        PROCESSOR += arm64
    endif
endif

install-migrate:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.$(OS)-$(PROCESSOR).tar.gz | tar xvz