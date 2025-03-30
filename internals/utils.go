package internals

func isColliding(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return x1 < x2+w2 &&
		x1+w1 > x2 &&
		y1 < y2+h2 &&
		y1+h1 > y2
}

func getTextWidthAndHeight(text string) (int,int){
	textWidth := len(text) * 7 // approx. 7 pixels per character
	textHeight := 13

	return textWidth, textHeight
}