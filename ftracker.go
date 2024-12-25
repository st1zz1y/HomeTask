package ftracker

import (
	"fmt"
	"math"
)

const (
	lenStep   = 0.65  // Средняя длина шага.
	mInKm     = 1000  // Количество метров в километре.
	minInH    = 60    // Количество минут в часе.
	kmhInMsec = 0.278 // Коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // Количество сантиметров в метре.
)

// distance возвращает дистанцию (в километрах), которую преодолел пользователь за время тренировки.
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	dist := distance(action)
	return dist / duration
}

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool*countPool) / mInKm / duration
}

// RunningSpentCalories возвращает количество потраченных калорий при беге.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	speed := meanSpeed(action, duration)
	return ((runningCaloriesMeanSpeedMultiplier * speed * runningCaloriesMeanSpeedShift) * weight / mInKm) * duration * minInH
}

const (
	runningCaloriesMeanSpeedMultiplier = 18.0  // Множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 1.79 // Сдвиг средней скорости.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	heightInMeters := height / cmInM
	speed := meanSpeed(action, duration)
	speedInMetersPerSecond := speed * kmhInMsec
	return ((walkingCaloriesWeightMultiplier * weight + (math.Pow(speedInMetersPerSecond, 2)/heightInMeters)*walkingSpeedHeightMultiplier*weight) * duration * minInH)
}

const (
	walkingCaloriesWeightMultiplier = 0.035 // Множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // Множитель роста.
)

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}

const (
	swimmingCaloriesMeanSpeedShift   = 1.1 // Среднее количество сжигаемых калорий при плавании.
	swimmingCaloriesWeightMultiplier = 2.0 // Множитель веса при плавании.
)

// ShowTrainingInfo возвращает строку с информацией о тренировке.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	var dist, speed, calories float64
	switch trainingType {
	case "Бег":
		dist = distance(action)
		speed = meanSpeed(action, duration)
		calories = RunningSpentCalories(action, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
	case "Ходьба":
		dist = distance(action)
		speed = meanSpeed(action, duration)
		calories = WalkingSpentCalories(action, duration, weight, height)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
	case "Плавание":
		speed = swimmingMeanSpeed(lengthPool, countPool, duration)
		dist = float64(lengthPool*countPool) / mInKm
		calories = SwimmingSpentCalories(lengthPool, countPool, duration, weight)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
	default:
		return "Неизвестный тип тренировки"
	}
}
