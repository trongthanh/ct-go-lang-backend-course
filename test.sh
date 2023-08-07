# Integration test for the API

## Register
echo "Register new user /register"
sleep 1
curl -i -X "POST" "http://localhost:8090/api/public/register" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
	"username": "thanh",
	"password": "12345678",
	"full_name": "Thanh Tran",
	"address": "200 Duong 3/2, Ho Chi Minh city"
}'

echo "--------------------------"

## Login
echo "Login /login"
sleep 1
#Make the POST request and store the response in a variable
response=$(curl -i -X "POST" "http://localhost:8090/api/public/login" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
	"username": "thanh",
	"password": "12345678"
}')

# Extract the token from the response using grep and awk
TOKEN=$(echo "$response" | grep -oE '{"Token":"[^"]+' | awk -F'"' '{print $4}')

# Print the token (optional, you can use it as needed in your script)
echo "Recieved token: $TOKEN"

echo "--------------------------"

echo "Get user info /self"
sleep 1
## Self
curl -i "http://localhost:8090/api/private/self" \
     -H "Authorization: Bearer $TOKEN"

echo "--------------------------"
echo "Upload image /upload-imag"
sleep 1
## Upload image
curl -i -X POST "http://localhost:8090/api/private/upload-image" \
     -H "Authorization: Bearer $TOKEN" \
	 -F 'file=@/Users/trantrongthanh/Pictures/tot.png'

echo ALL TESTS COMPLETED
