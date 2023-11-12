package structs

type ArtifactPayload struct {
	ArtifactID   int64  `json:"artifact_id"`
	ArtifactName string `json:"artifact_name"`
	RunID        int64  `json:"run_id"`
}
