DOCKER_REGISTORY = 921647845311.dkr.ecr.eu-west-1.amazonaws.com
IMAGE = $(DOCKER_REGISTORY)/stats

build-arm:
	docker build --build-arg GOARCH=arm64 --build-arg GOARM=7 -t stats:arm64v7 .

build-amd:
	docker build --build-arg GOARCH=amd64 -t stats:amd64 .

build: build-arm build-amd

publish: build
	aws ecr get-login-password --region eu-west-1 --profile private | \
		docker login --username AWS --password-stdin $(DOCKER_REGISTORY)
	docker tag stats:amd64 $(IMAGE):amd64
	docker tag stats:arm64v7 $(IMAGE):arm64v7
	docker push $(IMAGE):amd64
	docker push $(IMAGE):arm64v7

manifest:
	docker manifest create $(IMAGE) $(IMAGE):amd64 $(IMAGE):arm64v7
	docker manifest annotate $(IMAGE) $(IMAGE):arm64v7 --os linux --arch arm
	docker manifest push $(IMAGE)

run-worker: build
	docker run --rm stats:latest worker

run-bot: build
	docker run --rm --env-file .env -p 80:80 stats:latest bot
