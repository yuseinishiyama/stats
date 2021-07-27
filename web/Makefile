build:
	GOOS=linux go build -o artifact/ ./cmd/stats
	docker build -t stats .

publish: build
	docker tag stats:latest 921647845311.dkr.ecr.eu-west-1.amazonaws.com/stats:latest
	docker push 921647845311.dkr.ecr.eu-west-1.amazonaws.com/stats:latest

run-worker: build
	docker run --rm stats:latest worker

run-bot: build
	docker run --rm --env-file .env -p 80:80 stats:latest bot
