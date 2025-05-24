package mocks

import "time"

type TimerMock struct {
	mockedTime time.Time
}

func NewTimerMock(mockedTime time.Time) *TimerMock {
	return &TimerMock{mockedTime: mockedTime}
}

func (t *TimerMock) Now() time.Time {
	return t.mockedTime
}

func (t *TimerMock) Since(_ time.Time) time.Duration {
	return time.Millisecond
}
