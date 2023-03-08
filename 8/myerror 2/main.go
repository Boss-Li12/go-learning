package main

import "fmt"

//这是一个特有的areaError结构体，它只需要实现Error()接口即可
//里面包含err信息
//实际上可以认为areaError继承了error
type areaError struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

//实现Error()接口，返回信息
func (e *areaError) Error() string {
	return e.err
}

//error的方法用来明确错误原因
//相当于子类的新方法
func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

//相当于子类的新方法
func (e *areaError) widthNegative() bool {
	return e.width < 0
}


func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	//错误返回
	if err != "" {
		err += "\n"
		//返回的error变量实际上就是新建一个areaError的错误类型（注意一般都加&)
		return 0, &areaError{err, length, width}
	} //err 文本用来给错误提示信息
	//正常返回
	return length * width, nil
}

func main() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		fmt.Print(err)
		//err.(*areaError)由地址变回值
		//类型断言
		if err, ok := err.(*areaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
	}
	fmt.Println("area of rect", area)
}
