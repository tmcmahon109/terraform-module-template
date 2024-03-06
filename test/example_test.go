package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformResources(t *testing.T) {
	t.Parallel()

	// website::tag::2:: Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// website::tag::1:: The path to where our Terraform code is located
		TerraformDir: "../",
	})

	// website::tag::6:: At the end of the test, run `terraform destroy` to clean up any resources that were created.
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::3:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// website::tag::4:: Run `terraform output` to get the desired outputs
	sgName := terraform.Output(t, terraformOptions, "sg_name")
	assert.Equal(t, "variable_test", sgName)

}
