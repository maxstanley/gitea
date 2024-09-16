package structs

import "code.gitea.io/gitea/modules/json"

var _ Payloader = &WikiPayload{}

//  __      __.__ __   .__
// /  \    /  \__|  | _|__|
// \   \/\/   /  |  |/ /  |
//  \        /|  |    <|  |
//   \__/\  / |__|__|_ \__|
//        \/          \/

// HookWikiAction an action that happens to a wiki page
type HookWikiAction string

const (
	// HookWikiCreated created
	HookWikiCreated HookWikiAction = "created"
	// HookWikiEdited edited
	HookWikiEdited HookWikiAction = "edited"
	// HookWikiDeleted deleted
	HookWikiDeleted HookWikiAction = "deleted"
)

// WikiPayload payload for repository webhooks
type WikiPayload struct {
	Action     HookWikiAction `json:"action"`
	Repository *Repository    `json:"repository"`
	Sender     *User          `json:"sender"`
	Page       string         `json:"page"`
	Comment    string         `json:"comment"`
}

// JSONPayload JSON representation of the payload
func (p *WikiPayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", " ")
}
