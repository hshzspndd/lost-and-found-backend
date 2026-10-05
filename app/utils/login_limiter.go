package utils

import (
	"sync"
	"time"
)

const (
	maxLoginFailCount = 5                // 连续失败次数阈值
	loginLockDuration = 15 * time.Minute // 锁定时长
	failCount         = 30 * time.Minute // 失败计数重置周期
)

type loginAttempt struct {
	failCount   int
	lastFailAt  time.Time
	lockedUntil time.Time
}

var (
	loginMu       sync.Mutex
	loginAttempts = make(map[string]*loginAttempt)
)

// 检查是否被锁定
func IsLoginLocked(key string) bool {
	loginMu.Lock()
	defer loginMu.Unlock()

	a, ok := loginAttempts[key]
	if !ok {
		return false
	}
	if time.Now().Before(a.lockedUntil) {
		return true
	}
	delete(loginAttempts, key)
	return false
}

// 记录一次登录失败
func RecordLoginFail(key string) {
	loginMu.Lock()
	defer loginMu.Unlock()

	now := time.Now()
	a, ok := loginAttempts[key]
	if !ok {
		a = &loginAttempt{}
		loginAttempts[key] = a
	}
	// 重置计数
	if now.Sub(a.lastFailAt) > failCount {
		a.failCount = 0
	}
	a.failCount += 1
	a.lastFailAt = now
	if a.failCount >= maxLoginFailCount {
		a.lockedUntil = now.Add(loginLockDuration)
	}
}

// 登录成功后清零失败计数
func ResetLoginFail(key string) {
	loginMu.Lock()
	defer loginMu.Unlock()
	delete(loginAttempts, key)
}
