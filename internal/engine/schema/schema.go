package schema

import "encoding/json"

type PartInfo struct {
	Index int    `json:"index"`
	MD5   string `json:"md5"`
	Size  int64  `json:"size"`
}

type FileInfo struct {
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	GID        uint32 `json:"gid"`
	UID        uint32 `json:"uid"`
	Device     uint64 `json:"device"`
	DeviceID   uint64 `json:"device_id"`
	BlockSize  int64  `json:"block_size"`
	Blocks     int64  `json:"blocks"`
	AccessTime int64  `json:"atime"`
	ModTime    int64  `json:"mtime"`
	ChangeTime int64  `json:"ctime"`

	Path      string      `json:"path"`
	MD5       string      `json:"md5"`
	PartInfos []*PartInfo `json:"part_info"`
}

func (fi *FileInfo) String() string {
	v, _ := json.Marshal(fi)
	return string(v)
}
