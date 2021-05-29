package diskInfoStruct

// diskInfoWin 系统的一条磁盘记录
type diskInfoWin struct {
	VolumeId int    `json:"volume_id"` //卷标
	Ltr      string `json:"ltr"`       //盘符
	Label    string `json:"label"`     //名称
	Fs       string `json:"fs"`        //文件系统类型 NTFS是现代文件系统默认的，FAT32是老旧的32位二进制管理文件的方式
	Type_    string `json:"type_"`     //分区类型partition代表分区类型
	Size_    uint64 `json:"size_"`     //磁盘大小
	Status   string `json:"status"`    //磁盘健康状态
	Info_    string `json:"info"`      //根目录、系统、隐藏
}

// AllDiskInfoWin 所有的Linux磁盘信息
type AllDiskInfoWin struct {
	Name_ string      `json:"name"`
	Info_ diskInfoWin `json:"info"`
}

// DiskInfoLinux 的一条记录
type DiskInfoLinux struct {
	Name       string `json:"name"`        //块设备名
	Size_      uint64 `json:"size_"`       //容量大小
	Type       string `json:"type"`        //是否是磁盘上的一个分区sda/sdb是磁盘sr0是只读存储rom
	MountPoint string `json:"mount_point"` //mountpoint挂载点
	PkName     string `json:"pk_name"`     //
	FsType     string `json:"fs_type"`     //文件系统类型
}

// AllDiskInfoLinux 所有的win磁盘信息
type AllDiskInfoLinux struct {
	Name string        `json:"name"`
	Info DiskInfoLinux `json:"info"`
}
