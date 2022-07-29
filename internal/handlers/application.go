package handlers

import "log"

//Application структура для хранения зависимостей логеров
type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}
