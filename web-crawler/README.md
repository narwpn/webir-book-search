# book-search

## Web Crawler

### Environment Variables
```
CRAWLER_THREADS=8

REDIS_HOST=localhost:6379
REDIS_PASSWORD=1q2w3e4r
REDIS_DB=0

MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=<your-access-key>
MINIO_SECRET_KEY=<your-secret-key>
MINIO_BUCKET=html

POSTGRES_DSN='host=localhost user=admin password=1q2w3e4r dbname=book_search port=5432 sslmode=disable TimeZone=Asia/Shanghai'
```

### Run
```
cd web-crawler
go mod tidy
go run .
```
