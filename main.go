package main

import (
	"os"
)

func main() {
	file, _ := os.OpenFile("0", os.O_RDWR|os.O_CREATE, 0755)
	file.Seek(0, 2)
	file.Write([]byte("haa"))
	file.Sync()
	file.Close()
}
