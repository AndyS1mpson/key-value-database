package size_parser

import "errors"

var ErrIncorrectSize = errors.New("incorrect size")

// ParseSize parse size format (for example "4Kb")
func ParseSize(text string) (int, error) {
	if len(text) == 0 || text[0] < '0' || text[0] > '9' {
		return 0, ErrIncorrectSize
	}

	idx := 0
	size := 0
	for idx < len(text) && text[idx] >= '0' && text[idx] <= '9' {
		number := int(text[idx] - '0')
		size = size*10 + number
		idx++
	}

	parameter := text[idx:]
	switch parameter {
	case "GB", "Gb", "gb":
		return size << 30, nil
	case "MB", "Mb", "mb":
		return size << 20, nil
	case "KB", "Kb", "kb":
		return size << 10, nil
	case "B", "b", "":
		return size, nil
	default:
		return 0, ErrIncorrectSize
	}
}
