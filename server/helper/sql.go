package helper

import (
	"gopkg.in/guregu/null.v4"
)

func StringToNullTime(field string) null.Time {
	if field == "" {
		return null.NewTime(GetCurrentDatetime(), false)
	} else {
		t := FromDatetimeFormat(field)
		return null.TimeFrom(t)
	}

}
