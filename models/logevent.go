package models

import "github.com/astaxie/beego/orm"

import "log"

type LogEvent struct {
	Id            int
	LogStreamName string
	IngestionTime int64  `orm:"null"`
	Message       string `orm:"null"`
	Timestamp     int64  `orm:"null"`
}

func init() {
	orm.RegisterModel(new(LogEvent))
}

func InsertLogEvent(m *LogEvent) {
	o := orm.NewOrm()
	_, err := o.Insert(m)
	if err != nil {
		log.Fatal(err.Error())
	}
}
