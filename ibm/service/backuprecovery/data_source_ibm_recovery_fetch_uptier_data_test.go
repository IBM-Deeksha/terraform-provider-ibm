// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmRecoveryFetchUptierDataDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveryFetchUptierDataDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_recovery_fetch_uptier_data.recovery_fetch_uptier_data_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery_fetch_uptier_data.recovery_fetch_uptier_data_instance", "archive_u_id"),
				),
			},
		},
	})
}

func testAccCheckIbmRecoveryFetchUptierDataDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_recovery_fetch_uptier_data" "recovery_fetch_uptier_data_instance" {
			archiveUId = "archiveUId"
		}
	`)
}
