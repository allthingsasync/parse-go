package text

type SourceText struct {
	text             []rune
	length           int
	current_location int
	lexeme_start     int
}

func NewSource(text string) SourceText {
    t := []rune(text)
	return SourceText{
		text:             t,
		length:           len(t),
		lexeme_start:     0,
		current_location: 0,
	}
}

func (this *SourceText) Length() int {
    return this.length
}

func (this *SourceText) Peek(offset int) rune {
	if this.current_location+offset >= this.length {
		return -1
	}
	return this.text[this.current_location+offset]
}

func (this *SourceText) Current() rune {
	return this.Peek(0)
}

func (this *SourceText) LA() rune {
	return this.Peek(1)
}

func (this *SourceText) IsEnd() bool {
	return this.current_location >= this.length
}

func (this *SourceText) Advance(count int) {
	this.current_location += count
}

func (this *SourceText) Eat() rune {
	r := this.Current()
	this.Advance(1)
	return r
}
