package packagex

func Chess(sizeY, sizeX int) string {
	str := ""
	for i := 0; i < sizeY; i++ {
		for j := 0; j < sizeX; j++ {
			if (i+j)%2 == 0 {
				str += " "
			} else {
				str += "#"
			}
		}
		str += "\n"
	}
	return str
}
