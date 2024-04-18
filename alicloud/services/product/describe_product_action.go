package product

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type ProductActionResponse struct {
	Code int              `json:"code"`
	Data []*ProductAction `json:"data"`
}

type ProductAction struct {
	DirID     int                 `json:"dir_id"`
	Orders    int                 `json:"orders"`
	APIType   string              `json:"api_type"`
	NodeTitle string              `json:"node_title"`
	Title     string              `json:"title"`
	Children  []*ProductActionSet `json:"children"`
}

type ProductActionSet struct {
	ID             string              `json:"id"`
	Product        string              `json:"product"`
	Name           string              `json:"name"`
	ProductVersion string              `json:"version"`
	Method         string              `json:"method"`
	Hidden         bool                `json:"hidden"`
	Deprecated     bool                `json:"deprecated"`
	UpdatedAt      time.Time           `json:"updatedAt"`
	CreatedAt      time.Time           `json:"createdAt"`
	DirID          int                 `json:"dir_id"`
	Orders         int                 `json:"orders"`
	APIType        string              `json:"api_type"`
	NodeTitle      string              `json:"node_title"`
	ParentNodeID   int                 `json:"parent_node_id"`
	DocID          int                 `json:"doc_id"`
	Title          string              `json:"title"`
	Summary        string              `json:"summary"`
	Children       []*ProductActionSet `json:"children"`
}

func DescribeProductActions(name string, version string) ([]*ProductAction, error) {
	client := resty.New()

	r, err := client.R().
		SetQueryParams(map[string]string{
			"product": name,
			"version": version,
		}).
		Get("https://api.aliyun.com/api/product/apiDir")
	if err != nil {
		log.Fatalf("client ERROR: %s", err)
		return nil, nil
	}
	content := r.Body()
	var resp ProductActionResponse
	resp.FromJsonByte(content)
	if resp.Code != 0 {
		// todo: add log
		return nil, nil
	}

	return resp.Data, nil
}

func (r *ProductActionResponse) FromJsonByte(data []byte) error {
	f := actionResponse{}
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	return json.Unmarshal(data, &r)
}
