package domains

type Stats struct {
	Path string `json:"path" gorm:"not null;column:path"`
	IP   string `json:"ip" gorm:"not null;column:ip"`
}

type StatsResults struct {
	Path            string `json:"path"`
	APICallCount    string `json:"api_call_count"`
	UniqueCallCount string `json:"unique_call_count"`
}

type StatsService interface {
	CreateStats(path string, ip string) error
	GetAllStats(limit string, offset string) ([]StatsResults, error)
}

type StatsRepository interface {
	CreateStats(stats Stats) error
	GetAllStats(limit int, offset int) ([]StatsResults, error)
}
