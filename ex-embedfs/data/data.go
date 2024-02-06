package data

import (
	"embed"
	"fmt"
	"os"
)

//go:embed people
var peopleFile embed.FS

func GetPeopleFromEmbeddedFile() {

	peopleFile, _ := peopleFile.Open("people/people.json")
	peopleFileInfo, _ := peopleFile.Stat()
	fmt.Println("embedded people:", peopleFileInfo.Name(), peopleFileInfo.Size(), peopleFileInfo.ModTime())

	dirFs := os.DirFS("data/people")
	file, _ := dirFs.Open("people.json")
	info, _ := file.Stat()
	fmt.Println("embedded people file:", info.Name(), info.Size(), info.ModTime())
}
