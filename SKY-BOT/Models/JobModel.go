package Models

import "time"

type Metadata struct {
	ReleaseNumber string `json:"release_number"`
	Label         string `json:"label"`
	Version       string `json:"version"`
	Vendor        string `json:"vendor"`
	Prefix        string `json:"prefix"`
}

type Task struct {
	JobID     string    `json:"job_id"`
	CreatedAt time.Time `gorm:"primaryKey" json:"created_at"`
	Type      string    `json:"type"`
	Data      Metadata  `gorm:"embedded;embeddedPrefix:data_" json:"data"`
}

var JobChannel = make(chan Task)

type WebhookResponse struct {
	JobID       string `json:"job_id"`
	Status      string `json:"status"`
	ContainerID string `json:"container_id"`
}

// Custom table name
var TaskTableName = "queued_jobs" // Default table

func (Task) TableName() string {
	return TaskTableName
}
