package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("неверный формат количества шагов")
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("неверный формат времени")
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	return steps, parts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength
	distanceKm := distance / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	durationInHours := duration.Hours()
	meanSpeed := dist / durationInHours
	return meanSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	var calories float64
	var calErr error

	switch activityType {
	case "Ходьба":
		calories, calErr = WalkingSpentCalories(steps, weight, height, duration)
		if calErr != nil {
			log.Println(calErr)
			return "", calErr
		}
	case "Бег":
		calories, calErr = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(calErr)
			return "", calErr
		}
	default:
		err := fmt.Errorf("неизвестный тип тренировки")
		log.Println(err)
		return "", err
	}
	durationInHours := duration.Hours()

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, durationInHours, dist, speed, calories)
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0.0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0.0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0.0, fmt.Errorf("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0.0, fmt.Errorf("вес должен быть больше 0")
	}
	if height <= 0 {
		return 0.0, fmt.Errorf("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0.0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient
	return calories, nil
}
