package convert

import (
	"context"

	models_actions "code.gitea.io/gitea/models/actions"
	"code.gitea.io/gitea/models/perm"
	access_model "code.gitea.io/gitea/models/perm/access"
	api "code.gitea.io/gitea/modules/structs"
)

// ToRepo converts an ActionRun to api.ActionRun
func ToActionRun(ctx context.Context, run *models_actions.ActionRun) *api.ActionRun {
	return &api.ActionRun{
		ID:          run.ID,
		Title:       run.Title,
		Repo:        ToRepo(ctx, run.Repo, access_model.Permission{AccessMode: perm.AccessModeOwner}),
		Owner:       ToUser(ctx, run.Repo.Owner, nil),
		TriggerUser: ToUser(ctx, run.TriggerUser, nil),
		WorkflowID:  run.WorkflowID,
		Index:       run.Index,
		Ref:         run.Ref,
		CommitSHA:   run.CommitSHA,
		Status:      run.Status.String(),
		Version:     run.Version,
		Started:     run.Started.FormatLong(),
		Stopped:     run.Stopped.FormatLong(),
		Created:     run.Created.FormatLong(),
		Updated:     run.Updated.FormatLong(),
	}
}
