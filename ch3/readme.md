topz container runs as size car. topz exposes /topz api which returns running processes.

```bash
docker run nginx:latest
```

save nginx container id in APP_ID
export APP_ID=<hash>
build localtopz client 
```bash
docker build -t localtopz:latest .
```

```
docker run --pid=container:${APP_ID} -p 8080:8080 localtopz:latest
```

curl localhost:8080/topz
