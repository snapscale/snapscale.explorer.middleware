package _type

type Details struct {
	Message    string `json:"message"`
	File       string `json:"file"`
	LineNumber int64  `json:"line_number"`
	Method     string `json:"method"`
}

type Error2 struct {
	Code    uint64    `json:"code"`
	Name    string    `json:"name"`
	What    string    `json:"what"`
	Details []Details `json:"details"`
}

type Error struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
	Error   Error2 `json:"error"`
}
