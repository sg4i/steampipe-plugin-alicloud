package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-alicloud/alicloud/services/product"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudProductAction(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_product_action",
		Description: "Alicloud Product Action",
		List: &plugin.ListConfig{
			ParentHydrate: listProduct,
			Hydrate:       listProductAction,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "name of the product action.",
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: "title of the product action.",
				Transform:   transform.FromField("Title"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "description of the product action.",
				Transform:   transform.FromField("Summary"),
			},
			{
				Name:        "api_type",
				Type:        proto.ColumnType_STRING,
				Description: "api type of the product action.",
				Transform:   transform.FromField("APIType"),
			},
			{
				Name:        "product_name",
				Type:        proto.ColumnType_STRING,
				Description: "name of the product.",
				Transform:   transform.FromField("Product"),
			},
			{
				Name:        "product_version",
				Type:        proto.ColumnType_STRING,
				Description: "version of the product.",
				Transform:   transform.FromField("ProductVersion"),
			},
			{
				Name:        "parameter_type",
				Type:        proto.ColumnType_STRING,
				Description: "parameter type of the product action.",
				Hydrate:     getProductActionParam,
				Transform:   transform.FromField("ParameterType"),
			},
		},
	}
}

// // LIST FUNCTION
func listProductAction(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	p := h.Item.(*product.Product)
	// name := "Ess"
	// version := "2022-02-22"
	response, err := product.DescribeProductActions(p.Name, p.DefaultVersion)
	if err != nil {
		return nil, err
	}

	for _, data := range response {
		for _, item := range data.Children {
			d.StreamListItem(ctx, item)
		}
	}
	return nil, nil
}

// // HYDRATE FUNCTIONS
func getProductActionParam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	p := h.Item.(*product.ProductActionSet)
	// name := "Ess"
	// version := "2022-02-22"
	response, err := product.DescribeProductActionParams(p.Product, p.ProductVersion, p.Name)
	if err != nil {
		return nil, err
	}

	return response, nil
}
