# REST API

Repositor ini berisi kode yang saya gunakan dalam sesi REST API

## Panduan

Pastikan Golang Version : 1.14 sudah terinstall

### Running

Pada folder yang bersangkutan jalanin perintah berikut :

```
go run model.go handler.go main.go
```

dan tunggu sampai perintah di bawah ini muncul

```
...
[GIN-debug] Listening and serving HTTP on :9001
```

## Frontend

```
http://localhost:9001
```

## Endpoint

GET /material - Get all material
GET /material/:id - Get specific material
POST /material - Create new material
PUT /material/:id - Update material
PATCH /material/:id - Release material
DELETE /material/:id - Delete material
