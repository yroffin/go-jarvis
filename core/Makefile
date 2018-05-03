WHERE=$(PWD)
MODULE_ARMEL=${WHERE}/go-jarvis-0.0.1-SNAPSHOT.armel
MODULE_ARMHF=${WHERE}/go-jarvis-0.0.1-SNAPSHOT.armhf
MODULE_LOCAL=${WHERE}/go-jarvis-0.0.1-SNAPSHOT.bin

all: clean ${MODULE_ARMEL} ${MODULE_ARMHF}

live:
	gin -p 3000 -a 8080 -i -- -http=8080 -Djarvis.slack.api=$SLACK_API -Djarvis.zway.password=$ZWAY_PASSWORD

local: clean ${MODULE_LOCAL}

clean:
	rm -rf ${WHERE}/armel
	rm -rf ${WHERE}/armhf
	rm -f ${MODULE_ARMHF}
	rm -f ${MODULE_ARMEL}
	sh dist.sh
	mkdir -p dist/version
	sh info.sh > dist/version/version.json
	go get -u github.com/gobuffalo/packr/...
	packr

${MODULE_LOCAL}:
	# module
	go install -installsuffix bin
	ls -lrt ${GOPATH}/bin
	mv -f ${GOPATH}/bin/core ${GOPATH}/go-jarvis-0.0.1-SNAPSHOT.bin

${MODULE_ARMEL}:
	# module
	CC=arm-linux-gnueabi-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 CGO_CFLAGS="-march=armv6j -mfloat-abi=soft" go install -installsuffix armel
	mv -f ${GOPATH}/bin/linux_arm/core ${GOPATH}/go-jarvis-0.0.1-SNAPSHOT.armel

${MODULE_ARMHF}:
	# module
	CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CGO_CFLAGS="" go install -installsuffix armhf
	mv -f ${GOPATH}/bin/linux_arm/core ${GOPATH}/go-jarvis-0.0.1-SNAPSHOT.armhf

tag:
	git config --local user.name "Yannick Roffin"
	git config --local user.email "yroffin@gmail.com"
	git tag "$(shell date +'%Y%m%d-%H%M%S')"
	git push --tags
