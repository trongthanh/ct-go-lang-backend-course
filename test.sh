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
response=$(curl -i -X "POST" "http://localhost:8090/api/public/user/login" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
	"email": "thanh@chotot.vn",
	"password": "12345678"
}')

# Extract the token from the response using grep and awk
TOKEN=$(echo "$response" | grep -oE '{"Token":"[^"]+' | awk -F'"' '{print $4}')

# Print the token (optional, you can use it as needed in your script)
echo "Recieved token: $TOKEN"

echo "--------------------------"

echo "Get own user profile /me"
sleep 1
## Self
curl -i "http://localhost:8090/api/private/user/me" \
     -H "Authorization: Bearer $TOKEN"

echo "--------------------------"
echo "Create post (with image upload) /post/create"
sleep 1
## Upload image
curl -i -X POST "http://localhost:8090/api/private/post/create" \
     -H "Authorization: Bearer $TOKEN" \
	 -F "caption=Hello world" \
	 -F 'file=@/Users/trantrongthanh/Pictures/tot.png'

echo "--------------------------"
echo "List posts /post/all"
sleep 1
curl -i "http://localhost:8090/api/private/post/all" \
     -H "Authorization: Bearer $TOKEN" \
     -H 'Content-Type: application/json; charset=utf-8'


echo "--------------------------"
echo "Like a post"
sleep 1
curl -i -X POST "http://localhost:8090/api/private/post/like/65227b0f0dff15d1c45490ad" \
     -H "Authorization: Bearer $TOKEN" \
     -H 'Content-Type: application/json; charset=utf-8'


echo ALL TESTS COMPLETED


