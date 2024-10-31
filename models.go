package main

type BinRequest struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

type Bin struct {
	Id          string `json:"id"`
	Text        string `json:"text"`
	Title       string `json:"title"`
	Timestamp   string `json:"timestamp"`
	SeenCounter int64  `json:"seen"`
	StarCounter int64  `json:"stars"`
}
