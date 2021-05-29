/*
 * @Author: your name
 * @Date: 2021-05-18 10:36:23
 * @LastEditTime: 2021-05-28 19:33:50
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \Go\src\systemcommand\bytesOutputHandle\handleStr_linux.go
 */
package bytesOutputHandle

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"systemcommand/diskInfoStruct"
	"systemcommand/outputBytes"
)

// HandleStringLinux HandleString --------------------------------------------linux下的命令解析
func HandleString() []byte {
	outString := outputBytes.InputCommand()
	//outString := "NAME=\"/dev/sda\" SIZE=\"21474836480\" TYPE=\"disk\" MOUNTPOINT=\"\" PKNAME=\"\" FSTYPE=\"\"\nNAME=\"/dev/sda1\" SIZE=\"1073741824\" TYPE=\"part\" MOUNTPOINT=\"/boot\" PKNAME=\"/dev/sda\" FSTYPE=\"xfs\"\nNAME=\"/dev/sda2\" SIZE=\"20400046080\" TYPE=\"part\" MOUNTPOINT=\"\" PKNAME=\"/dev/sda\" FSTYPE=\"LVM2_member\"\nNAME=\"/dev/mapper/centos-root\" SIZE=\"18249416704\" TYPE=\"lvm\" MOUNTPOINT=\"/\" PKNAME=\"/dev/sda2\" FSTYPE=\"xfs\"\nNAME=\"/dev/mapper/centos-swap\" SIZE=\"2147483648\" TYPE=\"lvm\" MOUNTPOINT=\"[SWAP]\" PKNAME=\"/dev/sda2\" FSTYPE=\"swap\"\nNAME=\"/dev/sr0\" SIZE=\"4712300544\" TYPE=\"rom\" MOUNTPOINT=\"/run/media/jarnmaster/CentOS 7 x86_64\" PKNAME=\"\" FSTYPE=\"iso9660\"  "

	OK1 := getQuotedStrings1(outString)
	for k, v := range OK1 {
		//a1 := strings.Replace(v, "\"", " ", -1)
		a1 := strings.Replace(v, "[^a-zA-Z]", " ", 1)
		OK1[k] = a1
	}

	//fmt.Println(OK1)
	//进行基本的结构体构建
	allDiskInfo := make(map[string]diskInfoStruct.DiskInfoLinux)
	for i := 0; i < len(OK1)/6; i++ {
		var diskInfo diskInfoStruct.AllDiskInfoLinux
		diskInfo.Info = diskInfoStruct.DiskInfoLinux{Name: OK1[i*6], Size_: toUnit64(OK1[i*6+1]), Type: OK1[i*6+2], MountPoint: OK1[i*6+3], PkName: OK1[i*6+4], FsType: OK1[i*6+5]}
		diskInfo.Name = OK1[i*6]
		allDiskInfo[diskInfo.Name] = diskInfo.Info
	}
	//fmt.Println(allDiskInfo)
	a4, err := json.MarshalIndent(allDiskInfo, "", " ")
	if err != nil {
		fmt.Println("json 错误")
	}
	fmt.Printf("%s\n", a4)

	return a4
}

//类型转换函数
func toUnit64(s string) uint64 {
	ss, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Println(err)

		// 输出: int64, -3546343826724305832
	}
	return uint64(ss)
}

func getQuotedStrings1(s string) []string {
	var re1 = regexp.MustCompile(`"(.*?)"`) // Note the matching group (submatch).
	ms := re1.FindAllStringSubmatch(s, -1)
	ss := make([]string, len(ms))
	for i, m := range ms {
		ss[i] = m[1]
	}
	return ss
}
