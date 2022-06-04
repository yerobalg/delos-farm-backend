package domains

type Stats struct {
	APICount int64 `json:"api_count"`
	UniqueCallCount int64 `json:"unique_call_count"`
}