package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(st string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		fmt.Println("датасет пуст")
	}
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Printf("ошибка при парсинге %s элемента: %v\n", v, err)
			continue
		}

	}
	if len(dataset) > 0 {
		s, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка при выводе Actioninfo: %v\n", err)
		}
		fmt.Println(s)
	}
}
