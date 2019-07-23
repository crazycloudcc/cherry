package logtool

import (
	"io"
	"os"
	"strconv"
	"time"
)

var i = 0

func (this *Logger) save(args string) error {
	if this.hour != int32(time.Now().Hour()) {
		i = 0
		creatTime := time.Now().Format("2006-01-02-15")
		this.fileName = creatTime + "-" + strconv.Itoa(i) + ".txt"
		file, err := os.Create(this.conf.Path + this.fileName)
		if err != nil {
			return err
		}
		this.file = file
		this.writeSize = 0
	}
	//判断已经写入文件大小,超过限定大小，重新创建新文件
	if this.writeSize+int32(len([]rune(args))) > this.conf.Size {
		i++
		creatTime := time.Now().Format("2006-01-02-15")
		this.fileName = creatTime + "-" + strconv.Itoa(i) + ".txt"
		file, err := os.Create(this.conf.Path + this.fileName)
		if err != nil {
			return err
		}
		this.file = file
		this.writeSize = 0
	}

	//写入文件
	num, err := io.WriteString(this.file, "<p>"+args+"</p>"+"\r\n")
	if err != nil {
		return err
	}
	//写入数据累加
	this.writeSize = this.writeSize + int32(num)
	return nil

}
