package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	datastringSlise := strings.Split(datastring, ",")
	if len(datastringSlise) != 3 {
		return errors.New("Неверные данные Parse")
	}

	step, err := strconv.Atoi(datastringSlise[0])
	if err != nil {
		return errors.New("Неверное значение шагов")
	}
	if step <= 0 {
		return errors.New("Неверное значение шагов")
	}
	t.Steps = step

	if datastringSlise[1] == "" {
		return errors.New("Пустая строка")
	}
	t.TrainingType = datastringSlise[1]

	time, err := time.ParseDuration(datastringSlise[2])
	if err != nil {
		return errors.New("Ошибка времени")
	}
	if time <= 0 {
		return errors.New("Неправильное время")
	}
	t.Duration = time

	return nil
}

func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height) // Дистанция в км
	if dist <= 0 {
		return "", errors.New("Неверная дистанция в ActionInfo")
	}

	sp := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration) // скорость
	if sp <= 0 {
		return "", errors.New("неверная скорость в ActionInfo")
	}

	d := t.Duration
	hours := float64(d) / float64(time.Hour)

	switch t.TrainingType {
	case "Бег":
		runColories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", errors.New("ошибка рассчёта калорий для бега в ActionInfo")
		}
		if runColories <= 0 {
			return "", errors.New("калорий меньше или ноль в ActionInfo")
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, hours, dist, sp, runColories), nil
	case "Ходьба":
		walkColories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", errors.New("ошибка рассчёта калорий для ходьбы в ActionInfo")
		}
		if walkColories <= 0 {
			return "", errors.New("калорий меньше или ноль в ActionInfo")
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, hours, dist, sp, walkColories), nil

	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}
