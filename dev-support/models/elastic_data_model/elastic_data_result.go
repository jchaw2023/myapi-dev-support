package elastic_data_model

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}
type Total struct {
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}

type Hit struct {
	Index   string        `json:"_index"`
	Type    string        `json:"_type"`
	Id      string        `json:"_id"`
	Score   float64       `json:"_score"`
	Ignored []string      `json:"_ignored,omitempty"`
	Source  interface{}   `json:"_source"`
	Sort    []interface{} `json:"sort"`
}
type Hits struct {
	Total    *Total      `json:"total"`
	MaxScore interface{} `json:"max_score"`
	Hits     []*Hit      `json:"hits"`
}
type EsDataResult struct {
	Took     int     `json:"took"`
	TimedOut bool    `json:"timed_out"`
	Shards   *Shards `json:"_shards"`
	Hits     *Hits   `json:"hits"`
}
