package hidden

import "syscall"

func IsHidden(path string) (bool, error) {
	p, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return false, err
	}
	attrs, err := syscall.GetFileAttributes(p)
	if err != nil {
		return false, err
	}
	return attrs&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}
