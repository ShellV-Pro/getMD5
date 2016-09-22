package main

import "flag"
import "fmt"
import "io"
import "os"
import "crypto/md5"

var lstring *string = flag.String("s", "nil", "string to md5")
var lfile *string = flag.String("f", "nil", "get file md5")

func main() {
	flag.Parse()

	if *lfile != "nil" {

		Filestomd5 := *lfile
		infile, inerr := os.Open(Filestomd5)
		if inerr == nil {
			md5h := md5.New()
			io.Copy(md5h, infile)
			fmt.Printf("%x", md5h.Sum([]byte("")))
		}
	}

	if *lstring != "nil" {

		Stringtomd5 := *lstring
		md5h := md5.New()
		md5h.Write([]byte(Stringtomd5))
		Result := md5h.Sum([]byte(""))
		fmt.Printf("%x", Result)
	}

	if *lstring == "nil" && *lfile == "nil" {
		fmt.Printf("===============================================\n")
		fmt.Printf("getMD5 MD5 get tool (一个可以获取文件和字符串MD5值的工具)\n")
		fmt.Printf("By ShellV \n")
		fmt.Printf("www.linux.dog \n")
		fmt.Printf("在Github上寻找这个项目: https://github.com/ShellV-Pro/getMD5 \n")
		fmt.Printf("QQ: 1657297533 E-mail: i@linux.dog 欢迎喜欢这款工具的小伙伴们加我O(∩_∩)O~~\n\n")
		fmt.Printf("帮助(⊙o⊙): ===================================\n")
		fmt.Printf("-f     \\\\ 你可以用这个参数来获取文件哒MD5值 \n")
		fmt.Printf("-s     \\\\ 你可以用这个参数来获取字符串哒MD5值 \n\n")
		fmt.Printf("举个栗子: =======================================\n")
		fmt.Printf("getMD5.exe -f acgbag.html \\\\ 获取了 acgbag.html 文件哒MD5 \n")
		fmt.Printf("getMD5.exe -s iloveyou-acgbag \\\\ 获取了 \"iloveyou-acgbag\" 这段字符串哒MD5 \n")
		fmt.Printf("getMD5.exe -f acgbag.html>>abc.txt \\\\ 获取了 acgbag.html 文件哒MD5 然后写入 abc.txt \n")
	}

}
