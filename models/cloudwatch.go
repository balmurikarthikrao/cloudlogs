package models

import "github.com/astaxie/beego/orm"

import "log"

type LogStreams struct {
	Id                  int    `json:"-"`
	Arn                 string `orm:"null";`
	CreationTime        int64  `orm:"null"`
	FirstEventTimestamp int64  `orm:"null"`
	LastEventTimestamp  int64  `orm:"null"`
	LastIngestionTime   int64  `orm:"null"`
	LogStreamName       string `orm:"null"`
	StoredBytes         int64  `orm:"null"`
	UploadSequenceToken string `orm:"null"`
}

func init() {
	orm.RegisterModel(new(LogStreams))
}

func InsertLog(Logdata *LogStreams) (Data *LogStreams, err error) {

	o := orm.NewOrm()
	_, err = o.Insert(Logdata)

	return Logdata, err
}

func ReadLogarn() (rows []orm.Params, count int64) {

	// o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("log_stream_name").
		From("log_streams")

	sql := qb.String()
	o := orm.NewOrm()
	count, err := o.Raw(sql).Values(&rows)
	if err != nil {
		log.Fatal(err.Error())
	}
	return

}
