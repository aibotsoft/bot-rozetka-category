#!/usr/bin/env bash
mkdir "api"


#curl -X GET --location "https://api.getpostman.com/collections/$POSTMAN_COLLECTION_ID" \
#    -H "x-api-key: $POSTMAN_API_KEY" > ./openapi/postman.json
#
#if [ $? -ne 0 ]; then
#  echo "curl fail"
#  exit 1
#fi
#jq '.collection' ./openapi/postman.json | postman2openapi > ./openapi/openapi.yaml

java -jar /home/ac/bin/openapi-generator-cli-5.3.0.jar generate -g go -i openapi/openapi-manual.yaml -o api \
  --type-mappings integer=int64 --type-mappings number=float64 --additional-properties=hideGenerationTimestamp=true,packageName=api