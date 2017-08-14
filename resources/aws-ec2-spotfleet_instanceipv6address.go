package resources

// AWS::EC2::SpotFleet.InstanceIpv6Address AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-instanceipv6address.html
type AWSEC2SpotFleetInstanceIpv6Address struct {

	// Ipv6Address AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-instanceipv6address.html#cfn-ec2-spotfleet-instanceipv6address-ipv6address
	Ipv6Address string `json:"Ipv6Address"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleetInstanceIpv6Address) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.InstanceIpv6Address"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleetInstanceIpv6Address) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
