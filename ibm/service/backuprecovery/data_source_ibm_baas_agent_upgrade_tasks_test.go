// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBaasAgentUpgradeTasksDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasAgentUpgradeTasksDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_agent_upgrade_tasks.baas_agent_upgrade_tasks_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasAgentUpgradeTasksDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_baas_agent_upgrade_tasks" "baas_agent_upgrade_tasks_instance" {
			X-IBM-Tenant-Id = "X-IBM-Tenant-Id"
			ids = [ 1 ]
		}
	`)
}
