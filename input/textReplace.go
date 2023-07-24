package input

// Represents a Word input.
type TextReplace struct {
	Input
}

/** Represents the find and replace values and its options.
 */
func NewTextReplace(text string, replaceText string, matchCase bool) *TextReplace {
	var p TextReplace
	p.text = text
	p.replaceText = replaceText
	p.matchCase = matchCase
	return &p
}

/** Gets the Find Text value. This string will be replaced with <see cref="ReplaceText"/> during conversion. */
func (p *TextReplace) Text() string {
	return p.text
}

/** sets the Find Text value. This string will be replaced with <see cref="ReplaceText"/> during conversion. */
func (p *TextReplace) SetText(value string) {
	p.text = value
}

/** Gets or sets ReplaceText value.This string will replace the <see cref="Text"/> during conversion*/
func (p *TextReplace) ReplaceText() string {
	return p.replaceText
}

/**sets ReplaceText value.This string will replace the <see cref="Text"/> during conversion*/
func (p *TextReplace) SetReplaceText(value string) {
	p.replaceText = value
}

/** If True, the search operation will be case sensitive. */
func (p *TextReplace) MatchCase() bool {
	return p.matchCase
}
func (p *TextReplace) SetMatchCase(value bool) {
	p.matchCase = value
}
