package ostest

import "testing"

func TestCreateDirectory(t *testing.T) {
	err := createDir("/Users/xingzheng/go-playground/OS/test/test/test/")
	if err != nil {
		panic("error")
	}
}

func TestCreateAllDirectory(t *testing.T) {
	err := createAllDir("/Users/xingzheng/go-playground/OS/test/test/test/")
	if err != nil {
		panic("error")
	}
}

func TestReflect(t *testing.T) {
	reflcts("/Users/xingzheng/go-playground/OS/test/test/test/")
}
