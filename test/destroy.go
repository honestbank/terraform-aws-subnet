package test

import (
	"log"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func destroy(t *testing.T, terraformTempDir string) {
	var terraformDestroyOptions = &terraform.Options{
		TerraformDir: terraformTempDir,
		VarFiles:     []string{"test/terratest.tfvars"},
		Vars:         map[string]interface{}{},
		Reconfigure:  true,
	}
	log.Println("Destroying resources...")
	terraform.Destroy(t, terraformDestroyOptions)
}
