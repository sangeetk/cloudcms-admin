all:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o cloudcms-admin .

prod:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o cloudcms-admin .
	- docker image rm reg.urantiatech.com/cloudcms/cloudcms-admin 
	docker build -t reg.urantiatech.com/cloudcms/cloudcms-admin .
	docker push reg.urantiatech.com/cloudcms/cloudcms-admin

dev:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o cloudcms-admin .
	- docker image rm localhost:32000/cloudcms/cloudcms-admin 
	docker build -t localhost:32000/cloudcms/cloudcms-admin .
	docker push localhost:32000/cloudcms/cloudcms-admin
