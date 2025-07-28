package cost

var ResourceCosts = map[string]float64{
	"aws_instance":    25.00,  // small EC2 instance
	"aws_s3_bucket":   5.00,   // storage baseline
	"aws_db_instance": 100.00, // RDS baseline
}
