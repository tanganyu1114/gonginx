package gonginx

//Location represents a location in nginx config
type Location struct {
	*Directive
	Modifier string
	Match    string
	Comments []string
}

//GetComment get the comment of directive
func (l *Location) GetComment() []string {
	return l.Comments
}
