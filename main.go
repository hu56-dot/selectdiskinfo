/*
 * @Author: your name
 * @Date: 2021-04-28 10:46:48
 * @LastEditTime: 2021-05-28 21:26:40
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \Go\src\systemcommand\main.go
 */
package main

import (
	"fmt"
	"systemcommand/bytesOutputHandle"
)

func main() {
	fmt.Printf("解析结果如下：%s\n", bytesOutputHandle.HandleString())

}
