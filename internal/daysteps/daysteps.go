package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	datastringSlise := strings.Split(datastring, ",")
	if len(datastringSlise) != 2 {
		return errors.New("Неверное значение в daysteps/Parse")
	}

	step, err := strconv.Atoi(datastringSlise[0])
	if err != nil {
		return errors.New("Неверное значение шагов")
	}
	if step <= 0 {
		return errors.New("Неверное значение шагов")
	}
	ds.Steps = step

	time, err := time.ParseDuration(datastringSlise[1])
	if err != nil {
		return errors.New("Ошибка времени")
	}
	if time <= 0 {
		return errors.New("Неправильное время")
	}
	ds.Duration = time

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height) // Дистанция в км
	if dist <= 0 {
		return "", errors.New("неверная дистанция в ActionInfo")
	}

	// Калории при ходьбе
	walkColories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", errors.New("ошибка рассчёта калорий для ходьбы в ActionInfo")
	}
	if walkColories <= 0 {
		return "", errors.New("калорий меньше или ноль в ActionInfo")
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, walkColories), nil

}
