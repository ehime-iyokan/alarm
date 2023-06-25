package alarm

import "time"

type Alarm struct {
	time         time.Time // 時間は [ns] 単位では比較しないため、nsec には 0 を設定する
	ringing      bool
	selectorTime int // alarm.selectorTime = 0:秒調整, 1:時間調整
}

// デフォルトのアラーム時刻を設定する。「？時：？分：0 秒」でアラームを鳴らすかを判定するため、 second は 0 を設定する
func (a *Alarm) SetDefaultTime(t time.Time) {
	a.time = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

// selectorTime が 0 なら "分"、1 なら "秒"を +1 する
func (a *Alarm) TimeIncrement() {
	if a.selectorTime == 0 {
		minuteIncrementer, _ := time.ParseDuration("1m")
		a.time = a.time.Add(minuteIncrementer)
	} else {
		hourIncrementer, _ := time.ParseDuration("1h")
		a.time = a.time.Add(hourIncrementer)
	}
}

// selectorTime が 0 なら "分"、1 なら "秒"を -1 する
func (a *Alarm) TimeDecrement() {
	if a.selectorTime == 0 {
		minuteDecrementer, _ := time.ParseDuration("-1m")
		a.time = a.time.Add(minuteDecrementer)
	} else {
		hourDecrementer, _ := time.ParseDuration("-1h")
		a.time = a.time.Add(hourDecrementer)
	}
}

// 今の時間とアラームの時間が一致しているかを確認し、一致している場合は funcAlarmOn を実行する
func (a *Alarm) AlarmOnIfTimeMatched(t_now time.Time, funcAlarmOn func()) {
	t_now = time.Date(t_now.Year(), t_now.Month(), t_now.Day(), t_now.Hour(), t_now.Minute(), t_now.Second(), 0, t_now.Location())
	if t_now.Equal(a.time) {
		a.ringing = true
		funcAlarmOn()
	}
}

// アラームをOFFにし、funcAlarmOff を実行する
func (a *Alarm) AlarmOff(funcAlarmOff func()) {
	a.ringing = false
	funcAlarmOff()
}

// 日付が変わっても時間が一致するようにするため、年月日を今の時間に同期させる。
func (a *Alarm) AdjustDay(t_now time.Time) {
	a.time = time.Date(t_now.Year(), t_now.Month(), t_now.Day(), a.time.Hour(), a.time.Minute(), a.time.Second(), 0, t_now.Location())
}

// selectorTime に値を設定する
func (a *Alarm) SetSelectorTime(value int) {
	a.selectorTime = value
}

// Alarm構造体のメンバ変数にアクセスするためのメソッド
func (a *Alarm) GetStatusRinging() bool {
	return a.ringing
}
func (a *Alarm) GetStatusSelectorTime() int {
	return a.selectorTime
}
func (a *Alarm) GetTime() time.Time {
	return a.time
}
