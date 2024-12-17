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
	cmInM     = 100   // Количество сантиметров в метре

	// Множители для расчета калорий
	runningCaloriesMeanSpeedMultiplier = 18   // Множитель для бега
	runningCaloriesMeanSpeedShift      = 1.79 // Сдвиг для бега

	walkingCaloriesWeightMultiplier = 0.035 // Множитель для веса при ходьбе
	walkingSpeedHeightMultiplier    = 0.029 // Множитель для роста при ходьбе

	swimmingCaloriesMeanSpeedShift   = 1.1 // Сдвиг для плавания
	swimmingCaloriesWeightMultiplier = 2   // Множитель для веса при плавании
)

// distance возвращает дистанцию (в километрах), которую преодолел пользователь за время тренировки.
func distance(action int) float64 {
	// Возвращаем расстояние, которое прошло действие в метрах, переведенное в километры
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return distance(action) / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	var dist float64
	var speed float64
	var calories float64

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
		// Исправленный расчет дистанции плавания
		dist = float64(lengthPool*countPool) / 1000 // Дистанция в километрах
		speed = swimmingMeanSpeed(lengthPool, countPool, duration)
		calories = SwimmingSpentCalories(lengthPool, countPool, duration, weight)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

// RunningSpentCalories возвращает количество потраченных калорий при беге.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	speed := meanSpeed(action, duration)
	// Формула для расчета калорий при беге
	return ((runningCaloriesMeanSpeedMultiplier * speed * runningCaloriesMeanSpeedShift) * weight / mInKm * duration * minInH)
}

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	speed := meanSpeed(action, duration)
	// Переводим скорость в метры в секунду
	speedInMetersPerSecond := speed * mInKm / (minInH * 60)
	// Формула для расчета калорий при ходьбе
	return ((walkingCaloriesWeightMultiplier * weight) + (math.Pow(speedInMetersPerSecond, 2) / height * walkingSpeedHeightMultiplier * weight)) * duration * minInH
}

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	// Средняя скорость плавания
	return float64(lengthPool*countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	// Формула для расчета калорий при плавании
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}
