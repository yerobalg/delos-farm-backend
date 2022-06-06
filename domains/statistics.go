package domains

type Stats struct {
	APICount int64 `json:"api_count"`
	UniqueCallCount int64 `json:"unique_call_count"`
}

type StatsService interface {
	GetStatistics(path string, ip string) (Stats)
}

type StatsRepository interface{
	CountAPICall(path string) (int64, error)
	CountUniqueCall(ip string) (int64, error)
}