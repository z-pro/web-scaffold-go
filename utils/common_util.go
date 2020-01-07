package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"os"
	"path/filepath"
	"strings"
)

func EnWithMd5(orignalStr string) string  {
	h := md5.New()
	h.Write([]byte(orignalStr)) // 需要加密的字符串为 sharejs.com
	return hex.EncodeToString(h.Sum(nil))
}

/*
获取程序运行路径
*/
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Debug(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}