package main

import (
	"fmt"
	"Hang" 
)


func main() {
	var lancer int
	fmt.Println("\n\nHangman")
	fmt.Println("/nChoisissez le mode de jeu")
	fmt.Println("1 . Mot au hasard.")
	fmt.Println("0 . Quitter.")
	fmt.Scanln(&lancer)
	switch lancer {
	case 1:
		Hang.Hasard()
    case 0: 
        main()
	default:
		fmt.Println("chois")
		return
	}
}