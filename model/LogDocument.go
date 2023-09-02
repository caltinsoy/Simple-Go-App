package model

import "fmt"

type LogDocument struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	AnyString string `json:"anyString"`
}

func (t LogDocument) String() string {
	return fmt.Sprintf("[%d : %s]", t.ID, t.AnyString)
}
