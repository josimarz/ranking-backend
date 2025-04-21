export AWS_ACCESS_KEY_ID="xyz"
export AWS_SECRET_ACCESS_KEY="123"
aws dynamodb create-table \
    --endpoint-url http://localhost:4566 \
    --region us-east-1 \
    --cli-input-json file://rank-table.json \
    --no-cli-pager