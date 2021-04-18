# Spooky
> ![golang-logo](spooky.jpg)

# How to run?
```bash
./main
```

# Input format

## Projects
```json
[
  {
    "id": 1,
    "url": "https://webhook.site/053b3d48-c724-4c39-8d4b-5d31aba0cb2c",
    "students": [
      { "code": "200000000", "name": "John Doe" }
    ],
    "endpoints": [
      { "key": "create_user", "path": "/" }
    ]
  }
]
```

## Test cases
```json
[
  {
    "name": "Create user",
    "endpoint_key": "create_user",
    "method": "POST",
    "status_code": 200,
    "auth": false,
    "params": [
      { "key": "name", "value": "Jhon", "aliases": ["full_name", "first_name", "last_name"]},
      { "key": "username", "value": "john_doe"},
      { "key": "email", "generate": true },
      { "key": "birthdate", "value": "956102016" },
      { "key": "password", "generate": true, "aliases": ["confirm_password"]}
    ],
    "response": [
      { "key": "id", "required": true },
      { "key": "email", "required": true, "match": true },
      { "key": "username", "required": true, "match": true },
      { "key": "auth_token", "required": true, "aliases": ["jwt", "token"] }
    ]
  }
]
```