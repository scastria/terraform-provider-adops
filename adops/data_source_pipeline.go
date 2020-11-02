package adops

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/scastria/terraform-provider-adops/adops/client"
	"strconv"
)

func dataSourcePipeline() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePipelineRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePipelineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	id := strconv.Itoa(d.Get("id").(int))
	body, err := c.HttpRequest("pipelines/"+id, "GET", bytes.Buffer{})
	if err != nil {
		return diag.FromErr(err)
	}
	pipeline := &client.Pipeline{}
	err = json.NewDecoder(body).Decode(pipeline)
	if err != nil {
		return diag.FromErr(err)
	}
	var diags diag.Diagnostics
	d.Set("name", pipeline.Name)
	d.SetId(id)
	return diags
}
