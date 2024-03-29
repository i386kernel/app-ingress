build:
	GOOS=linux GOARCH=amd64 go build -o ingresstest .
	chmod 777 ingresstest

imgbuild:
	docker build -t lnanjangud653/ingress-test:"${IMG_VER}" .

imgpush:
	docker push lnanjangud653/ingress-test:${IMG_VER}

