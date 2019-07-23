package logtool

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

var logPath string
var templates *template.Template

//log服务开启
func StartService(filePath string, port string) {
	logPath = filePath
	templates = template.Must(template.ParseFiles("html/index.html"))
	http.HandleFunc("/", readLogHandler)
	http.ListenAndServe(":"+port, nil)
}

func readLogHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		request.ParseForm()
		// fmt.Println("name:", request.URL, len(request.Form["filename"]), request.Form["filename"])
		if len(request.Form["filename"]) > 0 {
			// fmt.Println("get filename")
			text, err := ioutil.ReadFile(logPath + request.Form["filename"][0])
			if err != nil {
				fmt.Println("读取错误：", err)
			}
			io.WriteString(writer, string(text))
			return
		} else {
			list := getFileList(logPath)
			var fileList []string
			for i := len(list) - 1; i >= 0; i-- {
				fileList = append(fileList, list[i])
			}

			// for _, value := range fileList {
			// 	fmt.Println(value)
			// }
			templates.ExecuteTemplate(writer, "index.html", fileList)
			return
		}

	} else {
		fmt.Println(request)
		return
	}
}
func getFileList(path string) []string {
	var files []string
	list, _ := ioutil.ReadDir(path)
	for _, info := range list {
		//遍历目录下的内容，获取文件详情，同os.Stat(filename)获取的信息
		info.Name() //文件名
		info.Sys()  //系统信息
		//fmt.Println(info.Name())
		files = append(files, info.Name())
		//if info.IsDir() == true {
		//	fmt.Println("是目录")
		//}
	}
	return files
}
