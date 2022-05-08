HELM_PLUGIN_DIR ?= $(shell helm env | grep HELM_PLUGINS | cut -d\" -f2)/helm-etcd

all: build install

build:
	./hack/gobuild.sh

install:
	mkdir -p $(HELM_PLUGIN_DIR)
	cp -f plugin.yaml $(HELM_PLUGIN_DIR)/plugin.yaml
	cp -f bin/helm-etcd $(HELM_PLUGIN_DIR)/helm-etcd
	cp -f config.json $(HELM_PLUGIN_DIR)/config.json