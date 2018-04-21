all:
	cd frontend && make all
	cd backend && make all

backend:
	cd backend && make all

frontend:
	cd frontend && make all

tag:
	git config --local user.name "Yannick Roffin"
	git config --local user.email "yroffin@gmail.com"
	git tag "$(shell date +'%Y%m%d-%H%M%S')"
	git push --tags
