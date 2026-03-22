init env
```
eval $(minikube docker-env)
```

build containers
```
make docker-build-all
```

create all kubernetes resources
```
create-kubernetes
```

start portforwarding 
```
make app-forward-port
```

there are some bugs in the application

Post
```
curl --request POST \
  --url http://localhost:8080/whatis/is
```

Get
```
curl --request GET \
  --url http://localhost:8080/whatis
```