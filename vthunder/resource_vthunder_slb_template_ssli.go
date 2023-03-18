package vthunder

//vThunder resource TemplateSSLI

import (
	log "github.com/sourcegraph-ce/logrus"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTemplateSSLI() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateSSLICreate,
		Update: resourceTemplateSSLIUpdate,
		Read:   resourceTemplateSSLIRead,
		Delete: resourceTemplateSSLIDelete,
		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTemplateSSLICreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateSSLI (Inside resourceTemplateSSLICreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateSSLI(d)
		logger.Println("[INFO] received formatted data from method data to TemplateSSLI --")
		d.SetId(name)
		go_vthunder.PostTemplateSSLI(client.Token, data, client.Host)

		return resourceTemplateSSLIRead(d, meta)

	}
	return nil
}

func resourceTemplateSSLIRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading TemplateSSLI (Inside resourceTemplateSSLIRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetTemplateSSLI(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceTemplateSSLIUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateSSLI   (Inside resourceTemplateSSLIUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateSSLI(d)
		logger.Println("[INFO] received formatted data from method data to TemplateSSLI ")
		d.SetId(name)
		go_vthunder.PutTemplateSSLI(client.Token, name, data, client.Host)

		return resourceTemplateSSLIRead(d, meta)

	}
	return nil
}

func resourceTemplateSSLIDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateSSLIDelete) " + name)
		err := go_vthunder.DeleteTemplateSSLI(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateSSLI(d *schema.ResourceData) go_vthunder.SSLI {
	var vc go_vthunder.SSLI
	var c go_vthunder.SsliInstance
	c.Name = d.Get("name").(string)
	c.Type = d.Get("type").(string)
	c.UserTag = d.Get("user_tag").(string)
	vc.UUID = c
	return vc
}
