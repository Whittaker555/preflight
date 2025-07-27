# PreFlight API

PreFlight is a simple Go service that accepts a JSON representation of a
Terraform plan and returns a summary with a very basic cost estimate.

## Building and running

```bash
# build the server
go build -o preflight ./cmd

# run (listens on port 3000 by default)
./preflight
```

## Terraform plan JSON example

The API expects the body of the request to match the structure below. You can
produce this JSON by running `terraform show -json <plan file> > plan.json` and
then posting `plan.json` to the API.

```json
{
  "resource_changes": [
    {
      "address": "aws_instance.example",
      "type": "aws_instance",
      "name": "example",
      "change": {
        "actions": ["create"]
      }
    },
    {
      "address": "aws_s3_bucket.data",
      "type": "aws_s3_bucket",
      "name": "data",
      "change": {
        "actions": ["create"]
      }
    }
  ]
}
```

## Analysing a plan

With the server running and the example JSON saved as `plan.json` you can send
it to the `/plan/upload` endpoint either as a file upload or directly as JSON.

### Uploading a file

```bash
curl -X POST \
  -F "file=@plan.json" \
  http://localhost:3000/api/v1/plan/upload
```

### Sending raw JSON

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  --data @plan.json \
  http://localhost:3000/api/v1/plan/upload
```

The response includes the number of resources, the total monthly cost estimate
and a list of each resource with its individual estimated cost.

## Uploading analysis results

If you need to store the JSON output from a previous analysis you can POST it
back to the API:

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  --data @result.json \
  http://localhost:3000/api/v1/plan/result/upload
```

The API will write the uploaded JSON to a file under the `results` directory and
return the path to the saved file.

## Postman Collection

The `postman` directory contains a ready-to-use Postman collection and environment for testing the API's health check and plan upload endpoints.

## GitHub Actions Deployment

The repository contains a `Deploy` workflow that packages the API as an AWS Lambda function and applies the Terraform in `infra/`.
To run it you must configure the following secrets in your GitHub repository:

- `AWS_REGION` – AWS region where the Lambda and API Gateway will be created
- `APP_NAME` – name for the Lambda function and API Gateway resources
- `AWS_DEPLOY_ROLE` – ARN of an IAM role that GitHub Actions is allowed to assume

### Creating the deploy role

1. In the AWS console open **IAM → Identity providers** and create an OIDC provider for `https://token.actions.githubusercontent.com` (if you do not already have one).
2. Create a new role for **Web identity** that trusts this provider and restrict the subject to your repository, e.g. `repo:<your org>/<your repo>:*`.
3. Attach policies that permit managing Lambda, API Gateway and IAM resources (using `AdministratorAccess` is easiest for initial tests).
4. Copy the role ARN and store it as the `AWS_DEPLOY_ROLE` secret in GitHub.
