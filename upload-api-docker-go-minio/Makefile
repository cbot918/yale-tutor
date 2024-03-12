IMAGE=upload-app
CONTAINER=upload
CLEAN_VOLUME=upload-api-docker-go-minio_minio-data
### test minio
minio-check:
	go run cmd/minio/minio.go

minio-start:
	mc admin service start myminio

minio-stop:
	mc admin service stop myminio

### docker
docker-build: dockerfile
	docker build -t $(IMAGE) .
docker-run: 
	docker run -it --name $(CONTAINER) -p 3456:3456 $(IMAGE)
docker-exec:
	docker exec -it $(CONTAINER) /bin/sh

docker-clean:
	docker container rm $(CONTAINER)
	docker image rm $(IMAGE)

docker-clean-volume:
	docker volume rm $(CLEAN_VOLUME)

.PHONY: minio
.SILENT: minio