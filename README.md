# TestJavan

## How to use

```
- Create .env file and copy env key from .env.example then fill it with your own value
- run go get ./...
- run go run *.go
```

# API HTTP Documentation

Header

```
{
    "Content-Type": "application/json",
    "Member-ID": "1",
    "Device-Type": "desktop" (ex: desktop, android, ios)
    "Device-Token": "token-123"
}
```

## Get Family By ID

### TCP Message args

### TCP Message args

Payload (JSON String)

```
{
    "method": "GET",
    "category": "family",
    "args": {
        "member_id": 1
    }
}
```

### HTTP GET /family/member/:id

Result

```
{
    "data": {
        "member_id": 1,
        "member_name": "Banban",
        "gender": "F",
        "created_at": "2023-01-14T16:23:53.858527Z",
        "updated_at": "2023-01-15T02:09:11.192971Z",
        "deleted_at": null,
        "assets": null
    },
    "error": {
        "message": "",
        "status_code": 0
    },
    "status": "success"
}
```

## Create New Family Member

### POST /family/member

Body

```
{
    "member_name": "New User 1",
    "gender": "M"
}
```

Result

```
{
    "data": null,
    "error": {
        "message": "",
        "status_code": 0
    },
    "status": "success"
}
```

## Update Family Member By ID

### PUT /family/member/:id

Body

```
{
    "member_name": "Update user 1",
    "gender": "F"
}
```

Result

```
{
    "data": null,
    "error": {
        "message": "",
        "status_code": 0
    },
    "status": "success"
}
```

## Delete Family Member By ID

### DELETE /family/member/:id

Result

```
{
    "data": null,
    "error": {
        "message": "",
        "status_code": 0
    },
    "status": "success"
}
```

## Get Asset By ID

### GET /asset/:id

Result

```
{
    "data": {
        "asset_id": 1,
        "asset_name": "Samsung Universe 9",
        "created_at": "2023-01-15T09:32:35.849754Z",
        "updated_at": "2023-01-15T11:48:59.613508Z",
        "deleted_at": null
    },
    "error": {
        "message": "",
        "status_code": 0
    },
    "status": "success"
}
```

## Create New Asset

### POST /asset

Body

```
{
    "asset_name": "Asset baru 1"
}
```

Result

```
{
    "data": null,
    "error": {
        "message": "",
        "status_code": 0
    },
    "status": "success"
}
```

## Update Asset By ID

### PUT /asset/:asset_id

Body

```
{
    "asset_name": "Update Asset baru 1"
}
```

Result

```
{
    "data": null,
    "error": {
        "message": "",
        "status_code": 0
    },
    "status": "success"
}
```

## Delete Asset By ID

### DELETE /asset/:asset_id

Result

```
{
    "data": null,
    "error": {
        "message": "",
        "status_code": 0
    },
    "status": "success"
}
```

## Add Asset To Family Member

### POST /family/member/:family_member_id/asset

Body

```
{
    "asset_name": "Asset mobil baru 2"
}
```

Result

```
{
    "data": null,
    "error": {
        "message": "",
        "status_code": 0
    },
    "status": "success"
}
```

## Remove Asset from Family Member

### DELETE /family/member/:family_member_id/asset

Body

```
{
    "asset_id": 7
}
```

Result

```
{
    "data": null,
    "error": {
        "message": "",
        "status_code": 0
    },
    "status": "success"
}
```
