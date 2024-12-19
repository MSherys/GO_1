package ftracker

import (
	"fmt"
	"math"
)

const (
	mInKm       = 1000.0 // Метров в километре
	stepLength  = 0.65   // Длина шага в метрах
	minInH      = 60.0   // Минут в часе
	swimFactor  = 1.1    // Коэффициент для расчёта калорий плавания
	runFactor   = 18.0   // Коэффициент для расчета калорий при беге
	runCoef     = 1.79   // Дополнительный коэффициент для расчёта калорий при беге
	walkFactor1 = 0.035  // Первый коэффициент для расчёта калорий при ходьбе
	walkFactor2 = 0.029  // Второй коэффициент для расчёта калорий при ходьбе
)

// Расчёт дистанции в километрах на основе количества шагов.
func distance(action int) float64 {
	return float64(action) * stepLength / mInKm
}

// Средняя скорость в км/ч для бега и ходьбы.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return distance(action) / duration
}

// Средняя скорость в плавании.
func swimmingSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	totalDistance := float64(lengthPool*countPool) / mInKm
	return totalDistance / duration
}

// Расчёт калорий для бега.
func RunningSpentCalories(action int, duration, weight float64) float64 {
	speed := meanSpeed(action, duration)
	return (runFactor * speed * runCoef * weight / mInKm) * duration * minInH
}

// Расчёт калорий для ходьбы.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	speed := meanSpeed(action, duration) * 1000 / 3600 // Скорость в м/с
	return (walkFactor1*weight + math.Pow(speed, 2)/(height)*walkFactor2*weight) * duration * minInH
}

// Расчёт калорий для плавания.
func SwimmingSpentCalories(lengthPool, countPool int, weight float64) float64 {
	totalDistance := float64(lengthPool*countPool) / mInKm
	return totalDistance * swimFactor * weight
}

// Вывод информации о тренировке.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	switch trainingType {
	case "Бег":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := RunningSpentCalories(action, duration, weight)
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			trainingType, duration, distance, speed, calories,
		)
	case "Ходьба":
		distance := distance(action)
		speed := meanSpeed(action, duration)
		calories := WalkingSpentCalories(action, duration, weight, height)
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			trainingType, duration, distance, speed, calories,
		)
	case "Плавание":
		distance := float64(lengthPool*countPool) / mInKm
		speed := swimmingSpeed(lengthPool, countPool, duration)
		calories := SwimmingSpentCalories(lengthPool, countPool, weight)
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			trainingType, duration, distance, speed, calories,
		)
	default:
		return "неизвестный тип тренировки"
	}
}
