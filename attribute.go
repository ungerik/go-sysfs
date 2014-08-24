package sysfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

type Attribute struct {
	Path string
	File *os.File
}

func (attrib *Attribute) Exists() bool {
	return fileExists(attrib.Path)
}

func (attrib *Attribute) Open() (err error) {
	attrib.File, err = os.OpenFile(attrib.Path, os.O_RDWR|syscall.O_NONBLOCK, 0660)
	return err
}

func (attrib *Attribute) Close() (err error) {
	err = attrib.File.Close()
	attrib.File = nil
	return err
}

func (attrib *Attribute) Ioctl(request, arg uintptr) (result uintptr, errno syscall.Errno) {
	if attrib.File == nil {
		attrib.Open()
		defer attrib.Close()
	}
	result, _, errno = syscall.Syscall(syscall.SYS_IOCTL, attrib.File.Fd(), request, arg)
	return result, errno
}

func (attrib *Attribute) Read() (string, error) {
	if attrib.File == nil {
		attrib.Open()
		defer attrib.Close()
	}
	attrib.File.Seek(0, os.SEEK_SET)
	data, err := ioutil.ReadAll(attrib.File)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (attrib *Attribute) Write(value string) error {
	if attrib.File == nil {
		attrib.Open()
		defer attrib.Close()
	}
	attrib.File.Seek(0, os.SEEK_SET)
	_, err := attrib.File.WriteString(value)
	return err
}

func (attrib *Attribute) Print(value interface{}) error {
	if attrib.File == nil {
		attrib.Open()
		defer attrib.Close()
	}
	attrib.File.Seek(0, os.SEEK_SET)
	_, err := fmt.Fprint(attrib.File, value)
	return err
}

func (attrib *Attribute) Scan(value interface{}) error {
	if attrib.File == nil {
		attrib.Open()
		defer attrib.Close()
	}
	attrib.File.Seek(0, os.SEEK_SET)
	_, err := fmt.Fscan(attrib.File, value)
	return err
}

func (attrib *Attribute) Printf(format string, args ...interface{}) error {
	if attrib.File == nil {
		attrib.Open()
		defer attrib.Close()
	}
	attrib.File.Seek(0, os.SEEK_SET)
	_, err := fmt.Fprintf(attrib.File, format, args...)
	return err
}

func (attrib *Attribute) Scanf(format string, args ...interface{}) error {
	if attrib.File == nil {
		attrib.Open()
		defer attrib.Close()
	}
	attrib.File.Seek(0, os.SEEK_SET)
	_, err := fmt.Fscanf(attrib.File, format, args...)
	return err
}

func (attrib *Attribute) ReadBytes() []byte {
	if attrib.File == nil {
		attrib.Open()
		defer attrib.Close()
	}
	attrib.File.Seek(0, os.SEEK_SET)
	data, _ := ioutil.ReadAll(attrib.File)
	return data
}

func (attrib *Attribute) WriteBytes(data []byte) error {
	if attrib.File == nil {
		attrib.Open()
		defer attrib.Close()
	}
	_, err := attrib.File.WriteAt(data, 0)
	return err
}

func (attrib *Attribute) ReadByte() (byte, error) {
	if attrib.File == nil {
		attrib.Open()
		defer attrib.Close()
	}
	data := make([]byte, 1)
	_, err := attrib.File.ReadAt(data, 0)
	return data[0], err
}

func (attrib *Attribute) WriteByte(value byte) error {
	return attrib.WriteBytes([]byte{value})
}

func (attrib *Attribute) ReadInt() (int, error) {
	s, err := attrib.Read()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

func (attrib *Attribute) WriteInt(value int) error {
	return attrib.Write(strconv.Itoa(value))
}
