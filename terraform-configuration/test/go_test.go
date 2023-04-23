package test

import (
	"fmt"
	"testing"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEKSTerraformModule(t *testing.T) {
	t.Parallel()

	// Generate a random name to avoid naming conflicts
	expectedName := fmt.Sprintf("education-eks-%s", random.UniqueId())

	// Specify the path to the Terraform code
	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform-configuration",
		Vars: map[string]interface{}{
			"cluster_name": expectedName,
			"region":       "us-east-1",
		},
	}

	// Clean up resources after the test is complete
	defer terraform.Destroy(t, terraformOptions)

	// Create the EKS cluster
	terraform.InitAndApply(t, terraformOptions)

	// Use terraform output to get the value of the cluster name
	actualName := terraform.Output(t, terraformOptions, "cluster_name")

	// Assert that the expected and actual names match
	assert.Equal(t, expectedName, actualName)
}
