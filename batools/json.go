package batools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*json解析相关*/

//读取本地json文件并解析到目标结构中
func LoadJsonTo(path string, conf interface{}) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(path+"加载失败:", err.Error())
		return err
	}
	err = json.Unmarshal(buf, conf)
	if err != nil {
		fmt.Println(path+"解析失败:", err.Error())
		return err
	}
	return nil
}
