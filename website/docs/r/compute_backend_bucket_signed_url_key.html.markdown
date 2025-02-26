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
page_title: "Google: google_compute_backend_bucket_signed_url_key"
description: |-
  A key for signing Cloud CDN signed URLs for BackendBuckets.
---

# google\_compute\_backend\_bucket\_signed\_url\_key

A key for signing Cloud CDN signed URLs for BackendBuckets.


To get more information about BackendBucketSignedUrlKey, see:

* [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendBuckets)
* How-to Guides
    * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)

~> **Warning:** All arguments including `key_value` will be stored in the raw
state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/language/state/sensitive-data).

## Example Usage - Backend Bucket Signed Url Key


```hcl
resource "random_id" "url_signature" {
  byte_length = 16
}

resource "google_compute_backend_bucket_signed_url_key" "backend_key" {
  name           = "test-key"
  key_value      = random_id.url_signature.b64_url
  backend_bucket = google_compute_backend_bucket.test_backend.name
}

resource "google_compute_backend_bucket" "test_backend" {
  name        = "test-signed-backend-bucket"
  description = "Contains beautiful images"
  bucket_name = google_storage_bucket.bucket.name
  enable_cdn  = true
}

resource "google_storage_bucket" "bucket" {
  name     = "test-storage-bucket"
  location = "EU"
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  Name of the signed URL key.

* `key_value` -
  (Required)
  128-bit key value used for signing the URL. The key value must be a
  valid RFC 4648 Section 5 base64url encoded string.
  **Note**: This property is sensitive and will not be displayed in the plan.

* `backend_bucket` -
  (Required)
  The backend bucket this signed URL key belongs.


- - -


* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/global/backendBuckets/{{backend_bucket}}`


## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import

This resource does not support import.

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
