---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "Compute Engine"
page_title: "Google: google_compute_global_network_endpoint"
description: |-
  A Global Network endpoint represents a IP address and port combination that exists outside of GCP.
---

# google\_compute\_global\_network\_endpoint

A Global Network endpoint represents a IP address and port combination that exists outside of GCP.
**NOTE**: Global network endpoints cannot be created outside of a
global network endpoint group.


To get more information about GlobalNetworkEndpoint, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/)

## Example Usage - Global Network Endpoint


```hcl
resource "google_compute_global_network_endpoint" "default-endpoint" {
  global_network_endpoint_group = google_compute_global_network_endpoint_group.neg.name
  fqdn       = "www.example.com"
  port       = 90
}

resource "google_compute_global_network_endpoint_group" "neg" {
  name                  = "my-lb-neg"
  default_port          = "90"
  network_endpoint_type = "INTERNET_FQDN_PORT"
}
```

## Argument Reference

The following arguments are supported:


* `port` -
  (Required)
  Port number of the external endpoint.

* `global_network_endpoint_group` -
  (Required)
  The global network endpoint group this endpoint is part of.


- - -


* `ip_address` -
  (Optional)
  IPv4 address external endpoint.

* `fqdn` -
  (Optional)
  Fully qualified domain name of network endpoint.
  This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{project}}/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}`


## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


GlobalNetworkEndpoint can be imported using any of these accepted formats:

```
$ terraform import google_compute_global_network_endpoint.default projects/{{project}}/global/networkEndpointGroups/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
$ terraform import google_compute_global_network_endpoint.default {{project}}/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
$ terraform import google_compute_global_network_endpoint.default {{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
