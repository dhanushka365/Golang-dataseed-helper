package sdata


import (
	"github.com/DygiSecure/triton-microservice-backend/services/policy/internal/pkg/dao"
	"github.com/DygiSecure/triton-microservice-backend/services/policy/pkg/remediation"
	"github.com/DygiSecure/triton-microservice-backend/services/policy/pkg/severity"
)

func GetIntPointer(value int) *int {
	return &value
}

var RiskCategoryData = []*dao.RiskCategory{
{
	Name:"Identity Management and Weak Authentication",
},
}

var PolicyRuleData = []*dao.PolicyRule{
{
	PolicyFrameworkID:1,
	RemediationTypeID:2,
	GuidedRemediationTemplateID:1,
	PolicyRuleCategoryID:1,
	PolicyNo:"1.1",
	Definition:"Avoid the use of the \"root\" account",
	RiskCategoryID:1,
	Severity:severity.Low.String(),
	TritonConfigRuleRef:"value not defined",
	ScheduleType:"Periodic",
	IsEnabled:true,
},
{
	PolicyFrameworkID:1,
	RemediationTypeID:2,
	GuidedRemediationTemplateID:1,
	PolicyRuleCategoryID:1,
	PolicyNo:"1.2",
	Definition:"Ensure multi-factor authentication (MFA) is enabled for all IAM users that have a console password",
	RiskCategoryID:1,
	Severity:severity.Medium.String(),
	TritonConfigRuleRef:"mfa-enabled-for-iam-console-access",
	ScheduleType:"Periodic",
	IsEnabled:true,
},
}

var PolicyRuleCategoryData = []*dao.PolicyRuleCategory{
{
	Name:"Identity and Access Management",
	Description:"Apply fine-grained permissions to AWS services and resources",
},
{
	Name:"ACM Certificate",
	Description:"makes it easy to provision, manage, and deploy SSL/TLS certificates on AWS managed resource",
},
{
	Name:"API Gateway",
	Description:"an API management tool between a client and a collection of backend services",
},
}

var PolicyFrameworkData = []*dao.PolicyFramework{
{
	Code:"1.2.0",
	Name:"AWS CIS",
	Description:"The CIS Security Benchmarks program provides well-defined, unbiased, consensus-based industry best practices to help organizations assess and improve their security",
	IsEnabled:true,
	IsCustomized:false,
	Logo:"assets/img/logo/cis-logo.png",
},
}

var PolicyRuleAvailableCloudData = []*dao.PolicyRuleAvailableCloud{
{
	PolicyRuleID:1,
	CloudAssetID:4,
},
{
	PolicyRuleID:2,
	CloudAssetID:4,
},
{
	PolicyRuleID:3,
	CloudAssetID:4,
},
{
	PolicyRuleID:4,
	CloudAssetID:4,
},
{
	PolicyRuleID:5,
	CloudAssetID:4,
},
{
	PolicyRuleID:6,
	CloudAssetID:4,
},
{
	PolicyRuleID:7,
	CloudAssetID:4,
},
{
	PolicyRuleID:8,
	CloudAssetID:4,
},
{
	PolicyRuleID:9,
	CloudAssetID:4,
},
}

var AvailableCloudData = []*dao.AvailableCloud{
{
	CloudName:"AWS",
	ImageURL:"assets/img/logo/aws.png",
},
}

var GuidedRemediationTemplateData = []*dao.GuidedRemediationTemplate{
{
	Step:"<p>To configure MFA for a user Open the IAM console at <a href=\"https://console.aws.amazon.com/iam/\">https://console.aws.amazon.com/iam/</a>.Choose Users.</p><p>Choose the User name of the user to configure MFA for. Choose Security credentials and then choose Manage next to Assigned MFA device.</p><p>Follow the Manage MFA Device wizard to assign the type of device appropriate for your environment.</p><p>To learn how to delegate MFA setup to users, see How to Delegate Management of Multi-Factor Authentication to AWS IAM Users on the AWS Security Blog.</p>",
},
{
	Step:"<p>To modify the password policy Open the IAM console at <a href=\"https://console.aws.amazon.com/iam/.\">https://console.aws.amazon.com/iam/. </a></p><p>Choose Account settings. Select Requires at least one uppercase letter and then choose Apply password policy.</p>",
},
{
	Step:"<p>Remediation</p><p>To modify the password policy Open the IAM console at <a href=\"https://console.aws.amazon.com/iam/.\">https://console.aws.amazon.com/iam/. </a></p><p>Choose Account settings. Select Require at least one non-alphanumeric character and then choose Apply password policy.</p>",
},
}

var CloudAssetCategoryData = []*dao.CloudAssetCategory{
{
	Name:"Compute",
},
{
	Name:"Storages",
},
}

var CloudAssetData = []*dao.CloudAsset{
{
	AvailableCloudID:1,
	CloudAssetCategoryID:1,
	Name:"EC2",
	Description:"offers the broadest and deepest compute platform, with over 500 instances and choice of the latest processor, storage, networking, operating system, and purchase model to help best match the needs of workload",
},
{
	AvailableCloudID:1,
	CloudAssetCategoryID:1,
	Name:"Images (AMI)",
	Description:"supported and maintained image provided by AWS that provides the information required to launch an instance",
},
}

var RemediationTypeData = []*dao.RemediationType{
	{
	Name: remediation.Automated.String(),
	},
	{
	Name: remediation.Guided.String(),
	},
}
