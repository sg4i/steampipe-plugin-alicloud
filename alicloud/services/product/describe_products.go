package product

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type actionResponse map[string]interface{}

type ProductResponse struct {
	Code int        `json:"code"`
	Data []*Product `json:"data"`
}

type ProductDescription struct {
	Cn string `json:"zh"`
	En string `json:"cn"`
}

type Product struct {
	Id                 int64              `json:"id"`
	Name               string             `json:"name"`
	Versions           string             `json:"versions"`
	PublicVersion      string             `json:"public_version"`
	DefaultVersion     string             `json:"default_version"`
	AllVersions        string             `json:"all_versions"`
	DailyVersions      string             `json:"daily_versions"`
	PreVersions        string             `json:"pre_versions"`
	ShowNameCn         string             `json:"show_name_cn"`
	ShowNameEn         string             `json:"show_name_en"`
	ApiDoc             string             `json:"api_doc"`
	SdkDoc             string             `json:"sdk_doc"`
	Endpoint           Endpoint           `json:"endpoint"`
	BelongGroup        string             `json:"belong_group"`
	IsPublic           bool               `json:"is_public"`
	ApiStyle           string             `json:"api_style"`
	ShortName          string             `json:"short_name"`
	Status             int                `json:"status"`
	OpenUrl            string             `json:"open_url"`
	OpenUrlIntl        string             `json:"open_url_intl"`
	IsNeedActive       bool               `json:"is_need_active"`
	ActiveJson         string             `json:"active_json"`
	Description        ProductDescription `json:"description"`
	PipProductCode     string             `json:"pip_product_code"`
	DefaultPipCode     string             `json:"default_pip_code"`
	LocationCode       string             `json:"location_code"`
	SupportCli         bool               `json:"support_cli"`
	SupportSdkLanguage string             `json:"support_sdk_language"`
	SupportPopDoc      bool               `json:"support_pop_doc"`
	Keywords           string             `json:"keywords"`
	KeywordsEn         string             `json:"keywords_en"`
	RamSupport         []string           `json:"ram_support"`
	HiddenDara         int                `json:"hidden_dara"`
	HiddenApi          string             `json:"hidden_api"`
	SiteType           string             `json:"site_type"`
	CreatedAt          time.Time          `json:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt"`
	CategoryId         int64              `json:"categoryId"`
	CategoryCode       string             `json:"categoryCode"`
	CategoryName       string             `json:"categoryName"`
	CategoryNameEn     string             `json:"categoryNameEn"`
	Category2Id        int64              `json:"category2Id"`
	Category2Code      string             `json:"category2Code"`
	Category2Name      string             `json:"category2Name"`
	Category2NameEn    string             `json:"category2NameEn"`
}

func DescribeProducts() ([]*Product, error) {
	client := resty.New()

	r, err := client.R().
		Get("https://api.aliyun.com/api/products")
	if err != nil {
		log.Fatalf("client ERROR: %s", err)
		return nil, nil
	}
	content := r.Body()
	var resp ProductResponse
	resp.FromJsonByte(content)
	if resp.Code != 0 {
		// todo: add log
		return nil, nil
	}

	return resp.Data, nil
}

type Endpoint map[string]string

func (e *Endpoint) UnmarshalJSON(data []byte) error {
	// 解析外层的JSON字符串
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// 解析内层的JSON字符串（地图）
	tmp := make(map[string]string)
	if err := json.Unmarshal([]byte(str), &tmp); err != nil {
		return err
	}

	*e = tmp
	return nil
}

func (r *ProductResponse) FromJsonByte(data []byte) error {
	f := actionResponse{}
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	return json.Unmarshal(data, &r)
}
