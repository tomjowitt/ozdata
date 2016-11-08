package ozdata

type Country struct {
	Name string
	Code string
}

type State struct {
	Name          string
	Code          string
	Capital       string
	Country       Country
	PostcodeRange []PostcodeRange `json:"PostcodeRange"`
}

type PostcodeRange struct {
	Low  int64
	High int64
}
