//go:build linux
// +build linux

package agent

func ExtendedFileInfo(fi *FileInfo, s *syscall.Stat_t) {
	fi.GID = s.Uid
	fi.UID = s.Uid
	fi.Device = s.Rdev
	fi.DeviceID = s.Dev
	fi.BlockSize = int64(s.Blksize)
	fi.Blocks = s.Blocks
	fi.AccessTime = s.Atim.Nano()
	fi.ModTime = s.Mtim.Nano()
	fi.ChangeTime = s.Ctim.Nano()
}
