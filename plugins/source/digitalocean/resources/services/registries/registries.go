// Code generated by codegen; DO NOT EDIT.

package registries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Registries() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_registries",
		Resolver: fetchRegistriesRegistries,
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "storage_usage_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("StorageUsageBytes"),
			},
			{
				Name:     "storage_usage_bytes_updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StorageUsageBytesUpdatedAt"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
		},

		Relations: []*schema.Table{
			Repositories(),
		},
	}
}
