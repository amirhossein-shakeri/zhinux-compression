package compression

import (
	"time"

	"github.com/amirhossein-shakeri/zhinux-platform/types"
)

type CompressionJob struct {
	ID       types.ID // Internal ID(Fast joins)
	PublicID string   // Exposed UUID at public APIs(Safe public identifiers)

	Status     CompressionJobStatus
	StartedAt  *time.Time
	FinishedAt *time.Time
	Duration   *time.Duration

	InputSize  *int // Uncompressed bytes size
	OutputSize *int // Compressed bytes size

	CompressionPercentage *uint

	InputArtifactID  *types.ID // If uploaded/retried/queued/etc
	OutputArtifactID *types.ID // If uploaded/retried/queued/etc

	CompressionAlgorithm CompressionAlgorithm
	// CompressionOptions | CompressionArgs

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
