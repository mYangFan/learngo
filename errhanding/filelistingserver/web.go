package main

import (
	"net/http"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		err := handler(writer, request)
//		if err != nil {
//			switch {
//			case os.IsNotExist(err)
//			}
//		}
//	}
//}

func main()  {
	//http.HandleFunc("/list/", filelisting.HandleFileList)
	//
	//err := http.ListenAndServe(":8888", nil)
	//if err != nil {
	//	panic(err)
	//}
}