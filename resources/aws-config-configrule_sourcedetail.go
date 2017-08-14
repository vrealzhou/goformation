package resources

// AWS::Config::ConfigRule.SourceDetail AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html
type AWSConfigConfigRuleSourceDetail struct {

	// EventSource AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html#cfn-config-configrule-source-sourcedetail-eventsource
	EventSource string `json:"EventSource"`

	// MaximumExecutionFrequency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html#cfn-config-configrule-sourcedetail-maximumexecutionfrequency
	MaximumExecutionFrequency string `json:"MaximumExecutionFrequency"`

	// MessageType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source-sourcedetails.html#cfn-config-configrule-source-sourcedetail-messagetype
	MessageType string `json:"MessageType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigRuleSourceDetail) AWSCloudFormationType() string {
	return "AWS::Config::ConfigRule.SourceDetail"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigConfigRuleSourceDetail) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
