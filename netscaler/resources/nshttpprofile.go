package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerNshttpprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_nshttpprofile,
		Read:          read_nshttpprofile,
		Update:        update_nshttpprofile,
		Delete:        delete_nshttpprofile,
		Schema: map[string]*schema.Schema{
			"adpttimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"altsvc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"apdexcltresptimethreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"clientiphdrexpr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cmponpush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"conmultiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dropextracrlf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dropextradata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dropinvalreqs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2direct": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2headertablesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2initialwindowsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2maxconcurrentstreams": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2maxframesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2maxheaderlistsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2minseverconn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"incomphdrdelay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"markconnreqinval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"markhttp09inval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxheaderlen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxreq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxreusepool": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"minreusepool": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"persistentetag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"reqtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"reqtimeoutaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"reusepooltimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rtsptunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"spdy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"weblog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"websocket": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_nshttpprofile(d *schema.ResourceData) nitro.Nshttpprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Nshttpprofile{
		Adpttimeout: d.Get("adpttimeout").(string),
		Altsvc:      d.Get("altsvc").(string),
		Apdexcltresptimethreshold: d.Get("apdexcltresptimethreshold").(int),
		Clientiphdrexpr:           d.Get("clientiphdrexpr").(string),
		Cmponpush:                 d.Get("cmponpush").(string),
		Conmultiplex:              d.Get("conmultiplex").(string),
		Dropextracrlf:             d.Get("dropextracrlf").(string),
		Dropextradata:             d.Get("dropextradata").(string),
		Dropinvalreqs:             d.Get("dropinvalreqs").(string),
		Http2:                     d.Get("http2").(string),
		Http2direct:               d.Get("http2direct").(string),
		Http2headertablesize:      d.Get("http2headertablesize").(int),
		Http2initialwindowsize:    d.Get("http2initialwindowsize").(int),
		Http2maxconcurrentstreams: d.Get("http2maxconcurrentstreams").(int),
		Http2maxframesize:         d.Get("http2maxframesize").(int),
		Http2maxheaderlistsize:    d.Get("http2maxheaderlistsize").(int),
		Http2minseverconn:         d.Get("http2minseverconn").(int),
		Incomphdrdelay:            d.Get("incomphdrdelay").(int),
		Markconnreqinval:          d.Get("markconnreqinval").(string),
		Markhttp09inval:           d.Get("markhttp09inval").(string),
		Maxheaderlen:              d.Get("maxheaderlen").(int),
		Maxreq:                    d.Get("maxreq").(int),
		Maxreusepool:              d.Get("maxreusepool").(int),
		Minreusepool:              d.Get("minreusepool").(int),
		Name:                      d.Get("name").(string),
		Persistentetag:            d.Get("persistentetag").(string),
		Reqtimeout:                d.Get("reqtimeout").(int),
		Reqtimeoutaction:          d.Get("reqtimeoutaction").(string),
		Reusepooltimeout:          d.Get("reusepooltimeout").(int),
		Rtsptunnel:                d.Get("rtsptunnel").(string),
		Spdy:                      d.Get("spdy").(string),
		Weblog:                    d.Get("weblog").(string),
		Websocket:                 d.Get("websocket").(string),
	}

	return resource
}

func set_nshttpprofile(d *schema.ResourceData, resource *nitro.Nshttpprofile) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("adpttimeout", resource.Adpttimeout)
	d.Set("altsvc", resource.Altsvc)
	d.Set("apdexcltresptimethreshold", resource.Apdexcltresptimethreshold)
	d.Set("clientiphdrexpr", resource.Clientiphdrexpr)
	d.Set("cmponpush", resource.Cmponpush)
	d.Set("conmultiplex", resource.Conmultiplex)
	d.Set("dropextracrlf", resource.Dropextracrlf)
	d.Set("dropextradata", resource.Dropextradata)
	d.Set("dropinvalreqs", resource.Dropinvalreqs)
	d.Set("http2", resource.Http2)
	d.Set("http2direct", resource.Http2direct)
	d.Set("http2headertablesize", resource.Http2headertablesize)
	d.Set("http2initialwindowsize", resource.Http2initialwindowsize)
	d.Set("http2maxconcurrentstreams", resource.Http2maxconcurrentstreams)
	d.Set("http2maxframesize", resource.Http2maxframesize)
	d.Set("http2maxheaderlistsize", resource.Http2maxheaderlistsize)
	d.Set("http2minseverconn", resource.Http2minseverconn)
	d.Set("incomphdrdelay", resource.Incomphdrdelay)
	d.Set("markconnreqinval", resource.Markconnreqinval)
	d.Set("markhttp09inval", resource.Markhttp09inval)
	d.Set("maxheaderlen", resource.Maxheaderlen)
	d.Set("maxreq", resource.Maxreq)
	d.Set("maxreusepool", resource.Maxreusepool)
	d.Set("minreusepool", resource.Minreusepool)
	d.Set("name", resource.Name)
	d.Set("persistentetag", resource.Persistentetag)
	d.Set("reqtimeout", resource.Reqtimeout)
	d.Set("reqtimeoutaction", resource.Reqtimeoutaction)
	d.Set("reusepooltimeout", resource.Reusepooltimeout)
	d.Set("rtsptunnel", resource.Rtsptunnel)
	d.Set("spdy", resource.Spdy)
	d.Set("weblog", resource.Weblog)
	d.Set("websocket", resource.Websocket)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_nshttpprofile_key(d *schema.ResourceData) nitro.NshttpprofileKey {

	key := nitro.NshttpprofileKey{
		d.Get("name").(string),
	}
	return key
}

func create_nshttpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_nshttpprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_nshttpprofile(d)
	key := resource.ToKey()

	exists, err := client.ExistsNshttpprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetNshttpprofile(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_nshttpprofile(d, resource)
	} else {
		err := client.AddNshttpprofile(get_nshttpprofile(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetNshttpprofile(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_nshttpprofile(d, resource)
	}

	return nil
}

func read_nshttpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_nshttpprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_nshttpprofile(d)
	key := resource.ToKey()

	exists, err := client.ExistsNshttpprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetNshttpprofile(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_nshttpprofile(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_nshttpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_nshttpprofile")

	client := meta.(*nitro.NitroClient)

	update := nitro.NshttpprofileUpdate{}
	unset := nitro.NshttpprofileUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("dropinvalreqs") {
		updateFlag = true

		value := d.Get("dropinvalreqs").(string)
		update.Dropinvalreqs = value

		if value == "" {
			unsetFlag = true

			unset.Dropinvalreqs = true
		}

	}
	if d.HasChange("markhttp09inval") {
		updateFlag = true

		value := d.Get("markhttp09inval").(string)
		update.Markhttp09inval = value

		if value == "" {
			unsetFlag = true

			unset.Markhttp09inval = true
		}

	}
	if d.HasChange("markconnreqinval") {
		updateFlag = true

		value := d.Get("markconnreqinval").(string)
		update.Markconnreqinval = value

		if value == "" {
			unsetFlag = true

			unset.Markconnreqinval = true
		}

	}
	if d.HasChange("cmponpush") {
		updateFlag = true

		value := d.Get("cmponpush").(string)
		update.Cmponpush = value

		if value == "" {
			unsetFlag = true

			unset.Cmponpush = true
		}

	}
	if d.HasChange("conmultiplex") {
		updateFlag = true

		value := d.Get("conmultiplex").(string)
		update.Conmultiplex = value

		if value == "" {
			unsetFlag = true

			unset.Conmultiplex = true
		}

	}
	if d.HasChange("maxreusepool") {
		updateFlag = true

		value := d.Get("maxreusepool").(int)
		update.Maxreusepool = value

		if value == 0 {
			unsetFlag = true

			unset.Maxreusepool = true
		}

	}
	if d.HasChange("dropextracrlf") {
		updateFlag = true

		value := d.Get("dropextracrlf").(string)
		update.Dropextracrlf = value

		if value == "" {
			unsetFlag = true

			unset.Dropextracrlf = true
		}

	}
	if d.HasChange("incomphdrdelay") {
		updateFlag = true

		value := d.Get("incomphdrdelay").(int)
		update.Incomphdrdelay = value

		if value == 0 {
			unsetFlag = true

			unset.Incomphdrdelay = true
		}

	}
	if d.HasChange("websocket") {
		updateFlag = true

		value := d.Get("websocket").(string)
		update.Websocket = value

		if value == "" {
			unsetFlag = true

			unset.Websocket = true
		}

	}
	if d.HasChange("rtsptunnel") {
		updateFlag = true

		value := d.Get("rtsptunnel").(string)
		update.Rtsptunnel = value

		if value == "" {
			unsetFlag = true

			unset.Rtsptunnel = true
		}

	}
	if d.HasChange("reqtimeout") {
		updateFlag = true

		value := d.Get("reqtimeout").(int)
		update.Reqtimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Reqtimeout = true
		}

	}
	if d.HasChange("adpttimeout") {
		updateFlag = true

		value := d.Get("adpttimeout").(string)
		update.Adpttimeout = value

		if value == "" {
			unsetFlag = true

			unset.Adpttimeout = true
		}

	}
	if d.HasChange("reqtimeoutaction") {
		updateFlag = true

		value := d.Get("reqtimeoutaction").(string)
		update.Reqtimeoutaction = value

		if value == "" {
			unsetFlag = true

			unset.Reqtimeoutaction = true
		}

	}
	if d.HasChange("dropextradata") {
		updateFlag = true

		value := d.Get("dropextradata").(string)
		update.Dropextradata = value

		if value == "" {
			unsetFlag = true

			unset.Dropextradata = true
		}

	}
	if d.HasChange("weblog") {
		updateFlag = true

		value := d.Get("weblog").(string)
		update.Weblog = value

		if value == "" {
			unsetFlag = true

			unset.Weblog = true
		}

	}
	if d.HasChange("clientiphdrexpr") {
		updateFlag = true

		value := d.Get("clientiphdrexpr").(string)
		update.Clientiphdrexpr = value

		if value == "" {
			unsetFlag = true

			unset.Clientiphdrexpr = true
		}

	}
	if d.HasChange("maxreq") {
		updateFlag = true

		value := d.Get("maxreq").(int)
		update.Maxreq = value

		if value == 0 {
			unsetFlag = true

			unset.Maxreq = true
		}

	}
	if d.HasChange("persistentetag") {
		updateFlag = true

		value := d.Get("persistentetag").(string)
		update.Persistentetag = value

		if value == "" {
			unsetFlag = true

			unset.Persistentetag = true
		}

	}
	if d.HasChange("spdy") {
		updateFlag = true

		value := d.Get("spdy").(string)
		update.Spdy = value

		if value == "" {
			unsetFlag = true

			unset.Spdy = true
		}

	}
	if d.HasChange("http2") {
		updateFlag = true

		value := d.Get("http2").(string)
		update.Http2 = value

		if value == "" {
			unsetFlag = true

			unset.Http2 = true
		}

	}
	if d.HasChange("http2direct") {
		updateFlag = true

		value := d.Get("http2direct").(string)
		update.Http2direct = value

		if value == "" {
			unsetFlag = true

			unset.Http2direct = true
		}

	}
	if d.HasChange("altsvc") {
		updateFlag = true

		value := d.Get("altsvc").(string)
		update.Altsvc = value

		if value == "" {
			unsetFlag = true

			unset.Altsvc = true
		}

	}
	if d.HasChange("http2maxheaderlistsize") {
		updateFlag = true

		value := d.Get("http2maxheaderlistsize").(int)
		update.Http2maxheaderlistsize = value

		if value == 0 {
			unsetFlag = true

			unset.Http2maxheaderlistsize = true
		}

	}
	if d.HasChange("http2maxframesize") {
		updateFlag = true

		value := d.Get("http2maxframesize").(int)
		update.Http2maxframesize = value

		if value == 0 {
			unsetFlag = true

			unset.Http2maxframesize = true
		}

	}
	if d.HasChange("http2maxconcurrentstreams") {
		updateFlag = true

		value := d.Get("http2maxconcurrentstreams").(int)
		update.Http2maxconcurrentstreams = value

		if value == 0 {
			unsetFlag = true

			unset.Http2maxconcurrentstreams = true
		}

	}
	if d.HasChange("http2initialwindowsize") {
		updateFlag = true

		value := d.Get("http2initialwindowsize").(int)
		update.Http2initialwindowsize = value

		if value == 0 {
			unsetFlag = true

			unset.Http2initialwindowsize = true
		}

	}
	if d.HasChange("http2headertablesize") {
		updateFlag = true

		value := d.Get("http2headertablesize").(int)
		update.Http2headertablesize = value

		if value == 0 {
			unsetFlag = true

			unset.Http2headertablesize = true
		}

	}
	if d.HasChange("http2minseverconn") {
		updateFlag = true

		value := d.Get("http2minseverconn").(int)
		update.Http2minseverconn = value

		if value == 0 {
			unsetFlag = true

			unset.Http2minseverconn = true
		}

	}
	if d.HasChange("reusepooltimeout") {
		updateFlag = true

		value := d.Get("reusepooltimeout").(int)
		update.Reusepooltimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Reusepooltimeout = true
		}

	}
	if d.HasChange("maxheaderlen") {
		updateFlag = true

		value := d.Get("maxheaderlen").(int)
		update.Maxheaderlen = value

		if value == 0 {
			unsetFlag = true

			unset.Maxheaderlen = true
		}

	}
	if d.HasChange("minreusepool") {
		updateFlag = true

		value := d.Get("minreusepool").(int)
		update.Minreusepool = value

		if value == 0 {
			unsetFlag = true

			unset.Minreusepool = true
		}

	}
	if d.HasChange("apdexcltresptimethreshold") {
		updateFlag = true

		value := d.Get("apdexcltresptimethreshold").(int)
		update.Apdexcltresptimethreshold = value

		if value == 0 {
			unsetFlag = true

			unset.Apdexcltresptimethreshold = true
		}

	}
	key := get_nshttpprofile_key(d)

	if updateFlag {
		if err := client.UpdateNshttpprofile(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetNshttpprofile(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetNshttpprofile(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_nshttpprofile(d, resource)
	}

	return nil
}

func delete_nshttpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_nshttpprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_nshttpprofile(d)
	key := resource.ToKey()

	exists, err := client.ExistsNshttpprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteNshttpprofile(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
