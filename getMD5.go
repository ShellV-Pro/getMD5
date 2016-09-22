package main

import "flag"
import "fmt"
import "io"
import "os"
import "crypto/md5"

var lstring *string = flag.String("s", "", "string to md5")
var lfile *string = flag.String("f", "", "get file md5")

func main() {
	flag.Parse()

	if lfile != nil {
		Filestomd5 := *lfile
		infile, inerr := os.Open(Filestomd5)
		if inerr == nil {
			md5h := md5.New()
			io.Copy(md5h, infile)
			fmt.Printf("%x", md5h.Sum([]byte("")))
		}
	}

}
