package data

func (link *Link) ShortTitle() (shortTitle string) {
	if len(link.Title) > 20 {
		shortTitle = string([]rune(link.Title)[:20])
	} else {
		shortTitle = link.Title
	}
	return
}
