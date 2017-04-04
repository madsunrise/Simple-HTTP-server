//package main
//
//import (
//	"strings"
//	"os"
//	"io/ioutil"
//)
//
//
//type File struct {
//	status int
//	length int
//	content_type string
//	content []byte
//}
//
//var ROOT_PATH = "./DOCUMENT_ROOT/"
//
//func GetFile(url string) (File) {
//	if strings.Contains(url, "../") {
//		return File {
//			403,
//			0,
//			"",
//			nil,
//		}
//	}
//
//	request_path := ROOT_PATH + url;
//	info, err := os.Stat(request_path)
//	if err != nil {
//		if os.IsNotExist(err) {
//			return File{
//				404,
//				0,
//				"",
//				nil,
//			}
//		} else {
//			return File{
//				403,
//				0,
//				"",
//				nil,
//			}
//		}
//	}
//
//
//
//	file, err := ioutil.ReadFile(request_path)
//	if err != nil {
//		return File{
//			403,
//			0,
//			"",
//			nil,
//		}
//	}
//
//	return File {
//		200,
//		int(info.Size()),
//			"image/jpeg",
//		file,
//	}
//}
