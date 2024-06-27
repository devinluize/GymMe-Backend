package helper

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Paniciferror(err error) {
	if err != nil {
		panic(err)
	} else {
		return
	}
}

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		tx.Rollback()
		logrus.Info(err)
	} else {
		tx.Commit()
	}
}
