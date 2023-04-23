package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEksClusterCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform-eks",
		Vars: map[string]interface{}{
			"region": "us-west-2",
		},
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	region := aws.GetRandomStableRegion(t, nil, nil)
	session, err := aws.NewSessionE(t, &aws.Options{
		Region: region,
	})
	assert.NoError(t, err)

	clusterName := terraform.Output(t, terraformOptions, "cluster_name")

	err = aws.EksPing(t, session, clusterName)
	assert.NoError(t, err)

	// Add more assertions here to validate that the EKS cluster was created successfully.
}
