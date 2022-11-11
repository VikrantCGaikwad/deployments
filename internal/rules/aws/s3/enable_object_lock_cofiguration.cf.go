package s3

var cloudFormationEnableObjectLockConfigurationGoodExamples = []string{
	`
Resources:
  GoodExample:
    Properties:
      BucketEncryption:
        ServerSideEncryptionConfiguration:
          - ObjectLockEnabled: "Enabled"
            ObjectLockRule:
               	Days: 10
 				Mode: "COMPLIANCE"
  				Years: 0
    Type: AWS::S3::Bucket
`,
}

var cloudFormationEnableObjectLockConfigurationBadExamples = []string{
	`---
Resources:
  BadExample:
    Properties:
      BucketEncryption:
        ServerSideEncryptionConfiguration:
          
    Type: AWS::S3::Bucket
`,
}

var cloudFormationEnableObjectLockConfigurationLinks = []string{}

var cloudFormationEnableObjectLockConfigurationRemediationMarkdown = ``
