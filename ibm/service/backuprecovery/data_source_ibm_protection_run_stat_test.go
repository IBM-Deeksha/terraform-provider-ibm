// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProtectionRunStatDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionRunStatDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_protection_run_stat.protection_run_stat_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_protection_run_stat.protection_run_stat_instance", "run_id"),
				),
			},
		},
	})
}

func testAccCheckIbmProtectionRunStatDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_protection_run_stat" "protection_run_stat_instance" {
			runId = "runId"
			objects = [ 1 ]
			tenantIds = [ "tenantIds" ]
			includeTenants = true
			includeFinishedTasks = true
			startTimeUsecs = 1
			endTimeUsecs = 1
			maxTasksNum = 1
			excludeObjectDetails = true
			runTaskPath = "runTaskPath"
			objectTaskPaths = [ "objectTaskPaths" ]
		}
	`)
}
