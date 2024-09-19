// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasProtectionGroupBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	environment := "kPhysical"
	includedPath := "/"
	includedPathUpdate := "/data1/"
	protectionType := "kFile"
	objectId := 72

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasProtectionGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupConfigBasic(name, environment, includedPath, protectionType, objectId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasProtectionGroupExists("ibm_baas_protection_group.baas_protection_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupConfigBasic(name, environment, includedPathUpdate, protectionType, objectId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "name", name),
				),
			},
		},
	})
}

func testAccCheckIbmBaasProtectionGroupConfigBasic(name, environment, includedPath, protectionType string, objectId int) string {
	return fmt.Sprintf(`
			resource "ibm_baas_protection_policy" "baas_protection_policy_instance" {
				x_ibm_tenant_id = "%[1]s"
				name = "tf-name-policy-test-1"
				backup_policy {
						regular {
							incremental{
								schedule{
										day_schedule {
											frequency = 1
										}
										unit = "Days"
									}
							}
							retention {
								duration = 1
								unit = "Weeks"
							}
							primary_backup_target {
								use_default_backup_target = true
							}
						}
				}
				retry_options {
				retries = 3
				retry_interval_mins = 5
				}
			}

		resource "ibm_baas_protection_group" "baas_protection_group_instance" {
			x_ibm_tenant_id = "%[1]s"
			policy_id = ibm_baas_protection_policy.baas_protection_policy_instance.id
			name = "%[2]s"
			environment = "%[3]s"
			physical_params {
				protection_type = "%[5]s"
				file_protection_type_params {
				objects {
					id = %d
					file_paths{
						included_path = "%[4]s"
					}
				  }
				}
			}
		}
	`, tenantId, name, environment, includedPath, protectionType, objectId)
}

func testAccCheckIbmBaasProtectionGroupExists(n string, obj backuprecoveryv1.ProtectionGroupResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionGroupByIdOptions := &backuprecoveryv1.GetProtectionGroupByIdOptions{}

		getProtectionGroupByIdOptions.SetID(rs.Primary.ID)
		getProtectionGroupByIdOptions.SetXIBMTenantID(tenantId)

		protectionGroupResponse, _, err := backupRecoveryClient.GetProtectionGroupByID(getProtectionGroupByIdOptions)
		if err != nil {
			return err
		}

		obj = *protectionGroupResponse
		return nil
	}
}

func testAccCheckIbmBaasProtectionGroupDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_protection_group" {
			continue
		}

		getProtectionGroupByIdOptions := &backuprecoveryv1.GetProtectionGroupByIdOptions{}

		getProtectionGroupByIdOptions.SetID(rs.Primary.ID)
		getProtectionGroupByIdOptions.SetXIBMTenantID(tenantId)

		// Try to find the key
		groupResponse, response, err := backupRecoveryClient.GetProtectionGroupByID(getProtectionGroupByIdOptions)

		if err == nil {
			if strings.Contains(*groupResponse.Name, fmt.Sprintf("_DELETED_%s", rs.Primary.Attributes["name"])) {
				return nil
			}
			return fmt.Errorf("baas_protection_group still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_protection_group (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
