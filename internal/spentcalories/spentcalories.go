package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	//lenStep                    = 0.65 // средняя длина шага. (Убрал, ибо по ТЗ использовать не нужно)
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	if data == "" {
		return 0, "", time.Duration(0), errors.New("Тренировка: Пустая строка!")
	}

	sliceData := strings.Split(data, ",")
	if len(sliceData) != 3 {
		return 0, "", time.Duration(0), errors.New("Тренировка: Недостаточно аргументов в строке!")
	}

	steps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return 0, "", time.Duration(0), errors.New("Тренировка: Первый аргумент должен быть целым числом!")
	}

	if steps <= 0 {
		return 0, "", time.Duration(0), errors.New("Тренировка: Количество шагов не может быть отрицательным!")
	}

	training := sliceData[1]
	if training == "" {
		return 0, "", time.Duration(0), errors.New("Тренировка: Нужно название вида тренировки!")
	}

	duration, err := time.ParseDuration(sliceData[2])
	if err != nil {
		return 0, "", time.Duration(0), err
	}

	if duration <= 0 {
		return 0, "", time.Duration(0), errors.New("Тренировка: Продолжительность не может быть отрицательной или нулём!")
	}

	return steps, training, duration, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	length := height * stepLengthCoefficient
	stepDistance := (float64(steps) * length) / mInKm

	return stepDistance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 || steps <= 0 {
		return 0
	}

	avg := distance(steps, height) / duration.Hours()
	return avg
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, training, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	switch training {

	case "Бег":
		trainingDistance := distance(steps, height)
		speed := meanSpeed(steps, height, duration)

		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		doneStr := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f\n", training, duration.Hours(), trainingDistance, speed, calories)
		return doneStr, nil

	case "Ходьба":
		trainingDistance := distance(steps, height)
		speed := meanSpeed(steps, height, duration)

		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}

		doneStr := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f\n", training, duration.Hours(), trainingDistance, speed, calories)

		return doneStr, nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", training)
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Расчёт калорий (бег): Параметры не реалистичны!")
	}
	avgSpeed := meanSpeed(steps, height, duration)

	calories := (weight * avgSpeed * duration.Minutes()) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	calories *= walkingCaloriesCoefficient
	return calories, nil
}
