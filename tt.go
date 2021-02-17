package tt

// TextTicker for flowing text left and right within a fixed length
type TextTicker struct {
	message      []rune
	fixedLength  uint
	direction    direction
	currentIndex uint
}

// flow direction
type direction uint8

const (
	left  direction = iota
	right direction = iota
)

// NewTextTicker creates a new TextTicker
func NewTextTicker(message string, fixedLength uint) *TextTicker {
	return &TextTicker{
		message:      []rune(message),
		fixedLength:  fixedLength,
		direction:    left,
		currentIndex: 0,
	}
}

// NextText returns current ticker text
func (t *TextTicker) NextText() string {
	length := uint(len(t.message))

	if length <= t.fixedLength {
		return string(t.message)
	}

	text := string(t.message[t.currentIndex : t.currentIndex+t.fixedLength])

	if t.direction == right {
		if t.currentIndex > 0 {
			t.currentIndex--
		} else {
			t.currentIndex++
			t.direction = left
		}
	} else if t.direction == left {
		if t.currentIndex+t.fixedLength >= length {
			t.currentIndex--
			t.direction = right
		} else {
			t.currentIndex++
		}
	}

	return text
}
