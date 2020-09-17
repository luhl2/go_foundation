package file

import "os"

func CreateFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		println(err.Error())
	}
	f.WriteString("第一行\n,第二行\n")
	defer f.Close()
}
