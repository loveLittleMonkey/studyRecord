package try_test

import "testing"

func TestFirstTry(t *testing.T) {
	t.Log("my first try!")
}

/*
测试程序
源码文件以 _test结尾，xxx_test.go
测试方法以 Test开头，TestXXX(t *testing.T) {...}

go test first_test.go
*/