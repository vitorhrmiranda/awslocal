setup:
	mkdir -p bin
	go build -o bin/main lambda/user_gen_password/main.go
	zip main.zip bin/main
	./setup.sh
	rm main.zip

notify:
	awslocal sqs send-message \
		--queue-url http://awslocal:4566/000000000000/consumer \
		--message-body file://input.json

scan:
	awslocal dynamodb scan --table-name=users

reset:
	docker-compose restart awslocal
