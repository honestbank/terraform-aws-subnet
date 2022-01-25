package test

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	test_structure "github.com/gruntwork-io/terratest/modules/files"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

const localBackend = `
terraform {
	backend "local" {}
}
`

func setupTest() (string, error) {
	terraformTempDir, errCopying := test_structure.CopyTerragruntFolderToTemp("../", "terratest-")
	if errCopying != nil {
		return "", errCopying
	}

	backendFilePath := fmt.Sprintf("%s/%s", terraformTempDir, "backend.tf")
	errRemoving := os.Remove(backendFilePath)
	if errRemoving != nil {
		return "", errRemoving
	}

	errWritingFile := os.WriteFile(backendFilePath, []byte(localBackend), os.ModeAppend)
	if errWritingFile != nil {
		return "", errWritingFile
	}
	os.Chmod(backendFilePath, fs.FileMode(0777))
	return terraformTempDir, nil
}

const tfWorkspaceEnvVarName = "TF_WORKSPACE"
const targetWorkspace = "test"

func TestTerraformCodeInfrastructureInitialCredentials(t *testing.T) {
	Region := "ap-southeast-1"
	terraformTempDir, errSettingUpTest := setupTest()
	if errSettingUpTest != nil {
		t.Fatalf("Error setting up test :%v", errSettingUpTest)
	}
	defer os.RemoveAll(terraformTempDir)
	log.Printf("Temp folder: %s", terraformTempDir)
	terraformInitOptions := &terraform.Options{
		TerraformDir: terraformTempDir,
		VarFiles:     []string{"test/terratest.tfvars"},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": "ap-southeast-1",
			//"TF_LOG":             "TRACE",
		},
		Reconfigure: true,
	}

	defer destroy(t, terraformTempDir)
	terraform.Init(t, terraformInitOptions)
	terraform.WorkspaceSelectOrNew(t, terraformInitOptions, targetWorkspace)
	terraformValidateOptions := &terraform.Options{
		TerraformDir: terraformTempDir,
		EnvVars: map[string]string{
			tfWorkspaceEnvVarName: targetWorkspace,
		},
	}
	terraform.Validate(t, terraformValidateOptions)
	plan, errApplyingIdempotent := terraform.ApplyAndIdempotentE(t, terraformInitOptions)
	if errApplyingIdempotent != nil {
		t.Logf("Error applying plan: %v", errApplyingIdempotent)
		t.Fail()
	} else {
		t.Log(fmt.Sprintf("Plan worked: %s", plan))
	}
	vpc_id := terraform.Output(t, terraformValidateOptions, "vpc_id")
	t.Run("Test Subnet exists", func(t *testing.T) {
		a := assert.New(t)
		//aws.GetVpcById(os.Getenv())
		var vars map[string]interface{}
		terraform.GetAllVariablesFromVarFile(t, "terratest.tfvars", &vars)
		aws.CreateAwsSessionWithCreds(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), Region)
		subnets := aws.GetSubnetsForVpc(t, vpc_id, Region)
		exists := false
		for _, subnet := range subnets {
			if subnet.Tags["Name"] == vars["subnet_name"].(string) {
				exists = true
			}
		}
		a.Equal(true, exists)

	})
}
