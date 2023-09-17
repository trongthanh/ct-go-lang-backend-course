# Integration test for the API

## Register
echo "Register new user /register"
sleep 1
curl -i -X "POST" "http://localhost:8090/api/public/user/signup" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
	"email": "thanh@chotot.vn",
	"password": "12345678"
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

echo "--------------------------"

echo "Change password"
sleep 1
## Change password
curl -i "http://localhost:8090/api/private/change-password" \
     -H "Authorization: Bearer $TOKEN" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
	"current_password": "12345678",
	"new_password": "12345679",
	"repeat_password": "12345679"
}'


echo ALL TESTS COMPLETED


