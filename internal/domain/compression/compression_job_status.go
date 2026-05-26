package compression

type CompressionJobStatus string

const (
	CompressionJobStatusPending    CompressionJobStatus = "pending"
	CompressionJobStatusInProgress CompressionJobStatus = "in_progress"
	CompressionJobStatusSucceeded  CompressionJobStatus = "succeeded"
	CompressionJobStatusSuccess    CompressionJobStatus = "success" // todo: redundant? remove?
	CompressionJobStatusFailed     CompressionJobStatus = "failed"
	CompressionJobStatusCanceled   CompressionJobStatus = "canceled"
)
