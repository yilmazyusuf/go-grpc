# go-grpc
docker build -t parse_article .
docker run -p 8080:8080 -p 50051:50051 --name parse_article --rm parse_article
lsof -ti:5901 | xargs kill -9
http://localhost:8080/read_article/?url=https://stackoverflow.com/questions/25455233/how-to-solve-too-many-arguments-to-retfdsgurn-issue-in-golang

go run console_client.go https://stackoverflow.com/questions/25455233/how-to-solve-too-many-arguments-to-retfdsgurn-issue-in-golang
