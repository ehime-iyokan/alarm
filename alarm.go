package alarm

import "time"

type Alarm struct {
	time         time.Time
	ringing      bool
	selectorTime int // alarm.selectorTime = 0:秒調整, 1:時間調整
}

// デフォルトのアラーム時刻を設定する。second, ns は 0 が設定される
func (a *Alarm) setDefaultTime(t time.Time) {
	a.time = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

// selectorTime が 0 なら "分"、1 なら "秒"を +1 する
func (a *Alarm) timeIncrement() {
	if a.selectorTime == 0 {
		minuteIncrementer, _ := time.ParseDuration("1m")
		a.time = a.time.Add(minuteIncrementer)
	} else {
		hourIncrementer, _ := time.ParseDuration("1h")
		a.time = a.time.Add(hourIncrementer)
	}
}

// selectorTime が 0 なら "分"、1 なら "秒"を -1 する
func (a *Alarm) timeDecrement() {
	if a.selectorTime == 0 {
		minuteDecrementer, _ := time.ParseDuration("-1m")
		a.time = a.time.Add(minuteDecrementer)
	} else {
		hourDecrementer, _ := time.ParseDuration("-1h")
		a.time = a.time.Add(hourDecrementer)
	}
}

// 今の時間とアラームの時間が一致しているかを確認し、一致している場合は funcAlarmOn を実行する
func (a *Alarm) alarmOnIfTimeMatched(t_now time.Time, funcAlarmOn func()) {
	t_now = time.Date(t_now.Year(), t_now.Month(), t_now.Day(), t_now.Hour(), t_now.Minute(), t_now.Second(), 0, t_now.Location())
	if t_now.Equal(a.time) {
		a.ringing = true
		funcAlarmOn()
	}
}

// アラームをOFFにし、funcAlarmOff を実行する
func (a *Alarm) alarmOff(funcAlarmOff func()) {
	a.ringing = false
	funcAlarmOff()
}

// 日付が変わっても時間が一致するようにするため、年月日を今の時間に同期させる。
func (a *Alarm) adjustDay(t_now time.Time) {
	a.time = time.Date(t_now.Year(), t_now.Month(), t_now.Day(), a.time.Hour(), a.time.Minute(), a.time.Second(), 0, t_now.Location())
}
