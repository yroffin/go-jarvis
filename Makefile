all:
	cd frontend && make all
	cd backend && make all

backend:
	cd backend && make all

frontend:
	cd frontend && make all

TAG=$(shell date +'%Y%m%d-%H%M%S')

tag:
	# prepare tag
	git config --local user.name "Yannick Roffin"
	git config --local user.email "yroffin@gmail.com"
	git tag $(TAG)
	git push --tags
	echo $(TAG) > .TAG

