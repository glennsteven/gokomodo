package presentations

type (
	Paging struct {
		Limit int `url:"limit,omitempty" db:"limit,omitempty"`
		Page  int `url:"page,omitempty" db:"page,omitempty"`
	}
)

type MetaData struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPage  int `json:"total_page"`
	TotalCount int `json:"total_count"`
}
