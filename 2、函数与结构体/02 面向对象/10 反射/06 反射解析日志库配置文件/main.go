package main

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	//"errors"
	"fmt"
	"reflect"
)

/*
	解析日志库的配置文件
*/

// config 是一个日志配置项，字段名大写
type config struct{
	FilePath string `conf:"file_path" db:"name"`
	FileName string `conf:"file_name"`
	MaxSize int64 `conf:"max_size"`
}

// parseConf 从conf文件中读取内容赋值给结构体指针
func parseConf(fileName string, result interface{}) (err error){

	// 0.前提条件，数据校验
	// result必须是一个ptr类型
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)
	if t.Kind() != reflect.Ptr{
		err = errors.New("result must be a pointer")
		return
	}

	// result 不是结构体也不行
	tElem := t.Elem()  // t.Elem() 返回元素类型
	if tElem.Kind() != reflect.Struct{
		err = errors.New("result must be a struct")
		return
	}

	// 1.打开文件
	data, err := ioutil.ReadFile(fileName)  // data:[]byte
	if err != nil{
		err = fmt.Errorf("打开文件[%s]配置失败!", fileName)
		return err
	}

	// 2.将读取的文件数据按照行分割，得到一个行的切片
	lineSlice := strings.Split(string(data), "\r\n")  // \r\n换行
	fmt.Println(lineSlice)

	// 3.一行行解析
	for index, line := range lineSlice{
		line = strings.TrimSpace(line)  // 去除字符串收尾的空白
		// 若当前行中没有数据，或者字符串开头是 # 则不处理
		if len(line) == 0 || strings.HasPrefix(line, "#"){
			// 忽略空行和注释
			continue
		}

		// 解析是不是正经的配置项（判断是不是有等于号）
		// 获取字符串中 = 的索引，若有则返回 = 的索引 无则返回 -1
		equalIndex := strings.Index(line, "=")
		if equalIndex == -1{
			err = fmt.Errorf("第{%d}行语法错误", index)
			return
		}

		// 4.找到需要赋值的配置
		// 按照=分割每一行，左边的是key，右边的是value
		key := line[:equalIndex]
		value := line[equalIndex + 1:]
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		// 利用反射 给 result 赋值
		// 遍历结构体的每一个字段和key比较，匹配上了就给value赋值
		for i := 0; i < tElem.NumField(); i++{

			field := tElem.Field(i)  // 根据类型获取字段名
			tag := field.Tag.Get("conf")  // 获取conf名称
			if key == tag{

				// 给匹配上的value赋值
				filedType := field.Type  // 获取每个字段的类型
				fmt.Println(filedType, 1111)
				fmt.Println(filedType.Kind(), 2222)

				// filedType.Kind() 这样写才能让case的反射包来做判断
				switch filedType.Kind(){ // 判断类型

				case reflect.String:
					// 获取值
					fieldValue := v.Elem().FieldByName(field.Name)  // 根据字段名获取字段值
					// 动态给结构体字段赋值
					fieldValue.SetString(value)
				case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
					// 字符串的 整型 需要做 字符串 --> 整型 的转换
					value64, _ := strconv.ParseInt(value, 10, 64)  // 转为int64
					// 修改结构体，通过i获取结构体对应的字段
					v.Elem().Field(i).SetInt(value64)
				}
			}
		}

	}
	return
}


func main() {

	var c = &config{}  // 用于读取配置数据的
	err := parseConf("./2、函数与结构体/02 面向对象/10 反射/06 反射解析日志库配置文件/xxx.conf", c)  // 调用解析配置函数
	if err != nil{
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", c)  // 解析之后的数据
}
