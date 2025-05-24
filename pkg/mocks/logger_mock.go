package mocks

import (
	"fmt"
	"slices"
	"sync"
)

type LoggerMock struct {
	sync.RWMutex
	lines []string
}

func (l *LoggerMock) Fatal(v ...any) {
	l.Lock()
	defer l.Unlock()

	l.lines = append(l.lines, fmt.Sprintf("%s", v...))
}

func (l *LoggerMock) Printf(format string, v ...any) {
	l.Lock()
	defer l.Unlock()

	l.lines = append(l.lines, fmt.Sprintf(format, v...))
}

func (l *LoggerMock) Println(v ...any) {
	l.Lock()
	defer l.Unlock()

	l.lines = append(l.lines, fmt.Sprintf("%s", v...))
}

func (l *LoggerMock) HasLogged(line string) bool {
	l.RLock()
	defer l.RUnlock()

	return slices.Contains(l.lines, line)
}
