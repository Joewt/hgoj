package tools

import "strconv"

func StringToInt32(s string) (int32, error) {
	id , err := strconv.ParseInt(s,10,32)
	id64 := int32(id)
	if err != nil {
		return -1, err
	}
	return id64, nil
}

