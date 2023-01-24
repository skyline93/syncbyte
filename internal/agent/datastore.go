package agent

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"syscall"

	"github.com/pierrec/lz4/v4"
	log "github.com/sirupsen/logrus"
)

var blockSize int64 = 1024 * 1024

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

type Block struct {
	Data []byte
}

func (b *Block) Compress() ([]byte, error) {
	var (
		compressedBuf                  = make([]byte, lz4.CompressBlockBound(len(b.Data)))
		c             lz4.CompressorHC = lz4.CompressorHC{Level: lz4.Level2}
	)

	n, err := c.CompressBlock(b.Data, compressedBuf)
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return b.Data, nil
	}

	if n >= len(b.Data) {
		return b.Data, nil
	}

	return compressedBuf[:n], nil
}

func (b *Block) Hash() string {
	h := md5.New()
	h.Write(b.Data)
	return hex.EncodeToString(h.Sum(nil))
}

func UploadBigFile(filename string, callback func(string, []byte) error, fileInfo *FileInfo) error {
	var (
		buf       []byte = make([]byte, blockSize)
		partIndex int    = 1
		partInfos []*PartInfo
	)

	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fhReader := NewHashReader(file)

	for {
		rSize, err := fhReader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		block := Block{Data: buf[:rSize]}

		compressedBuf, err := block.Compress()
		if err != nil {
			return err
		}

		hash := block.Hash()

		log.Debugf("buf size: %d, compressedBuf size: %d", rSize, len(compressedBuf))

		if err := callback(hash, compressedBuf); err != nil {
			return err
		}

		pi := &PartInfo{Index: partIndex, MD5: hash, Size: int64(rSize)}
		partInfos = append(partInfos, pi)

		partIndex += 1
	}

	fileInfo.Name = fi.Name()
	fileInfo.Size = fi.Size()
	fileInfo.Path = filename
	fileInfo.MD5 = fhReader.Hash()
	fileInfo.PartInfos = partInfos

	s := fi.Sys().(*syscall.Stat_t)
	ExtendedFileInfo(fileInfo, s)

	return nil
}

func UploadSmallFile(filename string, callback func(string, []byte) error, fileInfo *FileInfo) error {
	var buf []byte = make([]byte, blockSize)

	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	hReader := NewHashReader(file)

	rSize, err := hReader.Read(buf)
	if err != nil {
		if err != io.EOF {
			return err
		}
	}

	block := Block{Data: buf[:rSize]}

	compressedBuf, err := block.Compress()
	if err != nil {
		return err
	}

	hash := block.Hash()
	log.Debugf("buf size: %d, compressedBuf size: %d", rSize, len(compressedBuf))

	if err := callback(hash, compressedBuf); err != nil {
		return err
	}

	fileInfo.Name = fi.Name()
	fileInfo.Size = fi.Size()
	fileInfo.Path = filename
	fileInfo.MD5 = hReader.Hash()

	s := fi.Sys().(*syscall.Stat_t)
	ExtendedFileInfo(fileInfo, s)

	return nil
}

type DataStore interface {
	UploadFile(filename string, fileInfo *FileInfo) error
}

type NasVolume struct {
	MountPoint string
}

func NewNasVolume() *NasVolume {
	return &NasVolume{MountPoint: Conf.Storage.NASVolumeMountPoint}
}

func (n *NasVolume) UploadFile(filename string, fileInfo *FileInfo) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}

	if fi.Size() >= blockSize {
		return UploadBigFile(filename, n.uploadFile, fileInfo)
	}

	return UploadSmallFile(filename, n.uploadFile, fileInfo)
}

func (n *NasVolume) uploadFile(filename string, buf []byte) error {
	file, err := os.Create(filepath.Join(n.MountPoint, filename))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(buf))
	if err != nil {
		return err
	}

	return nil
}
