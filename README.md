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

With the server running and the example JSON saved as `plan.json` you can send it
using `curl`:

```bash
curl -X POST \ 
  -H "Content-Type: application/json" \
  --data @plan.json \
  http://localhost:3000/api/v1/plan/analyse
```

The response includes the number of resources, the total monthly cost estimate
and a list of each resource with its individual estimated cost.
