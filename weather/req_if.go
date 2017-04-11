package weather
//package main


type weather_request interface{
	Init(ak string) (error)
	Do_Request(location string) (err error)
	Get_PM25() (pm25 int, err error)
}

