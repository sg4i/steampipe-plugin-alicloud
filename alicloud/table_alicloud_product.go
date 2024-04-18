package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-alicloud/alicloud/services/product"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudProduct(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_product",
		Description: "Alicloud Product",
		List: &plugin.ListConfig{
			Hydrate: listProduct,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "name of the product.",
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "versions",
				Type:        proto.ColumnType_STRING,
				Description: "version of the product.",
				Transform:   transform.FromField("Versions"),
			},
			{
				Name:        "title_cn",
				Type:        proto.ColumnType_STRING,
				Description: "cn title of the product.",
				Transform:   transform.FromField("ShowNameCn"),
			},
			{
				Name:        "title_en",
				Type:        proto.ColumnType_STRING,
				Description: "en title of the product.",
				Transform:   transform.FromField("ShowNameEn"),
			},
			{
				Name:        "title_short",
				Type:        proto.ColumnType_STRING,
				Description: "short title of the product.",
				Transform:   transform.FromField("ShortName"),
			},
			{
				Name:        "endpoint",
				Type:        proto.ColumnType_JSON,
				Description: "endpoint of the product.",
				Transform:   transform.FromField("Endpoint"),
			},
		},
	}
}

// // LIST FUNCTION
func listProduct(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	response, err := product.DescribeProducts()
	if err != nil {
		return nil, err
	}

	for _, data := range response {
		d.StreamListItem(ctx, data)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS
