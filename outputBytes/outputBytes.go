/*
 * @Author: your name
 * @Date: 2021-05-18 09:30:23
 * @LastEditTime: 2021-05-28 13:44:11
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \Go\src\systemcommand\outputBytes\outputBytes.go
 */
package outputBytes

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// InputCommand 执行WINDOWS命令获取字节输出函数
func InputCommand() string {
	//str := exec.Command("cmd", "-C","diskpart","list volume")
	//str := exec.Command("cmd","/C","wmic","volume get BlockSize")

	str := exec.Command("/bin/bash", "-c", `lsblk  -P -p -b -o NAME,SIZE,TYPE,MOUNTPOINT,PKNAME,FSTYPE`)

	//创建获取命令输出管道
	stdout, err := str.StdoutPipe() //当上面的命令开始运行时，这个管道就会连接命令输出
	if err != nil {
		fmt.Printf("Error: can not obtain stdout pipe for command:%s\n", err)

	}
	defer stdout.Close() //保证关闭输出流

	//--------------------------上面先把管道建起来，等待输出-------------
	if err := str.Start(); err != nil { //用于启动命令
		fmt.Println("Error: the command is err,", err) //这里为什么不是获取命令行的提示错误

	}

	bytes, err := ioutil.ReadAll(stdout) //读取所有的输出
	if err != nil {
		fmt.Println("Error: ReadALL Stdout:", err.Error()) //读取错误

	}
	//fmt.Printf("%v",bytes)
 
	return string(bytes)
}


