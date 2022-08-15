TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

pack:
	docker build -t hello-docker:$(TAG) .

upload:
	docker push hello-docker:$(TAG)

deploy:
	envsubst < k8s.yaml | kubectl apply -f -

ship: pack upload deploy
