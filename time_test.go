package alarm

import (
	"fmt"
	"testing"
	"time"
)

func TestSetDefaultTime(t *testing.T) {
	test := Alarm{}
	t.Run("setDefaultTime test", func(t *testing.T) {
		test.SetDefaultTime(time.Date(2000, 1, 1, 12, 1, 1, 1, time.FixedZone("Asia/Tokyo", 9*60*60)))
		time_expected := time.Date(2000, 1, 1, 12, 1, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
		if test.time.Unix() != time_expected.Unix() {
			str := "defaultTime\n"
			str += fmt.Sprintf("%d\n", test.time.Unix())
			str += fmt.Sprintf("%d\n", time_expected.Unix())
			t.Errorf(str)
		}
	})
}

func TestTimeIncrement(t *testing.T) {
	tests := []Alarm{
		{
			time:         time.Date(2000, 1, 1, 12, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
			ringing:      false,
			selectorTime: 0,
		},
		{
			time:         time.Date(2000, 1, 1, 12, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
			ringing:      true,
			selectorTime: 1,
		},
	}
	for i, test := range tests {
		if i == 0 {
			t.Run("minute increment test", func(t *testing.T) {
				expected := time.Date(2000, 1, 1, 12, 1, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
				test.TimeIncrement()
				if test.time.Equal(expected) == false {
					str := "timeIncrement(minute)\n"
					str += fmt.Sprintf("%d\n", test.time.Unix())
					str += fmt.Sprintf("%d\n", expected.Unix())
					t.Errorf(str)
				}
			})
		} else if i == 1 {
			t.Run("hour increment test", func(t *testing.T) {
				expected := time.Date(2000, 1, 1, 13, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
				test.TimeIncrement()
				if test.time.Equal(expected) == false {
					str := "timeIncrement(hour)\n"
					str += fmt.Sprintf("%d\n", test.time.Unix())
					str += fmt.Sprintf("%d\n", expected.Unix())
					t.Errorf(str)
				}
			})
		}
	}
}

func TestTimeDecrement(t *testing.T) {
	tests := []Alarm{
		{
			time:         time.Date(2000, 1, 1, 12, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
			ringing:      false,
			selectorTime: 0,
		},
		{
			time:         time.Date(2000, 1, 1, 12, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
			ringing:      true,
			selectorTime: 1,
		},
	}
	for i, test := range tests {
		if i == 0 {
			t.Run("minute Decrement test", func(t *testing.T) {
				expected := time.Date(2000, 1, 1, 11, 59, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
				test.TimeDecrement()
				if test.time.Equal(expected) == false {
					str := "timeDecrement(minute)\n"
					str += fmt.Sprintf("%d\n", test.time.Unix())
					str += fmt.Sprintf("%d\n", expected.Unix())
					t.Errorf(str)
				}
			})
		} else if i == 1 {
			t.Run("hour Decrement test", func(t *testing.T) {
				expected := time.Date(2000, 1, 1, 11, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
				test.TimeDecrement()
				if test.time.Equal(expected) == false {
					str := "timeDecrement(hour)\n"
					str += fmt.Sprintf("%d\n", test.time.Unix())
					str += fmt.Sprintf("%d\n", expected.Unix())
					t.Errorf(str)
				}
			})
		}
	}
}

func TestAlarmOn(t *testing.T) {
	test := Alarm{
		time:         time.Date(2000, 1, 1, 12, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
		ringing:      false,
		selectorTime: 0,
	}
	t.Run("alarmOn test", func(t *testing.T) {
		alarmRinging_expected := true
		func_expected := 1
		func_testdata := 0
		test.AlarmOnIfTimeMatched(time.Date(2000, 1, 1, 12, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)), func() { func_testdata = 1 })
		if test.ringing != alarmRinging_expected {
			str := "alarmRinging test\n"
			str += fmt.Sprintf("%T\n", test.ringing)
			str += fmt.Sprintf("%T\n", alarmRinging_expected)
			t.Errorf(str)
		}
		if func_testdata != func_expected {
			str := "alarmfunc test\n"
			str += fmt.Sprintf("%d\n", func_testdata)
			str += fmt.Sprintf("%d\n", func_expected)
			t.Errorf(str)
		}
	})
}

func TestAlarmOff(t *testing.T) {
	test := Alarm{
		time:         time.Date(2000, 1, 1, 12, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
		ringing:      true,
		selectorTime: 0,
	}
	t.Run("alarmOff test", func(t *testing.T) {
		alarmRinging_expected := false
		func_expected := 1
		func_testdata := 0
		test.AlarmOff(func() { func_testdata = 1 })
		if test.ringing != alarmRinging_expected {
			str := "alarmRinging test\n"
			str += fmt.Sprintf("%T\n", test.ringing)
			str += fmt.Sprintf("%T\n", alarmRinging_expected)
			t.Errorf(str)
		}
		if func_testdata != func_expected {
			str := "alarmfunc test\n"
			str += fmt.Sprintf("%d\n", func_testdata)
			str += fmt.Sprintf("%d\n", func_expected)
			t.Errorf(str)
		}
	})
}

func TestAdjustDay(t *testing.T) {
	test := Alarm{
		time: time.Date(2000, 1, 1, 12, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60)),
	}
	t.Run("timeAdjustDay test", func(t *testing.T) {
		test.AdjustDay(time.Date(2001, 2, 2, 13, 1, 1, 1, time.FixedZone("Asia/Tokyo", 9*60*60)))
		time_expected := time.Date(2001, 2, 2, 12, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
		if test.time.Unix() != time_expected.Unix() {
			str := "AdjustDay test\n"
			str += fmt.Sprintf("%d\n", test.time.Unix())
			str += fmt.Sprintf("%d\n", time_expected.Unix())
			t.Errorf(str)
		}
	})
}
