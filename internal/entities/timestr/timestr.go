package timestr

import (
	"fmt"
	"time"
)

const (
	layoutMin = "2006/01/02 15:04"
	layoutDay = "2006/01/02"
)

func Validate(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	_, errM := time.ParseInLocation(layoutMin, str, time.Local)
	_, errD := time.ParseInLocation(layoutDay, str, time.Local)

	if errM != nil && errD != nil {
		return "", fmt.Errorf("invalid layout: [%s, %s]", errM.Error(), errD.Error())
	} else if errM == nil && errD != nil {
		return str, nil
	} else if errM != nil && errD == nil {
		return str + " 00:00", nil
	}

	return "", fmt.Errorf("Parse failed for some reason")
}

func Parse(str string) (*time.Time, error) {
	tM, errM := time.ParseInLocation(layoutMin, str, time.Local)
	tD, errD := time.ParseInLocation(layoutDay, str, time.Local)

	if errM != nil && errD != nil {
		return nil, fmt.Errorf("invalid layout: [%s, %s]", errM.Error(), errD.Error())
	} else if errM == nil && errD != nil {
		return &tM, nil
	} else if errM != nil && errD == nil {
		return &tD, nil
	}

	return nil, fmt.Errorf("Parse failed for some reason")
}
