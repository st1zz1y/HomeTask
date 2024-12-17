package ftracker

import (
	"fmt"
)

const (
	lenStep   = 0.65  // средняя длина шага
	mInKm     = 1000  // количество метров в километре
	minInH    = 60    // количество минут в часе
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с
	cmInM     = 100   // количество сантиметров в метре

	runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости
	runningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге

	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста

	swimmingCaloriesMeanSpeedShift   = 1.1 // среднее количество сжигаемых калорий при плавании относительно скорости
	swimmingCaloriesWeightMultiplier = 2   // множитель веса при плавании
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
	switch trainingType {
	case "Бег":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := RunningSpentCalories(action, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case "Ходьба":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := WalkingSpentCalories(action, duration, weight, height)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case "Плавание":
		distance := float64(lengthPool*countPool) / mInKm
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

func RunningSpentCalories(action int, weight, duration float64) float64 {
	speed := meanSpeed(action, duration)
	return (speed*runningCaloriesMeanSpeedMultiplier + runningCaloriesMeanSpeedShift) * weight * duration
}

func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	speed := meanSpeed(action, duration)
	return (walkingCaloriesWeightMultiplier + walkingSpeedHeightMultiplier*speed) * weight * duration
}

func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool*countPool) / mInKm / duration
}

func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	return (speed*swimmingCaloriesMeanSpeedShift + swimmingCaloriesWeightMultiplier) * weight * duration
}
