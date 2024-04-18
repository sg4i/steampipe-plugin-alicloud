package api_tests

import (
	"testing"

	"github.com/turbot/steampipe-plugin-alicloud/alicloud/services/product"
)

func TestDescribeDescribeProducts(t *testing.T) {
	product.DescribeProducts()
}
