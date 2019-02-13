### This server required swagger and Go
### Please download the latest version of swagger and regenerate swagger files if you plan on making changes to swagger.yml
### Sample swagger generation and build mechanism

swagger generate server -f swagger.yml
go build .\cmd\upbound-app-metadata-server
.\upbound-app-metadata-server --port=8082

Sample usage: (sample input can be found under 'inputSample')
curl -X GET "http://localhost:8082/app?title=Valid%20app1"
curl -X POST "http://localhost:8082/app" -H "accept: application/json" -H "Content-Type: application/json" -d @input.txt
