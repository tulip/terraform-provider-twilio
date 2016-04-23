# Twilio Provider

The provider configuration block accepts the following arguments:

* ``account_sid`` - (Required) Your SID (application ID) for the the Twilio API. May alternatively be set via the
  ``TWILIO_SID`` environment variable.

* ``auth_token`` - (Required) The API auth token to use when making requests. May alternatively
  be set via the ``TWILIO_AUTH_TOKEN`` environment variable.

## Phone number resource

Does a search for available phone number via the Local API described here: https://www.twilio.com/docs/api/rest/available-phone-numbers and then purchases it

optional ["Advanced Filters"](https://www.twilio.com/docs/api/rest/available-phone-numbers#local-get-advanced-filters) are nested under `location` block. (and `In`s are dropped)

Purchase, update, and delete API here:
https://www.twilio.com/docs/api/rest/incoming-phone-numbers

Arguments mostly follow those documented there with snake case instead of CamelCase (though with `name` instead of `FriendlyName`).

TL;DR  read the [schema](https://github.com/donaldguy/terraform-provider-twilio/blob/master/resource_phonenumber_schema.go)


### Example Usage

```terraform
provider "twilio" {
    account_sid = "ASdiid03kjj40304mmd03043893434id"
    auth_token = "abcd1234"
}

resource "twilio_phonenumber" "mexico" {
  iso_country_code = "MX"
}

resource "twilio_phonenumber" "boston" {
    name = "boston reserved"

    location {
      near_lat_long {
        longitude = 42.3755210
        latitude = -71.0932520
      }
    }
}

resource "twilio_phonenumber" "virginia" {
    name = "Virginia"

    location {
      region = "VA"
    }

    sms_method = "POST"
    sms_url = "https://example.com/smsEndpoint"
}


```
