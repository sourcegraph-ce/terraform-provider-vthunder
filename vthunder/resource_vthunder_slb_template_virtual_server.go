package vthunder

//vThunder resource SlbTemplateVirtualServer

import (
	log "github.com/sourcegraph-ce/logrus"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateVirtualServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateVirtualServerCreate,
		Update: resourceSlbTemplateVirtualServerUpdate,
		Read:   resourceSlbTemplateVirtualServerRead,
		Delete: resourceSlbTemplateVirtualServerDelete,
		Schema: map[string]*schema.Schema{
			"conn_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"tcp_stack_tfo_backoff_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmp_lockup_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmpv6_lockup_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"tcp_stack_tfo_active_conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"icmpv6_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"tcp_stack_tfo_cookie_time_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_limit_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"subnet_gratuitous_arp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rate_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"icmp_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmpv6_lockup": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmp_lockup": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateVirtualServerCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateVirtualServer (Inside resourceSlbTemplateVirtualServerCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateVirtualServer(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateVirtualServer --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateVirtualServer(client.Token, data, client.Host)

		return resourceSlbTemplateVirtualServerRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateVirtualServerRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateVirtualServer (Inside resourceSlbTemplateVirtualServerRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateVirtualServer(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateVirtualServerUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateVirtualServer   (Inside resourceSlbTemplateVirtualServerUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateVirtualServer(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateVirtualServer ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateVirtualServer(client.Token, name, data, client.Host)

		return resourceSlbTemplateVirtualServerRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateVirtualServerDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateVirtualServerDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateVirtualServer(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateVirtualServer(d *schema.ResourceData) go_vthunder.VirtualServer {
	var vc go_vthunder.VirtualServer
	var c go_vthunder.VirtualServerInstance

	c.ConnLimit = d.Get("conn_limit").(int)
	c.ConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	c.Name = d.Get("name").(string)
	c.IcmpLockupPeriod = d.Get("icmp_lockup_period").(int)
	c.ConnLimitReset = d.Get("conn_limit_reset").(int)
	c.RateInterval = d.Get("rate_interval").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.Icmpv6RateLimit = d.Get("icmpv6_rate_limit").(int)
	c.SubnetGratuitousArp = d.Get("subnet_gratuitous_arp").(int)
	c.Icmpv6Lockup = d.Get("icmpv6_lockup").(int)
	c.ConnRateLimitReset = d.Get("conn_rate_limit_reset").(int)
	c.TCPStackTfoBackoffTime = d.Get("tcp_stack_tfo_backoff_time").(int)
	c.TCPStackTfoCookieTimeLimit = d.Get("tcp_stack_tfo_cookie_time_limit").(int)
	c.ConnLimitNoLogging = d.Get("conn_limit_no_logging").(int)
	c.Icmpv6LockupPeriod = d.Get("icmpv6_lockup_period").(int)
	c.ConnRateLimit = d.Get("conn_rate_limit").(int)
	c.TCPStackTfoActiveConnLimit = d.Get("tcp_stack_tfo_active_conn_limit").(int)
	c.IcmpLockup = d.Get("icmp_lockup").(int)
	c.IcmpRateLimit = d.Get("icmp_rate_limit").(int)

	vc.UUID = c
	return vc
}
