package eisoduration

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"text/template"
	"time"
)

var (
	ErrBadFormat = errors.New("bad format string")

	tmpl = template.Must(template.New("duration").Parse(`P{{if .Years}}{{.Years}}Y{{end}}{{if .Months}}{{.Months}}M{{end}}{{if .Weeks}}{{.Weeks}}W{{end}}{{if .Days}}{{.Days}}D{{end}}{{if .HasTimePart}}T{{end }}{{if .Hours}}{{.Hours}}H{{end}}{{if .Minutes}}{{.Minutes}}M{{end}}{{if .Seconds}}{{.Seconds}}S{{end}}`))

	full = regexp.MustCompile(`^P(?:(?P<year>\d+)Y)?(?:(?P<month>\d+)M)?(?:(?P<week>\d+)W)?(?:(?P<day>\d+)D)?(?:T(?:(?P<hour>\d+)H|(?P<minute>\d+)M|(?P<second>\d+)S)+)?$`)
)

type Duration struct {
	Years   int
	Months  int
	Weeks   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func FromString(dur string) (*Duration, error) {
	if !full.MatchString(dur) {
		return nil, ErrBadFormat
	}

	match := full.FindStringSubmatch(dur)
	d := &Duration{}

	for i, name := range full.SubexpNames() {
		part := match[i]
		if i == 0 || name == "" || part == "" {
			continue
		}

		val, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		switch name {
		case "year":
			d.Years = val
		case "month":
			d.Months = val
		case "week":
			d.Weeks = val
		case "day":
			d.Days = val
		case "hour":
			d.Hours = val
		case "minute":
			d.Minutes = val
		case "second":
			d.Seconds = val
		default:
			return nil, errors.New(fmt.Sprintf("unknown field %s", name))
		}
	}

	return d, nil
}

// String打印出传入的值。它并不严格按照
// ISO规范，但非常接近。特别是要完全符合它
// 需要四舍五入到下一个最大的单位。61秒到1分1
// 第二，例如。它还需要禁止几周的时间与
// 其他单位。
func (d *Duration) String() string {
	var s bytes.Buffer

	err := tmpl.Execute(&s, d)
	if err != nil {
		panic(err)
	}

	return s.String()
}

func (d *Duration) HasTimePart() bool {
	return d.Hours != 0 || d.Minutes != 0 || d.Seconds != 0
}

func (d *Duration) ToDuration() time.Duration {
	now := time.Now()

	// 计算年月对应的真实时间
	future := now.AddDate(d.Years, d.Months, 0)

	// 计算基于年月的持续时间
	yearMonthDuration := future.Sub(now)

	// 计算其他固定单位的持续时间
	fixedDuration := time.Duration(0)
	day := time.Hour * 24
	fixedDuration += day * 7 * time.Duration(d.Weeks)
	fixedDuration += day * time.Duration(d.Days)
	fixedDuration += time.Hour * time.Duration(d.Hours)
	fixedDuration += time.Minute * time.Duration(d.Minutes)
	fixedDuration += time.Second * time.Duration(d.Seconds)

	return yearMonthDuration + fixedDuration
}

func (d *Duration) ToDurationFrom(start time.Time) time.Duration {
	future := start.AddDate(d.Years, d.Months, 0)
	yearMonthDuration := future.Sub(start)

	fixedDuration := time.Duration(0)
	day := time.Hour * 24
	fixedDuration += day * 7 * time.Duration(d.Weeks)
	fixedDuration += day * time.Duration(d.Days)
	fixedDuration += time.Hour * time.Duration(d.Hours)
	fixedDuration += time.Minute * time.Duration(d.Minutes)
	fixedDuration += time.Second * time.Duration(d.Seconds)

	return yearMonthDuration + fixedDuration
}
