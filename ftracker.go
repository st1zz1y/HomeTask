package ftracker

import (
	"fmt"
	"math"
)

const (
	lenStep   = 0.65  // Средняя длина шага в метрах
	mInKm     = 1000  // Количество метров в километре
	minInH    = 60    // Количество минут в часе
	kmhInMsec = 0.278 // Коэффициент для преобразования км/ч в м/с

	// Множители для расчета калорий
	runningCaloriesMeanSpeedMultiplier = 18   // Множитель для бега
	runningCaloriesMeanSpeedShift      = 1.79 // Сдвиг для бега

	walkingCaloriesWeightMultiplier = 0.035 // Множитель для веса при ходьбе
	walkingSpeedHeightMultiplier    = 0.029 // Множитель для роста при ходьбе

	swimmingCaloriesMeanSpeedShift   = 1.1 // Сдвиг для плавания
	swimmingCaloriesWeightMultiplier = 2   // Множитель для веса при плавании
)

func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return distance(action) / duration
}

func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	var dist, speed, calories float64

	switch trainingType {
	case "Бег":
		dist = distance(action)
		speed = meanSpeed(action, duration)
		calories = RunningSpentCalories(action, weight, duration)
	case "Ходьба":
		dist = distance(action)
		speed = meanSpeed(action, duration)
		calories = WalkingSpentCalories(action, duration, weight, height)
	case "Плавание":
		dist = float64(lengthPool*countPool) / mInKm
		speed = swimmingMeanSpeed(lengthPool, countPool, duration)
		calories = SwimmingSpentCalories(lengthPool, countPool, duration, weight)
	default:
		return "неизвестный тип тренировки"
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
}

func RunningSpentCalories(action int, weight, duration float64) float64 {
	speed := meanSpeed(action, duration)
	return (runningCaloriesMeanSpeedMultiplier*speed + runningCaloriesMeanSpeedShift) * weight * duration
}

func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	speed := meanSpeed(action, duration)
	speedInMetersPerSecond := speed * kmhInMsec
	return (walkingCaloriesWeightMultiplier*weight +
		(math.Pow(speedInMetersPerSecond, 2)/height)*walkingSpeedHeightMultiplier*weight) * duration * minInH
}

func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool*countPool) / mInKm / duration
}

func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}
