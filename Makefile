
image:
	docker build -t wayha-sms-api:v1.0.0 -f Dockerfile .

container:
	docker run -p:9000:9000 --env-file ./cmd/sms/local.env \
	--name wayha-sms-api wayha-sms-api:v1.0.0