package structs

import (
	"strings"
	"time"

	"code.gitea.io/gitea/modules/json"
)

var (
	_ Payloader = &CreatePayload{}
	_ Payloader = &DeletePayload{}
	_ Payloader = &ForkPayload{}
	_ Payloader = &PushPayload{}
	_ Payloader = &RepositoryPayload{}
)

// PayloadUser represents the author or committer of a commit
type PayloadUser struct {
	// Full name of the commit author
	Name string `json:"name"`
	// swagger:strfmt email
	Email    string `json:"email"`
	UserName string `json:"username"`
}

// FIXME: consider using same format as API when commits API are added.
//        applies to PayloadCommit and PayloadCommitVerification

// PayloadCommit represents a commit
type PayloadCommit struct {
	// sha1 hash of the commit
	ID           string                     `json:"id"`
	Message      string                     `json:"message"`
	URL          string                     `json:"url"`
	Author       *PayloadUser               `json:"author"`
	Committer    *PayloadUser               `json:"committer"`
	Verification *PayloadCommitVerification `json:"verification"`
	// swagger:strfmt date-time
	Timestamp time.Time `json:"timestamp"`
	Added     []string  `json:"added"`
	Removed   []string  `json:"removed"`
	Modified  []string  `json:"modified"`
}

// PayloadCommitVerification represents the GPG verification of a commit
type PayloadCommitVerification struct {
	Verified  bool         `json:"verified"`
	Reason    string       `json:"reason"`
	Signature string       `json:"signature"`
	Signer    *PayloadUser `json:"signer"`
	Payload   string       `json:"payload"`
}

// _________                        __
// \_   ___ \_______   ____ _____ _/  |_  ____
// /    \  \/\_  __ \_/ __ \\__  \\   __\/ __ \
// \     \____|  | \/\  ___/ / __ \|  | \  ___/
//  \______  /|__|    \___  >____  /__|  \___  >
//         \/             \/     \/          \/

// CreatePayload FIXME
type CreatePayload struct {
	Sha     string      `json:"sha"`
	Ref     string      `json:"ref"`
	RefType string      `json:"ref_type"`
	Repo    *Repository `json:"repository"`
	Sender  *User       `json:"sender"`
}

// JSONPayload return payload information
func (p *CreatePayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

// ParseCreateHook parses create event hook content.
func ParseCreateHook(raw []byte) (*CreatePayload, error) {
	hook := new(CreatePayload)
	if err := json.Unmarshal(raw, hook); err != nil {
		return nil, err
	}

	// it is possible the JSON was parsed, however,
	// was not from Gogs (maybe was from Bitbucket)
	// So we'll check to be sure certain key fields
	// were populated
	switch {
	case hook.Repo == nil:
		return nil, ErrInvalidReceiveHook
	case len(hook.Ref) == 0:
		return nil, ErrInvalidReceiveHook
	}
	return hook, nil
}

// ________         .__          __
// \______ \   ____ |  |   _____/  |_  ____
//  |    |  \_/ __ \|  | _/ __ \   __\/ __ \
//  |    `   \  ___/|  |_\  ___/|  | \  ___/
// /_______  /\___  >____/\___  >__|  \___  >
//         \/     \/          \/          \/

// PusherType define the type to push
type PusherType string

// describe all the PusherTypes
const (
	PusherTypeUser PusherType = "user"
)

// DeletePayload represents delete payload
type DeletePayload struct {
	Ref        string      `json:"ref"`
	RefType    string      `json:"ref_type"`
	PusherType PusherType  `json:"pusher_type"`
	Repo       *Repository `json:"repository"`
	Sender     *User       `json:"sender"`
}

// JSONPayload implements Payload
func (p *DeletePayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

// ___________           __
// \_   _____/__________|  | __
//  |    __)/  _ \_  __ \  |/ /
//  |     \(  <_> )  | \/    <
//  \___  / \____/|__|  |__|_ \
//      \/                   \/

// ForkPayload represents fork payload
type ForkPayload struct {
	Forkee *Repository `json:"forkee"`
	Repo   *Repository `json:"repository"`
	Sender *User       `json:"sender"`
}

// JSONPayload implements Payload
func (p *ForkPayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

// __________             .__
// \______   \__ __  _____|  |__
//  |     ___/  |  \/  ___/  |  \
//  |    |   |  |  /\___ \|   Y  \
//  |____|   |____//____  >___|  /
//                      \/     \/

// PushPayload represents a payload information of push event.
type PushPayload struct {
	Ref          string           `json:"ref"`
	Before       string           `json:"before"`
	After        string           `json:"after"`
	CompareURL   string           `json:"compare_url"`
	Commits      []*PayloadCommit `json:"commits"`
	TotalCommits int              `json:"total_commits"`
	HeadCommit   *PayloadCommit   `json:"head_commit"`
	Repo         *Repository      `json:"repository"`
	Pusher       *User            `json:"pusher"`
	Sender       *User            `json:"sender"`
}

// JSONPayload FIXME
func (p *PushPayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

// ParsePushHook parses push event hook content.
func ParsePushHook(raw []byte) (*PushPayload, error) {
	hook := new(PushPayload)
	if err := json.Unmarshal(raw, hook); err != nil {
		return nil, err
	}

	switch {
	case hook.Repo == nil:
		return nil, ErrInvalidReceiveHook
	case len(hook.Ref) == 0:
		return nil, ErrInvalidReceiveHook
	}
	return hook, nil
}

// Branch returns branch name from a payload
func (p *PushPayload) Branch() string {
	return strings.ReplaceAll(p.Ref, "refs/heads/", "")
}

//__________                           .__  __
//\______   \ ____ ______   ____  _____|__|/  |_  ___________ ___.__.
// |       _// __ \\____ \ /  _ \/  ___/  \   __\/  _ \_  __ <   |  |
// |    |   \  ___/|  |_> >  <_> )___ \|  ||  | (  <_> )  | \/\___  |
// |____|_  /\___  >   __/ \____/____  >__||__|  \____/|__|   / ____|
//        \/     \/|__|              \/                       \/

// HookRepoAction an action that happens to a repo
type HookRepoAction string

const (
	// HookRepoCreated created
	HookRepoCreated HookRepoAction = "created"
	// HookRepoDeleted deleted
	HookRepoDeleted HookRepoAction = "deleted"
)

// RepositoryPayload payload for repository webhooks
type RepositoryPayload struct {
	Action       HookRepoAction `json:"action"`
	Repository   *Repository    `json:"repository"`
	Organization *User          `json:"organization"`
	Sender       *User          `json:"sender"`
}

// JSONPayload JSON representation of the payload
func (p *RepositoryPayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", " ")
}
