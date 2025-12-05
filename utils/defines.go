package utils

type Item struct {
	TitleValue,
	Desc string
}

func (i Item) Title() string       { return i.TitleValue }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.TitleValue }
