package web

import (
	"os"
	"os/exec"
	"path"
	"reflect"
	"syscall"

	"github.com/satori/go.uuid"
)

//UUID generate uuid v4 string
func UUID() string {
	return uuid.NewV4().String()
}

//Shell 终端命令操作
func Shell(cmd string, args ...string) error {
	bin, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}
	return syscall.Exec(bin, append([]string{cmd}, args...), os.Environ())
}

//Package 获得包路径
func Package(o interface{}, ds ...string) string {
	v := []string{os.Getenv("GOPATH"), "src", reflect.TypeOf(o).Elem().PkgPath()}
	v = append(v, ds...)
	return path.Join(v...)
}
