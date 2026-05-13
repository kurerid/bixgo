package bixgo

import "encoding/json"

type BitrixBoolean bool

func (b *BitrixBoolean) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*b = s == "Y" || s == "y"
	return nil
}

type BitrixProductPicture struct {
	Id         string `json:"id"`
	Url        string `json:"url"`
	UrlMachine string `json:"urlMachine"`
}

type BitrixChecklistItem struct {
	Id               string                    `json:"id"`
	TaskId           *string                   `json:"taskId"`
	CreatedBy        *string                   `json:"createdBy"`
	ParentId         *string                   `json:"parentId"`
	Title            string                    `json:"title"`
	SortIndex        string                    `json:"sortIndex"`
	IsComplete       BitrixBoolean             `json:"isComplete"`
	IsImportant      BitrixBoolean             `json:"isImportant"`
	ToggledBy        *string                   `json:"toggledBy"`
	ToggledDate      *string                   `json:"toggledDate"`
	UfChecklistFiles bool                      `json:"ufChecklistFiles"`
	Members          []BitrixUserDescription   `json:"members"`
	Attachments      []any                     `json:"attachments"`
	EntityId         *string                   `json:"entityId"`
	Action           BitrixChecklistItemAction `json:"action"`
}

type BitrixChecklistItemAction struct {
	Modify        bool `json:"modify"`
	Remove        bool `json:"remove"`
	Toggle        bool `json:"toggle"`
	Add           bool `json:"add"`
	AddAccomplice bool `json:"addAccomplice"`
}

type BitrixGroup struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Opened         bool    `json:"opened"`
	MembersCount   int     `json:"membersCount"`
	Image          *string `json:"image"`
	AdditionalData any     `json:"additionalData"`
}

type BitrixUserDescription struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Link         string  `json:"link"`
	Icon         *string `json:"icon"`
	WorkPosition *string `json:"workPosition"`
}

type BitrixAction struct {
	Accept             bool `json:"accept"`
	Decline            bool `json:"decline"`
	Complete           bool `json:"complete"`
	Approve            bool `json:"approve"`
	Disapprove         bool `json:"disapprove"`
	Start              bool `json:"start"`
	Pause              bool `json:"pause"`
	Delegate           bool `json:"delegate"`
	Remove             bool `json:"remove"`
	Edit               bool `json:"edit"`
	Defer              bool `json:"defer"`
	Renew              bool `json:"renew"`
	Create             bool `json:"create"`
	ChangeDeadline     bool `json:"changeDeadline"`
	ChecklistAddItems  bool `json:"checklistAddItems"`
	AddFavorite        bool `json:"addFavorite"`
	DeleteFavorite     bool `json:"deleteFavorite"`
	Rate               bool `json:"rate"`
	Take               bool `json:"take"`
	EditOriginator     bool `json:"edit.originator"`
	ChecklistReorder   bool `json:"checklist.reorder"`
	ElapsedTimeAdd     bool `json:"elapsedtime.add"`
	DayplanTimerToggle bool `json:"dayplan.timer.toggle"`
	EditPlan           bool `json:"edit.plan"`
	ChecklistAdd       bool `json:"checklist.add"`
	FavoriteAdd        bool `json:"favorite.add"`
	FavoriteDelete     bool `json:"favorite.delete"`
}

type CheckListTree struct {
	NodeId      int                 `json:"nodeId"`
	Fields      CheckListTreeFields `json:"fields"`
	Action      []BitrixAction      `json:"action"`
	Descendants []CheckListTree     `json:"descendants"`
}

type CheckListTreeFields struct {
	Id               string                  `json:"id"`
	CopiedId         *string                 `json:"copiedId"`
	EntityId         *string                 `json:"entityId"`
	UserId           *int                    `json:"userId"`
	CreatedBy        *string                 `json:"createdBy"`
	ParentId         *string                 `json:"parentId"`
	Title            string                  `json:"title"`
	SortIndex        *string                 `json:"sortIndex"`
	DisplaySortIndex *string                 `json:"displaySortIndex"`
	IsComplete       bool                    `json:"isComplete"`
	IsImportant      bool                    `json:"isImportant"`
	CompletedCount   int                     `json:"completedCount"`
	Members          []BitrixUserDescription `json:"members"`
	Attachments      []any                   `json:"attachments"`
	NodeId           *string                 `json:"nodeId"`
}
