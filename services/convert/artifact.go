package convert

import (
	"context"

	models_actions "code.gitea.io/gitea/models/actions"
	api "code.gitea.io/gitea/modules/structs"
)

// ToRepo converts an ActionRun to api.ActionRun
func ToArtifact(ctx context.Context, artifact *models_actions.ActionArtifact) *api.ArtifactPayload {
	return &api.ArtifactPayload{
		ArtifactID:   artifact.ID,
		ArtifactName: artifact.ArtifactName,
		RunID:        artifact.RunID,
	}
}
