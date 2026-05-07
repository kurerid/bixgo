package bixgo

import (
	"context"
	"encoding/json"
)

type CatalogsResponse struct {
	Catalogs []any `json:"catalogs"`
}

const catalogsListMethod = "catalog.catalog.list"

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
