package session

type Document struct {
	content    string
	operations []int
	language   string
}

func CreateEmptyDoc() *Document {
	newDoc := &Document{
		content:    "",
		operations: make([]int, 0),
		language:   "golang",
	}

	return newDoc
}
