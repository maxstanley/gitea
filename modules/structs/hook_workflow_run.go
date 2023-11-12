package structs

import "code.gitea.io/gitea/modules/json"

var _ Payloader = &WorkflowRunPayload{}

type HookWorkflowRunAction string

const (
	HookWorkflowRunRequested  HookWorkflowRunAction = "requested"
	HookWorkflowRunInProgress HookWorkflowRunAction = "in_progress"
	HookWorkflowRunCompleted  HookWorkflowRunAction = "completed"
)

type WorkflowRunPayload struct {
	Action     HookWorkflowRunAction `json:"action"`
	Repository *Repository           `json:"repository"`
	Sender     *User                 `json:"sender"`
	ActionRun  *ActionRun            `json:"action_run"`
	Artifacts  *[]ArtifactPayload    `json:"artifacts"`
}

func (p *WorkflowRunPayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}
