package structs

import "code.gitea.io/gitea/modules/json"

var _ Payloader = &PullRequestPayload{}

// __________      .__  .__    __________                                     __
// \______   \__ __|  | |  |   \______   \ ____  ________ __   ____   _______/  |_
//  |     ___/  |  \  | |  |    |       _// __ \/ ____/  |  \_/ __ \ /  ___/\   __\
//  |    |   |  |  /  |_|  |__  |    |   \  ___< <_|  |  |  /\  ___/ \___ \  |  |
//  |____|   |____/|____/____/  |____|_  /\___  >__   |____/  \___  >____  > |__|
//                                     \/     \/   |__|           \/     \/

// PullRequestPayload represents a payload information of pull request event.
type PullRequestPayload struct {
	Action            HookIssueAction `json:"action"`
	Index             int64           `json:"number"`
	Changes           *ChangesPayload `json:"changes,omitempty"`
	PullRequest       *PullRequest    `json:"pull_request"`
	RequestedReviewer *User           `json:"requested_reviewer"`
	Repository        *Repository     `json:"repository"`
	Sender            *User           `json:"sender"`
	CommitID          string          `json:"commit_id"`
	Review            *ReviewPayload  `json:"review"`
}

// JSONPayload FIXME
func (p *PullRequestPayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

// ReviewPayload FIXME
type ReviewPayload struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}
