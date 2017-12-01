package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsAutoscalingGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAwsAutoscalingGroupRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"launch_configuration": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"desired_capacity": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"min_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"max_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"default_cooldown": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"health_check_grace_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"health_check_type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"availability_zones": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"placement_group": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"load_balancers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"vpc_zone_identifier": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"termination_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"enabled_metrics": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"suspended_processes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"protect_from_scale_in": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},

			// "tags": tagsSchemaComputed(),  TODO
		},
	}
}
