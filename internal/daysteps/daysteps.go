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
	if data == "" {
		return 0, time.Duration(0), errors.New("Шаги: Пустая строка!")
	}

	sliceData := strings.Split(data, ",")
	if len(sliceData) != 2 || sliceData[0] == "" || sliceData[1] == "" {
		return 0, time.Duration(0), errors.New("Шаги: Недостаточно аргументов в строке!")
	}

	steps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return 0, time.Duration(0), errors.New("Шаги: Первый аргумент должен быть целым числом!")
	}

	if steps <= 0 {
		return 0, time.Duration(0), errors.New("Шаги: Количество шагов не может быть отрицательным!")
	}

	walkTime, err := time.ParseDuration(sliceData[1])
	if err != nil {
		return 0, time.Duration(0), err
	}

	if walkTime <= 0 {
		return 0, time.Duration(0), errors.New("Шаги: Продолжительность не может быть отрицательной или нулём!")
	}

	return steps, walkTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, walkTime, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, walkTime)
	if err != nil {
		return ""
	}

	doneStr := fmt.Sprintf("Количество шагов: %v.\nДистанция составила %0.2f км.\nВы сожгли %0.2f ккал.\n", steps, distance, calories)
	return doneStr
}
