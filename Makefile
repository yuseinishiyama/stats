VPC_ID=vpc-08c236d1a165e007d
PUBLIC_SUBNET1=subnet-04c56e3a1281a168f
PUBLIC_SUBNET2=subnet-014711872c4205e34
SECURITY_GROUP=sg-01a08b06d78e63c7a

build:
	GOOS=linux go build -o artifact/stats

publish: build
	docker build -t stats .
	docker tag stats:latest 921647845311.dkr.ecr.eu-west-1.amazonaws.com/stats:latest
	docker push 921647845311.dkr.ecr.eu-west-1.amazonaws.com/stats:latest

run:
	go run main.go

run-task:
	aws ecs run-task \
		--cluster default  \
		--task-definition first-run-task-definition \
		--network-configuration awsvpcConfiguration="{subnets=[$(PUBLIC_SUBNET1), $(PUBLIC_SUBNET2)],securityGroups=[$(SECURITY_GROUP)],assignPublicIp=ENABLED}" \
		--enable-execute-command \
		--launch-type FARGATE \
		--platform-version '1.4.0' \
		--profile private

