package bixgo

import (
	"context"
	"encoding/json"
	"fmt"
)

type Section struct {
	Id              int            `json:"id"`
	IblockId        int            `json:"iblockId"`
	IblockSectionId *int           `json:"iblockSectionId"`
	Name            string         `json:"name"`
	XMLID           string         `json:"xmlId"`
	Code            string         `json:"code"`
	Sort            int            `json:"sort"`
	Active          BitrixBoolean  `json:"active"`
	Description     string         `json:"description"`
	DescriptionType string         `json:"descriptionType"`
	CustomFields    map[string]any `json:"customFields"`
}

type SectionsResponse struct {
	Sections []Section `json:"sections"`
}

const sectionsListMethod = "catalog.section.list"

var sectionKnownFields = map[string]struct{}{
	"id":              {},
	"iblockId":        {},
	"iblockSectionId": {},
	"name":            {},
	"xmlId":           {},
	"code":            {},
	"sort":            {},
	"active":          {},
	"description":     {},
	"descriptionType": {},
}

func (c *Section) UnmarshalJSON(data []byte) error {
	type Alias Section
	if err := json.Unmarshal(data, (*Alias)(c)); err != nil {
		return err
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	customFieldsCount := len(raw) - len(sectionKnownFields)
	if customFieldsCount <= 0 {
		return nil
	}

	c.CustomFields = make(map[string]any, customFieldsCount)
	for key, val := range raw {
		if _, found := sectionKnownFields[key]; !found {
			var v any
			if err := json.Unmarshal(val, &v); err != nil {
				return fmt.Errorf("unmarshal custom field %s: %w", key, err)
			}
			c.CustomFields[key] = v
		}
	}

	return nil
}

/*
GetCatalogSections

	Возвращает список разделов торгового каталога, используя метод catalog.section.list
*/
func (c *Client) GetCatalogSections(
	ctx context.Context,
	params Params,
) (*ListResponse[CatalogsResponse], error) {
	rawResponse, err := c.callRaw(ctx, sectionsListMethod, params)
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
