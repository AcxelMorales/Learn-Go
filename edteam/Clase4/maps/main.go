package main

import "fmt"

func main() {
	// Si es en una línea, se omite la coma final
	// languages := make(map[string]string){
	// 	"es": "Español",
	// 	"en": "Inglés",
	// 	"fr": "Frances",
	// }

	languages := make(map[string]string)

	languages["es"] = "Español"
	languages["en"] = "Inglés"
	languages["fr"] = "Frances"

	fmt.Println(languages)

	delete(languages, "fr")

	fmt.Println(languages)

	if lng, ok := languages["pt"]; ok {
		fmt.Println("Exist", lng)
	} else {
		fmt.Println("No")
	}

}
