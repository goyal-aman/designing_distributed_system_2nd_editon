init:
	eval $(minikube docker-env)
docker-build-cache:
	docker build -t cache ./cache 
docker-build-app:
	docker build -t app ./app
docker-build-ambassador:
	docker build -t ambassador ./ambassador
docker-build-all:
	docker build -t ambassador ./ambassador
	docker build -t app ./app
	docker build -t cache ./cache 
delete-kubernetes:
	kubectl delete -f .
create-kubernetes:
	kubectl create -f .
app-forward-port:
	kubectl port-forward svc/app 8080:8080