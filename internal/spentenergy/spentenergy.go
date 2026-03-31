package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("Ноль шагов")
	}
	if weight <= 0 {
		return 0, errors.New("Вес равен нулю")
	}
	if height <= 0 {
		return 0, errors.New("Рост равен нулю")
	}
	if duration <= 0 {
		return 0, errors.New("время равно нулю")
	}
	speed := MeanSpeed(steps, height, duration)
	t := duration.Minutes()
	res := ((weight * speed * t) / minInH) * walkingCaloriesCoefficient
	return res, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("Ноль шагов")
	}
	if weight <= 0 {
		return 0, errors.New("Вес равен нулю")
	}
	if height <= 0 {
		return 0, errors.New("Рост равен нулю")
	}
	if duration <= 0 {
		return 0, errors.New("время равно нулю")
	}
	speed := MeanSpeed(steps, height, duration)
	t := duration.Minutes()
	res := (weight * speed * t) / minInH

	return res, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		return 0
	}

	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	t := duration.Hours()
	return dist / t
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		return 0
	}

	if height <= 0 {
		return 0
	}

	return ((height * stepLengthCoefficient) * float64(steps)) / mInKm
}
