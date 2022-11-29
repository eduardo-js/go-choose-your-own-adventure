package cyoa

type IStory map[string]IChapter

type IChapter struct {
	Title    string     `json:"title"`
	Scenario []string   `json:"scenario"`
	Options  []IOptions `json:"options"`
}

type IOptions struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}
