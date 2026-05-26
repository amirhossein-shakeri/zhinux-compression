package compression

import (
	"time"

	"github.com/amirhossein-shakeri/zhinux-platform/types"
)

type CompressionArtifact struct {
	ID       types.ID // Internal ID(Fast joins)
	PublicID string   // Exposed UUID at public APIs(Safe public identifiers)

	CompressionJobID types.ID

	// Storage StorageLocation
	Size int

	Metadata map[string]any

	CreatedAt time.Time
	DeletedAt *time.Time
}
