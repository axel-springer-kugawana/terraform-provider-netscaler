# netscaler_tmsessionpolicy

Terraform resource name : ```netscaler_tmsessionpolicy```

##### Arguments

| Name | Force new | Array | Type |
|----|----|----|----|
|action|No|No|[tmsessionaction.name](/doc/resources/tmsessionaction.md)|
|name|No|No|string|
|rule|No|No|string|

##### Argument Reference

See [official documentation page](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/traffic-management/tmsessionpolicy/tmsessionpolicy/) for possible values for these arguments and for an exhaustive list of arguments.

##### Example

```
resource "netscaler_tmsessionpolicy" "<resource_name>" {

    action = "${netscaler_tmsessionaction.<resource_name>.name}"
    name = "abc"
    rule = "abc"
}
```

