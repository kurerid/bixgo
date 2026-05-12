package bixgo

import (
	"context"
	"encoding/json"
	"fmt"
)

type Task struct {
	Id                   string                  `json:"id"`
	ParentId             string                  `json:"parentId"`
	Title                string                  `json:"title"`
	Description          string                  `json:"description"`
	ChatId               *int                    `json:"chatId"`
	Mark                 string                  `json:"mark"`
	Priority             string                  `json:"priority"`
	Multitask            BitrixBoolean           `json:"multitask"`
	NotViewed            BitrixBoolean           `json:"notViewed"`
	Replicate            BitrixBoolean           `json:"replicate"`
	StageId              *string                 `json:"stageId"`
	SprintId             *string                 `json:"sprintId"`
	BacklogId            *string                 `json:"backlogId"`
	CreatedBy            *string                 `json:"createdBy"`
	CreatedDate          *string                 `json:"createdDate"`
	ResponsibleId        *string                 `json:"responsibleId"`
	ChangedBy            *string                 `json:"changedBy"`
	ChangedDate          *string                 `json:"changedDate"`
	StatusChangedBy      *string                 `json:"statusChangedBy"`
	ClosedBy             *string                 `json:"closedBy"`
	ClosedDate           *string                 `json:"closedDate"`
	ActivityDate         *string                 `json:"activityDate"`
	DateStart            *string                 `json:"dateStart"`
	Deadline             *string                 `json:"deadline"`
	StartDatePlan        *string                 `json:"startDatePlan"`
	EndDatePlan          *string                 `json:"endDatePlan"`
	GUID                 *string                 `json:"guid"`
	XMLID                *string                 `json:"xmlId"`
	CommentsCount        *string                 `json:"commentsCount"`
	ServiceCommentsCount *string                 `json:"serviceCommentsCount"`
	AllowChangeDeadline  BitrixBoolean           `json:"allowChangeDeadline"`
	AllowTimeTracking    BitrixBoolean           `json:"allowTimeTracking"`
	TaskControl          BitrixBoolean           `json:"taskControl"`
	AddInReport          BitrixBoolean           `json:"addInReport"`
	ForkedByTemplateId   *string                 `json:"forkedByTemplateId"`
	TimeEstimate         *string                 `json:"timeEstimate"`
	TimeSpentInLogs      *string                 `json:"timeSpentInLogs"`
	MatchWorkTime        BitrixBoolean           `json:"matchWorkTime"`
	ForumTopicId         *string                 `json:"forumTopicId"`
	ForumId              *string                 `json:"forumId"`
	SiteId               *string                 `json:"siteId"`
	Subordinate          BitrixBoolean           `json:"subordinate"`
	ExchangeModified     *string                 `json:"exchangeModified"`
	ExchangeId           *string                 `json:"exchangeId"`
	OutlookVersion       *string                 `json:"outlookVersion"`
	ViewedDate           *string                 `json:"viewedDate"`
	Sorting              *string                 `json:"sorting"`
	DurationFact         *string                 `json:"durationFact"`
	IsMuted              BitrixBoolean           `json:"isMuted"`
	IsPinned             BitrixBoolean           `json:"isPinned"`
	IsPinnedInGroup      BitrixBoolean           `json:"isPinnedInGroup"`
	FlowId               *string                 `json:"flowId"`
	DescriptionInBbcode  BitrixBoolean           `json:"descriptionInBbcode"`
	Status               string                  `json:"status"`
	StatusChangedDate    *string                 `json:"statusChangedDate"`
	DurationPlan         *string                 `json:"durationPlan"`
	DurationType         *string                 `json:"durationType"`
	Favorite             BitrixBoolean           `json:"favorite"`
	GroupId              *string                 `json:"groupId"`
	Auditors             []string                `json:"auditors"`
	Accomplices          []string                `json:"accomplices"`
	Checklist            []BitrixChecklistItem   `json:"checklist"`
	Group                BitrixGroup             `json:"group"`
	Creator              BitrixUserDescription   `json:"creator"`
	Responsible          BitrixUserDescription   `json:"responsible"`
	AccomplicesData      []BitrixUserDescription `json:"accomplicesData"`
	AuditorsData         any                     `json:"auditorsData"`
	NewCommentsCount     int                     `json:"newCommentsCount"`
	Action               BitrixAction            `json:"action"`
	CheckListTree        CheckListTree           `json:"checkListTree"`
	CheckListCanAdd      BitrixBoolean           `json:"checkListCanAdd"`
	UFCRMTask            any                     `json:"ufCrmTask"`
	UFTaskWebDavFiles    any                     `json:"ufTaskWebdavFiles"`
	UfMailMessage        *string                 `json:"ufMailMessage"`
	CustomFields         map[string]any          `json:"customFields"`
}

type TasksResponse struct {
	Tasks []Task `json:"tasks"`
}

const taskListMethod = "tasks.task.list"

var taskKnownFields = map[string]struct{}{
	"id": {},
}

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
