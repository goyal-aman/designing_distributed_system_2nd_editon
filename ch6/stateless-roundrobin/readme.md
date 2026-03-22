demos roundrobins between replicas of service called ch5app
build image
```
make docker-build-app
```

create kubernetes resources (ch6app replicas and ch6app service)
```
make create-kubernetes
```

open logs of all three pods
```
kubectl get pods
kubectl logs <pod1> -f
kubectl logs <pod2> -f 
kubectl logs <pod3> -f
```


in my machine the port is not available on localhost, portforwarding only forwards traffic to single pod. so we test round robin by make request in kubernetes env
```
kubectl run --rm -it debug --image=busybox -- sh

# to get into existing pod
kubectl exec -it <podname> -- sh
```
make multiple calls call
```
wget http://ch6app.default.svc.cluster.local:8080
```
you should see logs in all three services one by one

clean the resources
```
make delete-kubernetes 
```