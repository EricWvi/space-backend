Mode ?= DEBUG
PORT_NUM := $(shell lsof -i | grep 8719 | awk -F ' ' '{print $$2}')

run: FORCE
	go build
	mv space-backend out
	cp config.yaml out/config.yaml
	rm out/log/space.log
	touch out/log/space.log
	cd out;./space-backend

deploy: FORCE
ifneq ($(PORT_NUM),)
	kill -9 $(PORT_NUM)
endif
	go build
	mv space-backend /home/test/space/back-deploy
	cp config.deploy.yaml /home/test/space/back-deploy/config.yaml
	export GIN_MODE=release;cd /home/test/space/back-deploy;./space-backend > gin.log &

FORCE:
