package extractor

// The extractors interface extract metadata from requests and responses for validation and settlement.
// Any service that implements those interface can be registered and utilized in the 0g serving system.

type ProviderReqRespExtractor interface {
	GetInputCount(reqBody []byte) (int64, error)
	GetOutputCount(respBody [][]byte) (int64, error)
	GetRespContent(resp []byte, encodingType string) ([]byte, error)
	StreamCompleted(output []byte) (bool, error)
}