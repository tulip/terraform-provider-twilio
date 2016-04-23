package main

import "github.com/hashicorp/terraform/helper/schema"

func resourcePhonenumber() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "\"FriendlyName\": A human readable description of the new incoming phone number resource, with maximum length 64 characters",
				Required:    true,
			},
			"iso_country_code": &schema.Schema{
				Type:        schema.TypeString,
				Description: "ISO 3166-1 alpha-2 code of country in which to buy a phone number",
				Optional:    true,
				Default:     "US",
				ForceNew:    true,
			},
			"location": &schema.Schema{
				Type:        schema.TypeSet,
				Description: "Information to use in targetting the number to buy",
				Optional:    true,
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"near_number": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Given a phone number, find a geographically close number within Distance miles. Distance defaults to 25 miles.",
							Optional:    true,
						},
						"near_lat_long": &schema.Schema{
							Type:        schema.TypeSet,
							Description: "Given a latitude/longitude pair lat,long find geographically close numbers within Distance miles.",
							Optional:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"latitude": &schema.Schema{
										Type:     schema.TypeFloat,
										Required: true,
									},
									"longitude": &schema.Schema{
										Type:     schema.TypeFloat,
										Required: true,
									},
								},
							},
						},
						"distance": &schema.Schema{
							Type:        schema.TypeInt,
							Description: "Specifies the search radius for a Near- query in miles. If not specified this defaults to 25 miles. Maximum searchable distance is 500 miles.",
							Optional:    true,
							Default:     25,
						},
						"postal_code": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Limit results to a particular postal code. Given a phone number, search within the same postal code as that number.",
							Optional:    true,
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Limit results to a particular region (i.e. State/Province). Given a phone number, search within the same Region as that number.",
							Optional:    true,
						},
						"rate_center": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires InLata to be set as well.",
							Optional:    true,
						},
						"LATA": &schema.Schema{
							Type:        schema.TypeString,
							Description: "Limit results to a specific Local access and transport area (LATA). Given a phone number, search within the same LATA as that number.",
							Optional:    true,
						},
					},
				},
			},
			"phone_number": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The actual phone number purchased",
				Computed:    true,
			},
			"api_version": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Calls to this phone number will start a new TwiML session with this API version. Either 2010-04-01 or 2008-08-01.",
				Optional:    true,
				Computed:    true,
			},
			"voice_caller_id_lookup": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Do a lookup of a caller's name from the CNAM database and post it to your app. Either true or false.",
				Optional:    true,
				Computed:    true,
			},
			"voice_url": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The URL that Twilio should request when somebody dials the phone number",
				Optional:    true,
				Computed:    true,
			},
			"voice_method": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The HTTP method that should be used to request the VoiceUrl. Either GET or POST.",
				Optional:    true,
				Computed:    true,
			},
			"voice_fallback_url": &schema.Schema{
				Type:        schema.TypeString,
				Description: "A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by VoiceUrl.",
				Optional:    true,
				Computed:    true,
			},
			"voice_fallback_method": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The HTTP method that should be used to request the VoiceFallbackUrl. Either GET or POST.",
				Optional:    true,
				Computed:    true,
			},
			"sms_url": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The URL that Twilio should request when somebody sends an SMS to the new phone number.",
				Optional:    true,
				Computed:    true,
			},
			"sms_method": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The HTTP method that should be used to request the SmsUrl. Either GET or POST.",
				Optional:    true,
				Computed:    true,
			},
			"sms_fallback_url": &schema.Schema{
				Type:        schema.TypeString,
				Description: "A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl.",
				Optional:    true,
				Computed:    true,
			},
			"sms_fallback_method": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The HTTP method that should be used to request the SmsFallbackUrl. Either GET or POST.",
				Optional:    true,
				Computed:    true,
			},
			"status_callback": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The URL that Twilio will request to pass status parameters (such as call ended) to your application",
				Optional:    true,
				Computed:    true,
			},
			"status_callback_method": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The HTTP method Twilio will use to make requests to the StatusCallback URL. Either GET or POST",
				Optional:    true,
				Computed:    true,
			},
		},
		Create: phonenumberCreate,
		Read:   phonenumberRead,
		Update: phonenumberUpdate,
		Delete: phonenumberDelete,
	}
}
