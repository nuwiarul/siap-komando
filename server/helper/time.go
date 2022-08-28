package helper

import (
	"time"

	"github.com/jinzhu/now"
	"github.com/vigneshuvi/GoDateFormat"
)

func GetCurrentDate() string {
	t := timeIn(time.Now(), "Asia/Makassar")
	return t.Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
}

func ToIntStringShortMonth(bulan int) string {
	month := "Jan"

	if bulan == 1 {
		month = "Jan"
	} else if bulan == 2 {
		month = "Feb"
	} else if bulan == 3 {
		month = "Mar"
	} else if bulan == 4 {
		month = "Apr"
	} else if bulan == 5 {
		month = "Mei"
	} else if bulan == 6 {
		month = "Jun"
	} else if bulan == 7 {
		month = "Jul"
	} else if bulan == 8 {
		month = "Aug"
	} else if bulan == 9 {
		month = "Sep"
	} else if bulan == 10 {
		month = "Okt"
	} else if bulan == 11 {
		month = "Nop"
	} else if bulan == 12 {
		month = "Des"
	}

	return month
}

func ToIntStringMonth(bulan int) string {
	month := "Januari"

	if bulan == 1 {
		month = "Januari"
	} else if bulan == 2 {
		month = "Februari"
	} else if bulan == 3 {
		month = "Maret"
	} else if bulan == 4 {
		month = "April"
	} else if bulan == 5 {
		month = "Mei"
	} else if bulan == 6 {
		month = "Juni"
	} else if bulan == 7 {
		month = "Juli"
	} else if bulan == 8 {
		month = "Agustus"
	} else if bulan == 9 {
		month = "September"
	} else if bulan == 10 {
		month = "Oktober"
	} else if bulan == 11 {
		month = "Nopember"
	} else if bulan == 12 {
		month = "Desember"
	}

	return month
}

func ToDateFormat(waktu time.Time) string {
	return waktu.Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
}

func ToDatetimeFormat(waktu time.Time) string {
	return waktu.Format(GoDateFormat.ConvertFormat("yyyy-mm-dd HH:MM:ss"))
}

func ToTimeFormat(waktu time.Time) string {
	return waktu.Format(GoDateFormat.ConvertFormat("HH:MM:ss"))
}

func FromTimeFormat(waktu string) time.Time {
	t, err := time.Parse(GoDateFormat.ConvertFormat("HH:MM:ss"), waktu)
	PanicIfError(err)
	return t

}

func FromDateFormat(tanggal string) time.Time {
	t, err := time.Parse(GoDateFormat.ConvertFormat("yyyy-mm-dd"), tanggal)
	PanicIfError(err)
	return t

}

func FromDatetimeFormat(datetime string) time.Time {
	t, err := time.Parse(GoDateFormat.ConvertFormat("yyyy-mm-dd HH:MM:ss"), datetime)
	PanicIfError(err)
	return t

}

func EndOfMonth(bulan int, tahun int) (string, string) {
	loc := timeIn(time.Now(), "Asia/Makassar")
	t := time.Date(tahun, time.Month(bulan), 18, 17, 51, 49, 123456789, loc.Location())
	end := now.With(t).EndOfMonth()
	start := now.With(t).BeginningOfMonth()
	endDay := end.Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
	startDay := start.Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
	return startDay, endDay
}

func EndOfYear(bulan int, tahun int) (string, string) {
	loc := timeIn(time.Now(), "Asia/Makassar")
	t := time.Date(tahun, time.Month(bulan), 18, 17, 51, 49, 123456789, loc.Location())
	end := now.With(t).EndOfYear()
	start := now.With(t).BeginningOfYear()
	endDay := end.Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
	startDay := start.Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
	return startDay, endDay
}

func GetCurrentDatetimeString() string {
	t := timeIn(time.Now(), "Asia/Makassar")
	return t.Format(GoDateFormat.ConvertFormat("yyyy-mm-dd HH:MM:ss"))
}

func GetCurrentDatetime() time.Time {
	waktu := GetCurrentDatetimeString()
	return FromDatetimeFormat(waktu)
}

func timeIn(t time.Time, name string) time.Time {
	loc, err := time.LoadLocation(name)
	PanicIfError(err)
	return t.In(loc)
}
