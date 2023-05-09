## make test

### register
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"zhuj", "password":"123456"}' \
    http://localhost:8000/auth/register

### login
curl -i -H "Content-Type: application/json" \
    -X POST \
    -d '{"username":"zhuj", "password":"123456"}' \
    http://localhost:8000/auth/login

### add entry
curl -d '{"Content":"test content 444"}' \
    -H "Authorization: Bearer <<JWT returned by login>>" \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8000/entry

### list entries
curl -d '{"Content":"test content"}' \
    -H "Authorization: Bearer <JWT return by login>>" \
    -H "Content-Type: application/json" \
    -X GET http://localhost:8000/entry
