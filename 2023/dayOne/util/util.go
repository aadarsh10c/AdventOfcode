package util

import(
	"log"
	"os"
	"strings"
)


//Takes in a txt file ,returns content as slice of string 
func ReadInput(fileName string) []string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	contentStr := string(content)
	//remove \r\n
	contentStrList := strings.Split(contentStr, "\r\n")

	
	return contentStrList
} 