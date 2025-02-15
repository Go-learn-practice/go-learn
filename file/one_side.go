package file

import (
	"io"
	"log"
	"os"
	"path"
)

// 分片读取文件内容

func OneSideReadWriteToDest() {
	list := getFiles(sourceDir)
	for _, l := range list {
		_, name := path.Split(l)
		destFileName := destDir + "one-side/" + name
		// 文件写入
		OneSideReadWrite(l, destFileName)
	}
}

func OneSideReadWrite(srcName, destName string) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dst, err := os.OpenFile(destName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		dst.Write(buf[:n])
	}
}
