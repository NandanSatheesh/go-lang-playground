package models

import "encoding/json"

func UnmarshalNamesResponse(data []byte) (NamesResponse, error) {
	var r NamesResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NamesResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NamesResponse struct {
	Data    Data   `json:"data"`
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

type Data struct {
	Names []string `json:"names"`
}

type Score struct {
	Name  string
	Value float32
}

type ScoreMap []Score

func (s ScoreMap) Len() int {
	return len(s)
}

func (s ScoreMap) Less(i int, j int) bool {
	return s[i].Value < s[j].Value
}

func (s ScoreMap) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}
