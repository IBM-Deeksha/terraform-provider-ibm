// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBaasConnectionRegistrationTokenBasic(t *testing.T) {
	connectionID := fmt.Sprintf("tf_connection_id_%d", acctest.RandIntRange(10, 100))
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasConnectionRegistrationTokenConfigBasic(connectionID, xIbmTenantID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_baas_connection_registration_token.baas_connection_registration_token_instance", "connection_id", connectionID),
					resource.TestCheckResourceAttr("ibm_baas_connection_registration_token.baas_connection_registration_token_instance", "x_ibm_tenant_id", xIbmTenantID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_connection_registration_token.baas_connection_registration_token",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasConnectionRegistrationTokenConfigBasic(connectionID string, xIbmTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_connection_registration_token" "baas_connection_registration_token_instance" {
			connection_id = "%s"
			x_ibm_tenant_id = "%s"
		}
	`, connectionID, xIbmTenantID)
}
