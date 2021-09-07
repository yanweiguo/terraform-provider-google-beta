// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccOSConfigGuestPolicies_osConfigGuestPoliciesBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckOSConfigGuestPoliciesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOSConfigGuestPolicies_osConfigGuestPoliciesBasicExample(context),
			},
			{
				ResourceName:            "google_os_config_guest_policies.guest_policies",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"guest_policy_id", "project"},
			},
		},
	})
}

func testAccOSConfigGuestPolicies_osConfigGuestPoliciesBasicExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_compute_image" "my_image" {
  provider = google-beta
  family  = "debian-9"
  project = "debian-cloud"
}

resource "google_compute_instance" "foobar" {
  provider = google-beta
  name           = "tf-test-guest-policy-inst%{random_suffix}"
  machine_type   = "e2-medium"
  zone           = "us-central1-a"
  can_ip_forward = false
  tags           = ["foo", "bar"]

  boot_disk {
    initialize_params {
      image = data.google_compute_image.my_image.self_link
    }
  }

  network_interface {
    network = "default"
  }

  metadata = {
    foo = "bar"
  }
}

resource "google_os_config_guest_policies" "guest_policies" {
  provider = google-beta
  guest_policy_id = "tf-test-guest-policy%{random_suffix}"

  assignment {
    instances = [google_compute_instance.foobar.id]
  }

  packages {
    name = "my-package"
    desired_state = "UPDATED"
  }
}
`, context)
}

func TestAccOSConfigGuestPolicies_osConfigGuestPoliciesPackagesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckOSConfigGuestPoliciesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOSConfigGuestPolicies_osConfigGuestPoliciesPackagesExample(context),
			},
			{
				ResourceName:            "google_os_config_guest_policies.guest_policies",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"guest_policy_id", "project"},
			},
		},
	})
}

func testAccOSConfigGuestPolicies_osConfigGuestPoliciesPackagesExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_os_config_guest_policies" "guest_policies" {
  provider = google-beta
  guest_policy_id = "tf-test-guest-policy%{random_suffix}"

  assignment {
    group_labels {
      labels = {
        color = "red",
        env = "test"
      }
    }

    group_labels {
      labels = {
        color = "blue",
        env = "test"
      }
    }
  }

  packages {
    name          = "my-package"
    desired_state = "INSTALLED"
  }

  packages {
    name          = "bad-package-1"
    desired_state = "REMOVED"
  }

  packages {
    name          = "bad-package-2"
    desired_state = "REMOVED"
    manager       = "APT"
  }

  package_repositories {
    apt {
      uri          = "https://packages.cloud.google.com/apt"
      archive_type = "DEB"
      distribution = "cloud-sdk-stretch"
      components   = ["main"]
    }
  }

  package_repositories {
    yum {
      id           = "google-cloud-sdk"
      display_name = "Google Cloud SDK"
      base_url     = "https://packages.cloud.google.com/yum/repos/cloud-sdk-el7-x86_64"
      gpg_keys     = ["https://packages.cloud.google.com/yum/doc/yum-key.gpg", "https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg"]
    }
  }
}
`, context)
}

func TestAccOSConfigGuestPolicies_osConfigGuestPoliciesRecipesExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckOSConfigGuestPoliciesDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOSConfigGuestPolicies_osConfigGuestPoliciesRecipesExample(context),
			},
			{
				ResourceName:            "google_os_config_guest_policies.guest_policies",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"guest_policy_id", "project"},
			},
		},
	})
}

func testAccOSConfigGuestPolicies_osConfigGuestPoliciesRecipesExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_os_config_guest_policies" "guest_policies" {
  provider = google-beta
  guest_policy_id = "tf-test-guest-policy%{random_suffix}"

  assignment {
    zones = ["us-east1-b", "us-east1-d"]
  }

  recipes {
    name          = "tf-test-guest-policy%{random_suffix}-recipe"
    desired_state = "INSTALLED"

    artifacts {
      id = "tf-test-guest-policy%{random_suffix}-artifact-id"

      gcs {
        bucket     = "my-bucket"
        object     = "executable.msi"
        generation = 1546030865175603
      }
    }

    install_steps {
      msi_installation {
        artifact_id = "tf-test-guest-policy%{random_suffix}-artifact-id"
      }
    }
  }
}
`, context)
}

func testAccCheckOSConfigGuestPoliciesDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_os_config_guest_policies" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{OSConfigBasePath}}projects/{{project}}/guestPolicies/{{guest_policy_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("OSConfigGuestPolicies still exists at %s", url)
			}
		}

		return nil
	}
}
