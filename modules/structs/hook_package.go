package structs

import "code.gitea.io/gitea/modules/json"

var _ Payloader = &PackagePayload{}

// HookPackageAction an action that happens to a package
type HookPackageAction string

const (
	// HookPackageCreated created
	HookPackageCreated HookPackageAction = "created"
	// HookPackageDeleted deleted
	HookPackageDeleted HookPackageAction = "deleted"
)

// PackagePayload represents a package payload
type PackagePayload struct {
	Action       HookPackageAction `json:"action"`
	Repository   *Repository       `json:"repository"`
	Package      *Package          `json:"package"`
	Organization *User             `json:"organization"`
	Sender       *User             `json:"sender"`
}

// JSONPayload implements Payload
func (p *PackagePayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}
