swagger generate server -f swagger.yml
go build .\cmd\upbound-app-metadata-server
.\upbound-app-metadata-server --port=8082

curl -X GET "http://localhost:52193/app?title=Valid%20app1"
curl -X POST "http://localhost:52193/app" -H "accept: application/json" -H "Content-Type: application/json" -d @input.txt
