package main

import "fmt"

// паттер =н для работы несовместимых типов с одним интерфейсом

// есть интерфейс Service с методом, который пересылает данные в виде xml
// есть json структура и к ней адаптер, который реализует этот интерфейс  

// плюсы: сокрытие
// минусы: усложнение/избыточность


//// один пакет ////

type JsonModel struct {
	// ...
}

func (j JsonModel) ConvertJsonToXML() string {
	return "xml"
}

// адаптер
type JsonAdapter struct {
	jsonModel *JsonModel
}

func (adapter *JsonAdapter) SendXMl() {
	xml := adapter.jsonModel.ConvertJsonToXML()
	fmt.Println("send", xml)
}

//// второй ////

type Service interface {
	SendXMl()
}

type XMLModel struct {

}

func (s *XMLModel) SendXMl() {
	fmt.Println("send")
}

func main() {
	a := JsonAdapter{&JsonModel{}}
	a.SendXMl()
}