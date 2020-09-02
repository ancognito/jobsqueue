package jobs

type ID int

type Type string

const (
	TimeCritical    Type = "TIME_CRITICAL"
	NotTimeCritical Type = "NOT_TIME_CRITICAL"
)

type Status string

const (
	Queued     Status = "QUEUED"
	InProgress Status = "IN_PROGRESS"
	Concluded  Status = "CONCLUDED"
)

type Job struct {
	ID     ID     `json:"ID"`
	Type   Type   `json:"Type"`
	Status Status `json:"Status"`
}
