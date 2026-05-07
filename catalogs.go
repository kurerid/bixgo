package bixgo

import (
	"context"
	"encoding/json"
	"fmt"
)

type Catalog struct {
	Id              int    `json:"id"`
	IblockId        int    `json:"iblockId"`
	IblockTypeId    string `json:"iblockTypeId"`
	Lid             string `json:"lid"`
	Name            string `json:"name"`
	ProductIblockId *int   `json:"productIblockId"`
	SkuPropertyId   *int   `json:"skuPropertyId"`
	Subscription    string `json:"subscription"`
	VatId           *int   `json:"vatId"`
	CustomFields    map[string]any
}

type CatalogsResponse struct {
	Catalogs []Catalog `json:"catalogs"`
}

const catalogsListMethod = "catalog.catalog.list"

var knownFields = map[string]struct{}{
	"id":              {},
	"iblockId":        {},
	"iblockTypeId":    {},
	"lid":             {},
	"name":            {},
	"productIblockId": {},
	"skuPropertyId":   {},
	"subscription":    {},
	"vatId":           {},
}

func (c *Catalog) UnmarshalJSON(data []byte) error {
	type Alias Catalog
	if err := json.Unmarshal(data, (*Alias)(c)); err != nil {
		return err
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	customFieldsCount := len(raw) - len(knownFields)
	if customFieldsCount <= 0 {
		return nil
	}

	c.CustomFields = make(map[string]any, customFieldsCount)
	for key, val := range raw {
		if _, found := knownFields[key]; !found {
			var v any
			if err := json.Unmarshal(val, &v); err != nil {
				return fmt.Errorf("unmarshal custom field %s: %w", key, err)
			}
			c.CustomFields[key] = v
		}
	}

	return nil
}

func (c *Client) GetCatalogs(
	ctx context.Context,
	params Params,
) (*ListResponse[CatalogsResponse], error) {
	rawResponse, err := c.callRaw(ctx, catalogsListMethod, params)
	if err != nil {
		return nil, err
	}
	var result ListResponse[CatalogsResponse]
	err = json.Unmarshal(rawResponse, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
