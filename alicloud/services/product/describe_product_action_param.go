package product

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type ProductActionParamResponse struct {
	Code int                 `json:"code"`
	Data *ProductActionParam `json:"data"`
}

type ProductActionParam struct {
	Product       string `json:"product"`
	Name          string `json:"name"`
	Version       string `json:"version"`
	ParameterType string `json:"parameter_type"`
	// Params        string `json:"params"`
	Method   string `json:"method"`
	Protocol string `json:"protocol"`
	Path     string `json:"path"`
	// Business           interface{}   `json:"business"`
	// Response           string        `json:"response"`
	// ResponseHeaders    string        `json:"response_headers"`
	Timeout    int  `json:"timeout"`
	Deprecated bool `json:"deprecated"`
	Hidden     bool `json:"hidden"`
	// ResourceTypeCode   interface{}   `json:"resource_type_code"`
	// ResourceTypeTitle  interface{}   `json:"resource_type_title"`
	// BodyStyle          interface{}   `json:"body_style"`
	SiteProps string    `json:"site_props"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Site      string    `json:"site"`
	Summary   string    `json:"summary"`
	// Request            string        `json:"request"`
	Title         string `json:"title"`
	Api           string `json:"api"`
	Description   string `json:"description"`
	RecordVersion int    `json:"record_version"`
	// AbcId              interface{}   `json:"abc_id"`
	// ContentMd          interface{}   `json:"content_md"`
	ExtraInfo         string `json:"extra_info"`
	RequestDemo       string `json:"request_demo"`
	RequestHeaderDesc string `json:"request_header_desc"`
	RequestParamsDesc string `json:"request_params_desc"`
	ResponseDemoDesc  string `json:"response_demo_desc"`
	// ErrorCodes         string        `json:"error_codes"`
	// ResponseHeadersDocs string       `json:"responseHeadersDocs"`
	// ResponseDemo       string        `json:"response_demo"`
	FormBody   bool       `json:"formBody"`
	Runtime    string     `json:"runtime"`
	Security   []Security `json:"security"`
	HostEnable bool       `json:"hostEnable"`
}

type Security struct {
	AK []string `json:"AK"`
}

func DescribeProductActionParams(name string, version string, action string) (*ProductActionParam, error) {
	client := resty.New()

	r, err := client.R().
		SetQueryParams(map[string]string{
			"productName": name,
			"version":     version,
			"apiName":     action,
		}).
		Get("https://api.aliyun.com/api/product/api")
	if err != nil {
		log.Fatalf("client ERROR: %s", err)
		return nil, nil
	}
	content := r.Body()
	var resp ProductActionParamResponse
	resp.FromJsonByte(content)
	if resp.Code != 0 {
		// todo: add log
		return nil, nil
	}

	return resp.Data, nil
}

func (r *ProductActionParamResponse) FromJsonByte(data []byte) error {
	f := actionResponse{}
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	return json.Unmarshal(data, &r)
}
