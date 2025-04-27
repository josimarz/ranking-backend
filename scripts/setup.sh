export AWS_ACCESS_KEY_ID="xyz"
export AWS_SECRET_ACCESS_KEY="123"
aws dynamodb create-table \
    --endpoint-url $ENDPOINT_URL \
    --region us-east-1 \
    --table-name $TABLE_NAME \
    --cli-input-json file://rank-table.json \
    --no-cli-pager

aws s3api create-bucket \
    --endpoint-url $ENDPOINT_URL \
    --region us-east-1 \
    --no-cli-pager \
    --acl public-read \
    --bucket ranking