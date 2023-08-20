VERSION = V1.0.1-$(shell git describe --tags --always)
ARTICLE=registry.cn-shenzhen.aliyuncs.com/pg/article
AUTH=registry.cn-shenzhen.aliyuncs.com/pg/auth
APIGW=registry.cn-shenzhen.aliyuncs.com/pg/apigw
IMAGE=$(AUTH):$(VERSION)
.PHONY: all auth article apigw
auth:
	docker buildx build --platform linux/amd64 -f ./auth/Dockerfile --load  -t $(AUTH):$(VERSION) ./auth && docker push $(AUTH):$(VERSION)

article:
	docker buildx build --platform linux/amd64 -f ./article/Dockerfile --load  -t $(ARTICLE):$(VERSION) ./article && docker push $(ARTICLE):$(VERSION)

apigw:
	docker buildx build --platform linux/amd64 -f ./apigw/Dockerfile --load  -t $(APIGW):$(VERSION) ./apigw && docker push $(APIGW):$(VERSION)

build:
	make auth
	make article
	make apigw

.PHONY: run

run:
	export IMAGE=$(AUTH):$(VERSION) && envsubst < ./auth/app.yaml | kubectl apply -f -

	export IMAGE=$(ARTICLE):$(VERSION) && envsubst < ./article/app.yaml | kubectl apply -f -

	export IMAGE=$(APIGW):$(VERSION)  && envsubst < ./apigw/app.yaml | kubectl apply -f -

clean:
	kubectl delete -f ./auth/app.yaml -n pg
	kubectl delete -f ./article/app.yaml -n pg
	kubectl delete -f ./apigw/app.yaml -n pg
