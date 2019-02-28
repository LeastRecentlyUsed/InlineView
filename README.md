# LandRegPriceInspector

An automated process which retrieves the latest UK Land registry prices of sold residential homes and makes pre-structured data available on a simple web page.  Utilises AWS cloud services.

A Data pipeline will extract the native CSV and cleanse the contents into a semi-structured NOSQL database and well as saving the raw source file.

# Project Build Activities 1 (code environment)
 
Install the AWS Go SDK ("/..." install dependancies)
  go get -u github.com/aws/aws-sdk-go/...

For AWS SDK docs see https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html
For Go docs see https://godoc.org/github.com/aws/aws-sdk-go/aws/session

Once installed set the AWS region in the aws config file
~/.aws/config (posix)
c:\Users\username\.aws\config (windows)

[default]
region = eu-west-1 

a list of region names is at https://docs.aws.amazon.com/general/latest/gr/rande.html

set the aws IAM access keys in the (shared) credentials file
~/.aws/credentials (posix)
c:\Users\username\.aws\credentials (windows)
see the link at https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html

[default]
aws_access_key_id = 
aws_secret_access_key = 

Use the AWS package (core SDK)
Use the AWS service package (cognito, dynamoDB, S3, etc)


# AWS Setup 1 (IAM and policies)

using the IAM secion of the console create a resource group and add a user within that group.

For the new user go to the 'security credentials' tab in the user summary and copy the access key and the secret access key for use in the 'credentials' config file.

For the resource group, in the 'permissions' tab of the group summary attach one of the standard policies, for example the S3 full access policy.

At this point it should be possible to query your AWS account for any S3 storage buckets.




