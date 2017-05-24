package datasipper

import (
	"encoding/json"
	"fmt"
)

func main() {
	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]

	//arg := os.Args[3]
	//fmt.Println(argsWithProg)
	//fmt.Println(argsWithoutProg)
	//fmt.Println(arg)

	db := DefaultConfig()
	j, err := db.ExecuteQuery("SELECT * FROM table")
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(&j, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}
