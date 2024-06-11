# instagram-clone

## Run the project
- Nếu chưa run docker or chưa start mysql thì có thể chạy
```shell
docker-compose up
```

- Sau đó chạy project bằng lệnh
```shell
CONNECTION_STRING=capybara:my_secret@tcp(localhost:3306)/users?parseTime=true PORT=:8001 go run .
```