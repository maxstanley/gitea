package structs

type ActionRun struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Repo        *Repository  `json:"repo"`
	Owner       *User        `json:"owner"`
	TriggerUser *User        `json:"trigger_user"`
	WorkflowID  string       `json:"workflow_id"`
	Index       int64        `json:"index"`
	Ref         string       `json:"ref"`
	CommitSHA   string       `json:"commit_sha"`
	Status      string       `json:"status"`
	Version     int          `json:"version"`
	Started     string       `json:"time_started"`
	Stopped     string       `json:"time_stopped"`
	Created     string       `json:"time_created"`
	Updated     string       `json:"time_updated"`
}
