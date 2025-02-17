// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPICloudConnectionsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICloudConnectionsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_cloud_connections.connections", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPICloudConnectionsDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_cloud_connections" "connections" {
			pi_cloud_instance_id = "%s"
		}`, pi_cloud_instance_id)
}
