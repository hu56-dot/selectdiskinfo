/*
 * @Author: your name
 * @Date: 2021-05-20 17:45:15
 * @LastEditTime: 2021-05-28 19:34:04
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \Go\src\systemcommand\bytesOutputHandle\handleStr_windows.go
 */
package bytesOutputHandle

import (
	"encoding/json"
	"fmt"

	"github.com/StackExchange/wmi"
)

type Storage struct {
	Name       string
	FileSystem string
	Total      uint64
	Free       uint64
	StatusInfo string
	VolumeName string
	//DeviceID   string
	//Caption    string
}

type storageInfo struct {
	Name       string
	Size       uint64
	FreeSpace  uint64
	FileSystem string
	StatusInfo string
	VolumeName string
	//DeviceID   string
	//Caption    string
}

func HandleString() string {
	var storageinfo []storageInfo
	var localStorages []Storage
	err := wmi.Query("Select * from Win32_LogicalDisk", &storageinfo)
	if err != nil {
		fmt.Println("查询error:", err)
	}

	for _, storage := range storageinfo {
		info := Storage{
			Name:       storage.Name,
			FileSystem: storage.FileSystem,
			Total:      storage.Size,
			Free:       storage.FreeSpace,
			StatusInfo: storage.StatusInfo,
			VolumeName: storage.VolumeName,
			//Caption:    storage.Caption,
			//DeviceID:	storage.DeviceID,
		}
		localStorages = append(localStorages, info)
	}

	byte, err := json.Marshal(localStorages)
	if err != nil {
		fmt.Println("解析错误")
	}
	//fmt.Printf("localStorages:=%s\n",byte)
	return string(byte)
}
