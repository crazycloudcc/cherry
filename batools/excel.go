//excel读取静态表
package batools

import (
	"fmt"
	"os"
	"reflect"
	"server/exproperty"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

var xlsxData map[string]*exproperty.PropSets

var xlsxPath string

/************************************************************************/
// public
/************************************************************************/

//初始化所有xlsx配置
func LoadAllFiles(path string) {

	xlsxPath = path
	fmt.Println("load xlsx path: ", path)
	xlsxData = make(map[string]*exproperty.PropSets)

	xlsxData[exproperty.XLSX_CONSTANT_FORMULA] = loadConstantExcel(exproperty.XLSX_CONSTANT_FORMULA)
	xlsxData[exproperty.XLSX_CONSTANT_GLOBAL] = loadConstantExcel(exproperty.XLSX_CONSTANT_GLOBAL)

	arrXLSX := exproperty.ARRAY_XLSX
	for i := 0; i < len(arrXLSX); i++ {
		xlsxData[arrXLSX[i]] = loadExcel(arrXLSX[i])
	}
}

//获取指定配置中的一条属性
//fname：xlsx文件名，不带后缀
//tid：表中一行数据存储后的索引
func GetProperty(fname string, tid int) interface{} {
	if xlsxData[fname] == nil {
		fmt.Println(fmt.Sprintf("ERROR: not found data! fname:[%s], tid:[%d]", fname, tid))
		return nil
	}
	return xlsxData[fname].GetProp(tid)
}

//获取指定数据表集合
func GetPropSets(fname string) *exproperty.PropSets {
	return xlsxData[fname]
}

// 获取指定数据表的有效行数.
func GetCount(fname string) int {
	l := len(xlsxData[fname].Props)
	return l
}

/************************************************************************/
// private
/************************************************************************/

//读取单个excel表格存储到Propertys中
func loadExcel(fname string) *exproperty.PropSets {
	pSets := new(exproperty.PropSets)
	pSets.Fname = fname
	xf, err := xlsx.OpenFile(fmt.Sprintf("%s/%s.xlsx", xlsxPath, fname))
	if err != nil {
		fmt.Printf("加载%s文件失败:%v\n", fname, err.Error())
		return nil
	}
	sheetCount := len(xf.Sheets)
	if sheetCount > 0 {
		//暂时只考虑单表单的情况
		sheet := xf.Sheets[0]
		readSheet(sheet, pSets)
	}
	return pSets
}

//读取表单
func readSheet(sheet *xlsx.Sheet, pSets *exproperty.PropSets) error {
	//读取描述字段
	readRowToSlice(sheet.Rows[0], &pSets.Descriptions)
	//读取属性字段名
	readRowToSlice(sheet.Rows[2], &pSets.Fields)
	rowNums, colNums := len(sheet.Rows), len(pSets.Fields)
	var row *xlsx.Row
	pSets.Props = make(map[int]interface{})
	for i := 3; i < rowNums; i++ {
		row = sheet.Rows[i]
		//反射实例化一个property类型

		prop := refPropertyWithName(pSets.Fname)
		str, _ := row.Cells[0].String()
		if str == "" {
			fmt.Printf("<%s.xlsx> 行未填充满 %d / %d\n", pSets.Fname, i, rowNums)
			break
		}
		tid, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("<%s.xlsx> 配置表读取异常 [%d] / [%d] [%v] \n", pSets.Fname, i, rowNums, err)
			break
		}
		for j := 0; j < colNums; j++ {
			if pSets.Fields[j] != "" {
				if str, _ := row.Cells[j].String(); str == "" {
					fmt.Printf("%s表数据配置tid:[%d] - %s缺少数据 %d / %d\n", pSets.Fname, tid, pSets.Fields[j], j, colNums)
					break //结束列
				}
				fillReflectFields(pSets.Fname, &prop, pSets.Fields[j], row.Cells[j])
				pSets.Props[tid] = prop.Interface()
			}
		}
	}
	return nil
}

//行数据转数组
func readRowToSlice(row *xlsx.Row, vals *[]string) {
	cellCount := len(row.Cells)
	for i := 0; i < cellCount; i++ {
		value, err := row.Cells[i].String()
		if err != nil {
			fmt.Println("读取单元格数据出错", err.Error())
			break
		}
		*vals = append(*vals, value)
	}
}

//从路径中提取文件名（不带后缀）
func extractFname(path string) string {
	str := strings.Split(path, "/")
	return strings.Split(str[len(str)-1], ".")[0]
}

//根据名称反射一个对象
func refPropertyWithName(fname string) reflect.Value {
	p := exproperty.XlsxNameMap[fname]
	if p == nil {
		fmt.Printf("没有%s.xlsx对应的Property结构体\n", fname)
		os.Exit(1)
	}
	t := reflect.TypeOf(p)
	prop := reflect.New(t).Elem()
	return prop
}

//填充结构体的反射字段
//key：表中字段名，对应行标为3
func fillReflectFields(fname string, prop *reflect.Value, key string, cell *xlsx.Cell) {
	//获取结构类型
	propType := prop.Type()
	//	fmt.Printf("----------------------------------------- %v\n", prop.Type())
	//获取结构中字段的数量
	fieldNums := prop.NumField()
	for i := 0; i < fieldNums; i++ {
		//通过类型查找字段，便于获取字段名
		fieldName := propType.Field(i).Name
		//获取单个字段的类型
		fieldType := prop.Field(i).Kind()
		key = strings.Split(key, "[")[0]
		if strings.Compare(strings.ToLower(key), strings.ToLower(fieldName)) == 0 {
			switch fieldType {
			case reflect.Bool:
				value := cell.Bool()
				prop.Field(i).SetBool(value)
				return
			case reflect.Int:
				value, _ := cell.Int64()
				prop.Field(i).SetInt(value)
				return
			case reflect.Int32:
				value, _ := cell.Int64()
				prop.Field(i).SetInt(value)
				return
			case reflect.String:
				value, _ := cell.String()
				prop.Field(i).SetString(value)
				return
			case reflect.Float32:
				value, _ := cell.Float()
				prop.Field(i).SetFloat(value)
				return
			case reflect.Slice:
				value, _ := cell.String()
				vals := strings.Split(value, ",")
				//				fmt.Printf("%s | key：%s value：%s %v\n", fname, fieldName, value, vals)
				field := prop.Field(i)
				formatSliceToField(fname, vals, &field)
				return
			case reflect.Struct:
				fs := prop.Field(i)
				fCount := fs.NumField()
				value, _ := cell.String()
				vals := parseCellToSliceArr(value)
				//				fmt.Println("子结构体NumField  :", fieldName, fCount, " ", vals)
				for i := 0; i < fCount; i++ {
					field := fs.Field(i)
					//内嵌的类型一定是切片：约定好的
					//					fmt.Printf(" %s======   %v %v\n", fname, vals, field.String())
					if field.Kind() == reflect.Slice {
						formatSliceToField(fname, vals[i], &field)
					} else {
						fmt.Println("内嵌类型不是切片！")
						return
					}
				}

				return
			}
		}
	}
}

//解析单元格中的数据到切片数组中
//单元格数据格式：（x1,x2;y1,y2）
//目标格式为：[[x1,y1],[x2,y2]]
func parseCellToSliceArr(v string) [][]string {
	if v == "null" {
		return [][]string{nil, nil}
	}
	//拆分为[x1,x2][y1,y2]
	arr1 := strings.Split(v, ";")
	//分成几部分
	eleCount := len(strings.Split(arr1[0], ","))
	//结果数组的个数取决于元素的个数
	rArr := make([][]string, eleCount)
	for i := 0; i < eleCount; i++ {
		rArr[i] = make([]string, len(arr1))
	}
	for i := 0; i < len(arr1); i++ {
		tmpArr := strings.Split(arr1[i], ",")
		for j := 0; j < eleCount; j++ {
			rArr[j][i] = tmpArr[j]
			// println(len(rArr[j]))
		}
		// fmt.Println("单一部分数据：", arr1[i], " ", eleCount, "结果：", rArr)
	}
	return rArr
}

//格式化[]interface到目标类型的切片中
func formatSliceToField(fn string, s []string, f *reflect.Value) {
	eleType := f.Type().Elem().Kind()
	var err error
	switch eleType {
	case reflect.Int:
		arr := make([]int, len(s))
		for j := 0; j < len(s); j++ {
			arr[j], err = strconv.Atoi(s[j])
			if err != nil && s[j] != "null" {
				panic(fmt.Sprintf("要求int类型！%v, file[%s]", eleType, fn))
				break
			}
		}
		// fmt.Println(arr)
		tmp := *f //获取字段reflect.value
		tmp = reflect.AppendSlice(tmp, reflect.ValueOf(arr))
		f.Set(tmp)
	case reflect.Int32:
		arr := make([]int32, len(s))
		var k int
		for j := 0; j < len(s); j++ {
			k, err = strconv.Atoi(s[j])
			if err != nil && s[j] != "null" {
				panic(fmt.Sprintf("要求int32类型！%v, file[%s]", eleType, fn))
				break
			}
			arr[j] = int32(k)
		}
		// fmt.Println(arr)
		tmp := *f //获取字段reflect.value
		tmp = reflect.AppendSlice(tmp, reflect.ValueOf(arr))
		f.Set(tmp)
	case reflect.Float32:
		arr := make([]float32, len(s))
		var v float64
		for j := 0; j < len(s); j++ {
			v, err = strconv.ParseFloat(s[j], 64)
			if err != nil && s[j] != "null" {
				panic("要求float32类型！")
				break
			}
			arr[j] = float32(v)
		}
		tmp := *f //获取字段reflect.value
		tmp = reflect.AppendSlice(tmp, reflect.ValueOf(arr))
		f.Set(tmp)
	case reflect.String:
		tmp := *f //获取字段reflect.value
		tmp = reflect.AppendSlice(tmp, reflect.ValueOf(s))
		f.Set(tmp)

	case reflect.Float64:
		arr := make([]float64, len(s))
		for j := 0; j < len(s); j++ {
			arr[j], err = strconv.ParseFloat(s[j], 64)
			if err != nil && s[j] != "null" {
				panic("要求float64类型！")
				break
			}
		}
		tmp := *f //获取字段reflect.value
		tmp = reflect.AppendSlice(tmp, reflect.ValueOf(arr))
		f.Set(tmp)
	}
}

/************************************************************************/
// 特殊处理constant_formula的tid：自动生成int型索引
/************************************************************************/
//读取单个constant_formula表格存储到Propertys中
func loadConstantExcel(fname string) *exproperty.PropSets {
	pSets := new(exproperty.PropSets)
	pSets.Fname = fname
	xf, err := xlsx.OpenFile(fmt.Sprintf("%s/%s.xlsx", xlsxPath, fname))
	if err != nil {
		fmt.Printf("加载%s文件失败:%v\n", fname, err.Error())
		return nil
	}
	sheetCount := len(xf.Sheets)
	if sheetCount > 0 {
		//暂时只考虑单表单的情况
		sheet := xf.Sheets[0]
		readConstantSheet(sheet, pSets)
	}
	return pSets
}

//读取表单
func readConstantSheet(sheet *xlsx.Sheet, pSets *exproperty.PropSets) error {
	//读取描述字段
	readRowToSlice(sheet.Rows[0], &pSets.Descriptions)
	//读取属性字段名
	readRowToSlice(sheet.Rows[2], &pSets.Fields)
	rowNums, colNums := len(sheet.Rows), len(pSets.Fields)
	var row *xlsx.Row
	pSets.Props = make(map[int]interface{})
	tid := 0
	for i := 3; i < rowNums; i++ {
		row = sheet.Rows[i]
		//反射实例化一个property类型

		prop := refPropertyWithName(pSets.Fname)
		for j := 0; j < colNums; j++ {
			if pSets.Fields[j] != "" {
				if str, _ := row.Cells[j].String(); str == "" {
					fmt.Printf("%s表数据配置tid:[%d] - %s缺少数据, %d / %d\n", pSets.Fname, tid, pSets.Fields[j], j, colNums)
					break //结束列
				}
				fillReflectFields(pSets.Fname, &prop, pSets.Fields[j], row.Cells[j])
				pSets.Props[tid] = prop.Interface()
			}
		}
		tid++
	}
	return nil
}
