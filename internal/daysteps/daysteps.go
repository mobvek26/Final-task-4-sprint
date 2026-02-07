package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid format")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, errors.New("invalid steps count format")
	}
	if steps <= 0 {
		return 0, 0, errors.New("step count must be greater  than 0")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, errors.New("invalid duration format")
	}
	if duration <= 0 {
		return 0, 0, errors.New("duration must be greter than 0")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("Error parssing packege: %v", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distance := float64(steps) * stepLength
	distanceKm := distance / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Printf("Error calculating calorie: %v", err)
		return ""
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
	return result
}
