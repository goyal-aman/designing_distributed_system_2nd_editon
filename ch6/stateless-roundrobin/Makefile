init:
	eval $(minikube docker-env)
docker-build-app:
	docker build -t ch6app .
delete-kubernetes:
	kubectl delete -f .
create-kubernetes:
	kubectl create -f .
app-forward-port:
	kubectl port-forward svc/app 8080:8080
port-forward:
	kubectl port-forward svc/ch6app 8080:8080