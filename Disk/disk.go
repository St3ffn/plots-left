package Disk

import (
	"syscall"
)

const (
	Byte = 1
	KiB  = 1024 * Byte
	MiB  = 1024 * KiB
	GiB  = 1024 * MiB
	TiB  = 1024 * GiB
)

// Disk represents abstraction of a disk
type Disk struct {
	// Path represents a directory path
	Path string
	// Total represents the total disk space
	Total Space
	// Free represents the free/available disk space
	Free Space
	// Used represents the used disk space
	Used Space
}

// Space represents disk space in bytes
type Space uint64

// Byte returns the disk Space in Bytes
func (s Space) Byte() uint64 {
	return uint64(s)
}

// Kibibyte returns the disk Space in Kibibyte
func (s Space) Kibibyte() float64 {
	return float64(s) / float64(KiB)
}

// Mebibyte returns the disk Space in Megibyte
func (s Space) Mebibyte() float64 {
	return float64(s) / float64(MiB)
}

// Gibibyte returns the disk Space in Gibibyte
func (s Space) Gibibyte() float64 {
	return float64(s) / float64(GiB)
}

// Tebibyte returns the disk Space in Gibibyte
func (s Space) Tebibyte() float64 {
	return float64(s) / float64(TiB)
}

// statfs represents the syscall to get the Statfs information
var statfs = syscall.Statfs

// NewDisk creates a new Disk from the given path.
// Details about the disk space will be fetched via statfs.
func NewDisk(path string) (*Disk, error) {
	fs := syscall.Statfs_t{}
	err := statfs(path, &fs)
	if err != nil {
		return nil, err
	}

	var disk Disk
	disk.Path = path
	disk.Total = Space(fs.Blocks * uint64(fs.Bsize))
	disk.Free = Space(fs.Bfree * uint64(fs.Bsize))
	disk.Used = disk.Total - disk.Free

	return &disk, nil
}
