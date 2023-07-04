# Session 03

## Register
curl -X "POST" "http://localhost:8090/api/public/register" \
     -H 'Content-Type: text/plain; charset=utf-8' \
     -d $'{
	"username": "thanh",
	"password": "12345678",
	"full_name": "Thanh Tran",
	"address": "200 Duong 3/2, Ho Chi Minh city"
}'


## Login
curl -X "POST" "http://localhost:8090/api/public/login" \
     -H 'Content-Type: text/plain; charset=utf-8' \
     -d $'{
	"username": "thanh",
	"password": "12345678"
}'

## Self
curl "http://localhost:8090/api/private/self" \
     -H 'Authorization: Bearer ***'

