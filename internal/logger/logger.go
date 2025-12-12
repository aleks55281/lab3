package logger

import (
	"fmt"
	"sync"
)

// GameLogger - одиночка
type GameLogger struct {
	logs []string
	mu   sync.RWMutex
}

var (
	instance *GameLogger
	once     sync.Once
)

// GetInstance - метод для получения одиночки
func GetInstance() *GameLogger {
	once.Do(func() {
		instance = &GameLogger{
			logs: make([]string, 0),
		}
	})
	return instance
}

// Log - поведение одиночки
func (gl *GameLogger) Log(message string) {
	gl.mu.Lock()
	defer gl.mu.Unlock()
	logMessage := fmt.Sprintf("[GAME LOG]: %s", message)
	gl.logs = append(gl.logs, logMessage)
	fmt.Println(logMessage)
}

func (gl *GameLogger) GetLogs() []string {
	gl.mu.RLock()
	defer gl.mu.RUnlock()
	return gl.logs
}
