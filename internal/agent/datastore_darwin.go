//go:build darwin

package agent

import "syscall"

func ExtendedFileInfo(fi *FileInfo, s *syscall.Stat_t) {
	fi.GID = s.Uid
	fi.UID = s.Uid
	fi.Device = uint64(s.Rdev)
	fi.DeviceID = uint64(s.Dev)
	fi.BlockSize = int64(s.Blksize)
	fi.Blocks = s.Blocks

	fi.AccessTime = s.Atimespec.Nano()
	fi.ModTime = s.Mtimespec.Nano()
	fi.ChangeTime = s.Ctimespec.Nano()
}
