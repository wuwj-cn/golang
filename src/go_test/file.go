package file

import (
	"syscall"
)

type File struct {
	fd int
	name string
}

var (
	Stdin = newFile(0, "/dev/stdin")
	Stdout = newFile(1, "/dev/stdout")
	Stderr = newFile(2, "/dev/stderr")
)
func newFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name} }

func Open(name string, mode int, perm uint32) (file *File, err error) {
	r, e := syscall.Open(name, mode, perm)
	if e != nil {
		err = e
	}
	return newFile(r, name), err
}

func (file *File) Close() error {
	if file == nil {
		return syscall.EINVAL
	}
	e := syscall.Close(file.fd)
	if e != nil {
		return e
	}
	return nil
}

func (file *File) Read(b []byte) (ret int, err error) {
	if file == nil {
		return -1, syscall.EINVAL
	}
	r, e := syscall.Read(file.fd, b)
	if e != nil {
		err = e
	}
	return int(r), err
}

func (file *File) Write(b []byte) (ret int, err error) {
	if file == nil {
		return -1, syscall.EINVAL
	}
	r, e := syscall.Write(file.fd, b)
	if e != nil {
		err = e
	}
	return int(r), err
}

func (file *File) String() string {
	return file.name
}
