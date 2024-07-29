// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmProtectionGroupRunRequestBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupRun
	runType := "kRegular"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProtectionGroupRunRequestDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionGroupRunRequestConfigBasic(runType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProtectionGroupRunRequestExists("ibm_protection_group_run_request.protection_group_run_request_instance", conf),
					resource.TestCheckResourceAttr("ibm_protection_group_run_request.protection_group_run_request_instance", "run_type", runType),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_protection_group_run_request.protection_group_run_request",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmProtectionGroupRunRequestConfigBasic(runType string) string {
	return fmt.Sprintf(`
		resource "ibm_protection_group_run_request" "protection_group_run_request_instance" {
			run_type = "%s"
		}
	`, runType)
}

func testAccCheckIbmProtectionGroupRunRequestExists(n string, obj backuprecoveryv1.ProtectionGroupRun) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionGroupRunOptions := &backuprecoveryv1.GetProtectionGroupRunOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProtectionGroupRunOptions.SetID(parts[0])
		getProtectionGroupRunOptions.SetRunID(parts[1])

		protectionGroupRunRequest, _, err := backupRecoveryClient.GetProtectionGroupRun(getProtectionGroupRunOptions)
		if err != nil {
			return err
		}

		obj = *protectionGroupRunRequest
		return nil
	}
}

func testAccCheckIbmProtectionGroupRunRequestDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_protection_group_run_request" {
			continue
		}

		getProtectionGroupRunOptions := &backuprecoveryv1.GetProtectionGroupRunOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProtectionGroupRunOptions.SetID(parts[0])
		getProtectionGroupRunOptions.SetRunID(parts[1])

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionGroupRun(getProtectionGroupRunOptions)

		if err == nil {
			return fmt.Errorf("protection_group_run_request still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for protection_group_run_request (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
