package bixgo

import (
	"context"
	"encoding/json"
	"fmt"
)

type ProjectOrGroup struct {
	Id              string         `json:"ID"`
	SiteId          string         `json:"SITE_ID"`
	Name            string         `json:"NAME"`
	Description     string         `json:"DESCRIPTION"`
	DateCreate      string         `json:"DATE_CREATE"`
	DateUpdate      string         `json:"DATE_UPDATE"`
	Active          string         `json:"ACTIVE"`
	Visible         string         `json:"VISIBLE"`
	Opened          string         `json:"OPENED"`
	Closed          string         `json:"CLOSED"`
	SubjectId       string         `json:"SUBJECT_ID"`
	OwnerId         string         `json:"OWNER_ID"`
	Keywords        string         `json:"KEYWORDS"`
	NumberOfMembers string         `json:"NUMBER_OF_MEMBERS"`
	DateActivity    string         `json:"DATE_ACTIVITY"`
	SubjectName     string         `json:"SUBJECT_NAME"`
	Project         string         `json:"PROJECT"`
	IsExtranet      string         `json:"IS_EXTRANET"`
	CustomFields    map[string]any `json:"customFields"`
}

const projectsAndGroupListMethod = "sonet_group.get"

var projectsKnownFields = map[string]struct{}{
	"ID":                {},
	"SITE_ID":           {},
	"NAME":              {},
	"DESCRIPTION":       {},
	"DATE_CREATE":       {},
	"DATE_UPDATE":       {},
	"ACTIVE":            {},
	"VISIBLE":           {},
	"OPENED":            {},
	"CLOSED":            {},
	"SUBJECT_ID":        {},
	"OWNER_ID":          {},
	"KEYWORDS":          {},
	"NUMBER_OF_MEMBERS": {},
	"DATE_ACTIVITY":     {},
	"SUBJECT_NAME":      {},
	"PROJECT":           {},
	"IS_EXTRANET":       {},
}

func (p *ProjectOrGroup) UnmarshalJSON(data []byte) error {
	type Alias ProjectOrGroup
	if err := json.Unmarshal(data, (*Alias)(p)); err != nil {
		return err
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	unknownCount := len(raw) - len(projectsKnownFields)
	if unknownCount <= 0 {
		return nil
	}

	for key, val := range raw {
		if _, found := projectsKnownFields[key]; found {
			continue
		}

		var v any
		if err := json.Unmarshal(val, &v); err != nil {
			return fmt.Errorf("unmarshal field %s: %w", key, err)
		}
		if p.CustomFields == nil {
			p.CustomFields = make(map[string]any)
		}
		p.CustomFields[key] = v

	}

	return nil
}

/*
GetProjectsAndGroups

	Возвращает список проектов и групп по фильтру, используя метод sonet_group.get
*/
func (c *Client) GetProjectsAndGroups(
	ctx context.Context,
	params Params,
) (*ListResponse[[]ProjectOrGroup], error) {
	rawResponse, err := c.callRaw(ctx, projectsAndGroupListMethod, params)
	if err != nil {
		return nil, err
	}
	var result ListResponse[[]ProjectOrGroup]
	err = json.Unmarshal(rawResponse, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
