Mode ?= DEBUG

run: FORCE
	go build
	mv space-backend out
ifeq ($(Mode), Release)
	cp config.deploy.yaml out/config.yaml
endif
ifeq ($(Mode), DEBUG)
	cp config.yaml out/config.yaml
	rm out/log/space.log
	touch out/log/space.log
endif
	cd out;./space-backend

FORCE: