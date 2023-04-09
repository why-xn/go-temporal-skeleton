## Go Gin & Temporal Skeleton

**Init Docker Environment**
```
make init
```
\
**Build Temporalite Image**
```
make build-temporalite
```
\
**Start Temporalite**
```
make start-temporalite
```
\
**Stop Temporalite**
```
make stop-temporalite
```
\
**Build App**
```
go build
```
\
**Start App**
```
./go-temporal-skeleton
```


### Test APIs
GET: http://localhost:8080/api/v1/students \
GET: http://localhost:8080/api/v1/students/123456