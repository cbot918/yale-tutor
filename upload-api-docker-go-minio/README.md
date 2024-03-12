# Deploy Upload API with Docker feat. Go Monio
一個佈署模板

## Ubuntu Linux/WSL2 Install ( mac 及 win 請參考官方網站)
- minio-server
```bash
下載
wget https://dl.min.io/server/minio/release/linux-amd64/archive/minio_20240310025348.0.0_amd64.deb -O minio.deb

安裝
sudo dpkg -i minio.deb
```

## Start Server
```bash
minio server ~/minio --console-address :9000
```

## Login in to minio UI
```bash
1. 網頁瀏覽 http://localhost:9000
2. 登入 minioadmin / minioadmin
3. 畫面左 Access Keys - 右上角 Create access key - 拿到 Access Key 跟 Secret Key, 放到 .env 裡面
4. 畫面左 Buckets - 右上角 Create Bucket - Bucket Name: testbucket - Create Bucket
```
## Getting Stated
[Github](https://github.com/cbot918/yale-tutor/upload-api-docker-go-minio)
```bash
git clone https://github.com/cbot918/yale-tutor/upload-api-docker-go-minio 

cd upload-api-docker-go-minio
```

### POC of upload png
```bash
go run cmd/minio/minio.go
# should see new picture in minio testbucket
```

### Upload API Service
```bash
go run .
# browse http://localhost:3456
# choose a png file to upload

# browse http://localhost:9000 to check testbucket have new picture just upload
```

### deploy with one click
```
docker-compose up
```
