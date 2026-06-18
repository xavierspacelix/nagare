package gestures

import (
	"time"

	"nagare/models"
)

type CooldownManager struct {
	cooldowns map[models.Gesture]time.Time
	config    map[models.Gesture]time.Duration
	clock     func() time.Time
}

func NewCooldownManager(defaultDuration time.Duration) *CooldownManager {
	cm := &CooldownManager{
		cooldowns: make(map[models.Gesture]time.Time),
		config:    make(map[models.Gesture]time.Duration),
		clock:     time.Now,
	}

	cm.config[models.GesturePinch] = defaultDuration
	cm.config[models.GesturePinchHold] = defaultDuration
	cm.config[models.GestureTwoFingerPinch] = defaultDuration
	cm.config[models.GestureSwipeLeft] = defaultDuration
	cm.config[models.GestureSwipeRight] = defaultDuration
	cm.config[models.GestureTwoFingerUp] = 500 * time.Millisecond
	cm.config[models.GestureTwoFingerDown] = 500 * time.Millisecond

	return cm
}

func (cm *CooldownManager) IsReady(gesture models.Gesture) bool {
	expiry, exists := cm.cooldowns[gesture]
	if !exists {
		return true
	}
	return cm.clock().After(expiry)
}

func (cm *CooldownManager) Start(gesture models.Gesture) {
	duration, exists := cm.config[gesture]
	if !exists {
		duration = 250 * time.Millisecond
	}
	cm.cooldowns[gesture] = cm.clock().Add(duration)
}

func (cm *CooldownManager) Reset(gesture models.Gesture) {
	delete(cm.cooldowns, gesture)
}

func (cm *CooldownManager) SetCooldown(gesture models.Gesture, duration time.Duration) {
	cm.config[gesture] = duration
}

func (cm *CooldownManager) Remaining(gesture models.Gesture) time.Duration {
	expiry, exists := cm.cooldowns[gesture]
	if !exists {
		return 0
	}
	remaining := expiry.Sub(cm.clock())
	if remaining < 0 {
		return 0
	}
	return remaining
}
