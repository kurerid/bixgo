package bixgo

import (
	"context"
	"encoding/json"
	"fmt"
)

type Task struct {
	CustomFields map[string]any `json:"customFields"`
}

type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}

const taskListMethod = "tasks.task.list"

var taskKnownFields = map[string]struct{}{}

func (t *Task) UnmarshalJSON(data []byte) error {
	type Alias Task
	if err := json.Unmarshal(data, (*Alias)(t)); err != nil {
		return err
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	unknownCount := len(raw) - len(taskKnownFields)
	if unknownCount <= 0 {
		return nil
	}

	for key, val := range raw {
		if _, found := taskKnownFields[key]; found {
			continue
		}

		var v any
		if err := json.Unmarshal(val, &v); err != nil {
			return fmt.Errorf("unmarshal field %s: %w", key, err)
		}
		if t.CustomFields == nil {
			t.CustomFields = make(map[string]any)
		}
		t.CustomFields[key] = v

	}

	return nil
}

/*
GetTasks

	Возвращает список задач по фильтру, используя метод sonet_group.get
*/
func (c *Client) GetTasks(
	ctx context.Context,
	params Params,
) (*ListResponse[TasksResponse], error) {
	rawResponse, err := c.callRaw(ctx, taskListMethod, params)
	if err != nil {
		return nil, err
	}
	var result ListResponse[TasksResponse]
	err = json.Unmarshal(rawResponse, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
