IMAGE=nooocode/usercenter:v1.0.2
run:
	DUBBO_GO_CONFIG_PATH="./dubbogo.yaml" go run main.go
build-image:
	docker build -t ${IMAGE} .