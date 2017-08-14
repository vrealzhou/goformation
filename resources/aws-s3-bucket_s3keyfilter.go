package resources

// AWS::S3::Bucket.S3KeyFilter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html
type AWSS3BucketS3KeyFilter struct {

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-notificationconfiguration-config-filter-s3key.html#cfn-s3-bucket-notificationconfiguraiton-config-filter-s3key-rules
	Rules []AWSS3BucketS3KeyFilterFilterRule `json:"Rules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketS3KeyFilter) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.S3KeyFilter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketS3KeyFilter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
