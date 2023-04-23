package test

import (
    "fmt"
    "os"
    "path/filepath"
    "testing"

    "github.com/gruntwork-io/terratest/modules/aws"
    "github.com/gruntwork-io/terratest/modules/random"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraform(t *testing.T) {
    terraformOptions := &terraform.Options{
        TerraformDir: "../",
        Vars: map[string]interface{}{
            "region":           "us-east-1",
            "suffix":           random.UniqueId(),
            "instance_type":    "t3.small",
        },
    }

    // Clean up everything at the end of the test
    defer terraform.Destroy(t, terraformOptions)

    // Deploy the infrastructure
    terraform.InitAndApply(t, terraformOptions)

    // Run Terraform output to get the outputs
    clusterEndpoint := terraform.Output(t, terraformOptions, "cluster_endpoint")
    clusterSecurityGroup := terraform.Output(t, terraformOptions, "cluster_security_group_id")
    region := terraform.Output(t, terraformOptions, "region")

    // Verify that the outputs are valid
    assert.NotEmpty(t, clusterEndpoint)
    assert.NotEmpty(t, clusterSecurityGroup)
    assert.Equal(t, "us-east-1", region)

    // Use the AWS SDK to make sure the cluster endpoint is reachable
    awsRegion := aws.GetRegion(t, region)
    awsSession, err := aws.NewSessionE(t, &aws.Options{
        Profile: awsProfile,
        Region:  awsRegion,
    })
    if err != nil {
        t.Fatalf("Error creating AWS session: %v", err)
    }
    pingResult := aws.EksPing(t, awsSession, clusterEndpoint)
    assert.True(t, pingResult)
}
