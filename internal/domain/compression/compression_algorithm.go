package compression

type CompressionAlgorithm string

const (
	CompressionAlgZSTD CompressionAlgorithm = "zstd"
	CompressionAlgGZip CompressionAlgorithm = "gzip"
	CompressionAlgZlib CompressionAlgorithm = "zlib"
	CompressionAlgZip  CompressionAlgorithm = "zip"
	CompressionAlgXZ   CompressionAlgorithm = "xz"
)
