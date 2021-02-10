package model

type SendMessages struct {
	Messages []Messages `json:"messages"`
}

type Messages struct {
	From         string         `json:"from"`
	Text         string         `json:"text"`
	Destinations []Destinations `json:"destinations"`
}

type Destinations struct {
	To string `json:"to"`
}

type SentResponse struct {
	BulkId   string                 `json:"bulkId"`
	Messages []MessagesSentResponse `json:"messages"`
}

type MessagesSentResponse struct {
	To        string `json:"to"`
	Status    Status `json:"status"`
	MessageId string `json:"messageId"`
}

type Status struct {
	GroupId     int64  `json:"groupId"`
	GroupName   string `json:"groupName"`
	Id          int64  `json:"id"`
	Name        string `json:"'name'"`
	Description string `json:"description"`
}
