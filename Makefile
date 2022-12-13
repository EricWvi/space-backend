run: FORCE
	go build
	mv space-backend out
	cp config.yaml out
	cd out;./space-backend


FORCE: