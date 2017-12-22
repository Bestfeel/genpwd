package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"time"
	"crypto/sha256"
	"crypto/sha512"
	"runtime"
	"os/exec"
	"flag"
)

const (
	COLOR_RED     = uint8(iota + 91)
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
	COLOR_MAGENTA
)

func MD5Pwd() string {
	m5 := md5.New()
	m5.Write([]byte(time.Now().String()))
	sum := m5.Sum(nil)
	return hex.EncodeToString(sum)
}

func Sha256Pwd() string {
	hash := sha256.New()
	hash.Write([]byte(time.Now().String()))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}

func Sha512Pwd() string {
	hash := sha512.New()
	hash.Write([]byte(time.Now().String()))
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}

func genPwd2Mac(len int) {

	if runtime.GOOS == "darwin" {
		out256, _ := exec.Command("openssl", "rand", "-base64", "256").Output()
		fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_MAGENTA, string(out256)[:len]+"\n")
		outShaSum, _ := exec.Command("bash", "-c", "date +%s |shasum |base64").Output()
		fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_MAGENTA, string(outShaSum)[:len]+"\n")
		outUrandom, _ := exec.Command("bash", "-c", "cat /dev/urandom | head -1 | md5 |base64 ").Output()
		fmt.Printf("\x1b[%d;1m%s\x1b[0m", COLOR_MAGENTA, string(outUrandom)[:len]+"\n")
	}
}

var (
	len = flag.Int("l", 32, "输入生成的密码长度")
)

func main() {
	flag.Parse()
	fmt.Printf("\x1b[%d;1mmd5:%s\x1b[0m", COLOR_MAGENTA, MD5Pwd()[:*len]+"\n")
	fmt.Printf("\x1b[%d;1msha256:%s\x1b[0m", COLOR_MAGENTA, Sha256Pwd()[:*len]+"\n")
	fmt.Printf("\x1b[%d;1msha512:%s\x1b[0m", COLOR_MAGENTA, Sha512Pwd()[:*len]+"\n")
	genPwd2Mac(*len)
}
