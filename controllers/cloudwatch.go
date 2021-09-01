package controllers

import (
	"cloud/models"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type MainController struct{}

func (MainController) LogStreams(c *gin.Context) {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("secret key", "token", ""),
	})
	if err != nil {
		fmt.Println("Got error in session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//// Get the configured credentials
	// details, eorr := sess.Config.Credentials.Get()

	svc := cloudwatchlogs.New(sess)

	///DescribeLogStreamsPages Method(Describes the Log information )

	var LogStreamData *cloudwatchlogs.DescribeLogStreamsOutput
	pageNum := 0
	eror := svc.DescribeLogStreamsPages(&cloudwatchlogs.DescribeLogStreamsInput{
		Descending:   aws.Bool(true),
		LogGroupName: aws.String("/aws/lambda/BVAPIDEV"),
	},
		func(page *cloudwatchlogs.DescribeLogStreamsOutput, lastPage bool) bool {
			pageNum++
			LogStreamData = page
			return pageNum < 1
		})
	// fmt.Println(Lo)
	if eror != nil {
		fmt.Println("Got error getting log stream:")
		fmt.Println(eror.Error())
		os.Exit(1)
	}

	// Logdata.Arn = LogStreamData.arn
	fmt.Println(LogStreamData)
	var users []*models.LogStreams
	// var user *models.LogStreams
	for _, v := range LogStreamData.LogStreams {
		Logdata := models.LogStreams{}
		if v.Arn != nil {
			Logdata.Arn = *v.Arn
		}
		if v.CreationTime != nil {
			Logdata.CreationTime = *v.CreationTime
		}
		if v.FirstEventTimestamp != nil {
			Logdata.FirstEventTimestamp = *v.FirstEventTimestamp
		}
		if v.LastEventTimestamp != nil {
			Logdata.LastEventTimestamp = *v.LastEventTimestamp
		}
		if v.LastIngestionTime != nil {
			Logdata.LastIngestionTime = *v.LastIngestionTime
		}
		if v.LogStreamName != nil {
			Logdata.LogStreamName = *v.LogStreamName
		}
		if v.StoredBytes != nil {
			Logdata.StoredBytes = *v.StoredBytes
		}
		if v.UploadSequenceToken != nil {
			Logdata.UploadSequenceToken = *v.UploadSequenceToken
		}

		user, err := models.InsertLog(&Logdata)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)

	}

}

func (MainController) LogEvents(c *gin.Context) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("API key secret", "token", ""),
	})
	if err != nil {
		fmt.Println("Got error in session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//// Get the configured credentials
	// details, eorr := sess.Config.Credentials.Get()

	svc := cloudwatchlogs.New(sess)
	rows, _ := models.ReadLogarn()
	for _, v := range rows {
		// fmt.Println("value:", v["log_stream_name"])
		value, ok := v["log_stream_name"].(string)
		if ok {
			resp, err := svc.GetLogEvents(&cloudwatchlogs.GetLogEventsInput{
				Limit:         aws.Int64(100),
				LogGroupName:  aws.String("/aws/lambda/BVAPIDEV"), //LOG-GROUP-NAME
				LogStreamName: aws.String(value),                  //LOG-STREAM-NAME
			})
			if err != nil {
				log.Fatal()
			}

			for _, v := range resp.Events {
				Events := models.LogEvent{}
				Events.LogStreamName = value
				Events.IngestionTime = *v.IngestionTime
				Events.Message = *v.Message
				Events.Timestamp = *v.Timestamp

				models.InsertLogEvent(&Events)
			}
		}

	}

}
