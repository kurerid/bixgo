package bixgo

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type Product struct {
	Id                 int                   `json:"id"`
	IblockId           int                   `json:"iblockId"`
	Name               string                `json:"name"`
	Active             BitrixBoolean         `json:"active"`
	Available          BitrixBoolean         `json:"available"`
	Code               string                `json:"code"`
	XMLID              string                `json:"xmlId"`
	BarcodeMulti       string                `json:"barcodeMulti"`
	Bundle             BitrixBoolean         `json:"bundle"`
	CanBuyZero         BitrixBoolean         `json:"canBuyZero"`
	CreatedBy          int                   `json:"createdBy"`
	ModifiedBy         int                   `json:"modifiedBy"`
	DateActiveFrom     string                `json:"dateActiveFrom"`
	DateActiveTo       string                `json:"dateActiveTo"`
	DateCreate         string                `json:"dateCreate"`
	TimestampX         string                `json:"timestampX"`
	IblockSectionId    *int                  `json:"iblockSectionId"`
	Measure            *int                  `json:"measure"`
	PreviewText        string                `json:"previewText"`
	DetailText         string                `json:"detailText"`
	PreviewPicture     *BitrixProductPicture `json:"previewPicture"`
	DetailPicture      *BitrixProductPicture `json:"detailPicture"`
	PreviewTextType    string                `json:"previewTextType"`
	DetailTextType     string                `json:"detailTextType"`
	Sort               int                   `json:"sort"`
	Subscribe          string                `json:"subscribe"`
	VatId              *int                  `json:"vatId"`
	VatIncluded        BitrixBoolean         `json:"vatIncluded"`
	Height             *float64              `json:"height"`
	Length             *float64              `json:"length"`
	Weight             *float64              `json:"weight"`
	Width              *float64              `json:"width"`
	QuantityTrace      string                `json:"quantityTrace"`
	Type               int                   `json:"type"`
	PurchasingCurrency string                `json:"purchasingCurrency"`
	PurchasingPrice    *float64              `json:"purchasingPrice"`
	Quantity           *float64              `json:"quantity"`
	QuantityReserved   *float64              `json:"quantityReserved"`
	RecurSchemeLength  *int                  `json:"recurSchemeLength"`
	RecurSchemeType    string                `json:"recurSchemeType"`
	TrialPriceId       *int                  `json:"trialPriceId"`
	WithoutOrder       BitrixBoolean         `json:"withoutOrder"`
	Properties         map[string]any        `json:"properties"`
	CustomFields       map[string]any        `json:"customFields"`
}

type ProductsResponse struct {
	Products []Product `json:"products"`
}

const productsListMethod = "catalog.product.list"

var productsKnownFields = map[string]struct{}{
	"id":                 {},
	"iblockId":           {},
	"name":               {},
	"active":             {},
	"available":          {},
	"code":               {},
	"xmlId":              {},
	"barcodeMulti":       {},
	"bundle":             {},
	"canBuyZero":         {},
	"createdBy":          {},
	"modifiedBy":         {},
	"dateActiveFrom":     {},
	"dateActiveTo":       {},
	"dateCreate":         {},
	"timestampX":         {},
	"iblockSectionId":    {},
	"measure":            {},
	"previewText":        {},
	"detailText":         {},
	"previewPicture":     {},
	"detailPicture":      {},
	"previewTextType":    {},
	"detailTextType":     {},
	"sort":               {},
	"subscribe":          {},
	"vatId":              {},
	"vatIncluded":        {},
	"height":             {},
	"length":             {},
	"weight":             {},
	"width":              {},
	"quantityTrace":      {},
	"type":               {},
	"purchasingCurrency": {},
	"purchasingPrice":    {},
	"quantity":           {},
	"quantityReserved":   {},
	"recurSchemeLength":  {},
	"recurSchemeType":    {},
	"trialPriceId":       {},
	"withoutOrder":       {},
}

func (p *Product) UnmarshalJSON(data []byte) error {
	type Alias Product
	if err := json.Unmarshal(data, (*Alias)(p)); err != nil {
		return err
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	unknownCount := len(raw) - len(productsKnownFields)
	if unknownCount <= 0 {
		return nil
	}

	for key, val := range raw {
		if _, found := productsKnownFields[key]; found {
			continue
		}

		var v any
		if err := json.Unmarshal(val, &v); err != nil {
			return fmt.Errorf("unmarshal field %s: %w", key, err)
		}

		if strings.HasPrefix(key, "property") {
			if p.Properties == nil {
				p.Properties = make(map[string]any)
			}
			propertyID := strings.TrimPrefix(key, "property")
			p.Properties[propertyID] = v
		} else {
			if p.CustomFields == nil {
				p.CustomFields = make(map[string]any)
			}
			p.CustomFields[key] = v
		}
	}

	return nil
}

/*
GetProducts

	Возвращает список товаров по фильтру, используя метод catalog.product.list
*/
func (c *Client) GetProducts(
	ctx context.Context,
	params Params,
) (*ListResponse[ProductsResponse], error) {
	rawResponse, err := c.callRaw(ctx, productsListMethod, params)
	if err != nil {
		return nil, err
	}
	var result ListResponse[ProductsResponse]
	err = json.Unmarshal(rawResponse, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
