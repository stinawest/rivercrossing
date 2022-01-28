package event

import "fmt"


/*
func Put(item string) string {
    return "[kylling rev korn ---\\ \\_korn_/ _________________/---]"
}


func ViewState() string {
	// Sjekke data som er lagret ("kylling til venstre", "rev til venstre")
	return "[kylling rev korn hs ---\\ \\__/ _________________/---]"
}




func CrossRiver() 
return //returnerer ny tilstand

func Reset()


func Println()


func main() {


	fmt.Println(item[0]) // getting results of the value
	fmt.Println(item[1])
	fmt.Println(item[2])
	fmt.Println(item[3])

	var kylling []string = item[0]
}

*/

func Start([]string) string {
	fmt.Println("Rivercrossing")
	fmt.Println("")
}


var item [4]string = [4]string { "kylling", "korn", "rev", "person"}
var kylling []string = item[:1]
var korn []string = item[1:2]
var rev []string = item[2:3]
var person []string = item[3:4]


func PutKy([]string) string {
	return "[kylling rev korn ---\\ \\_kylling_/ ______________/---]"
}

func PutKo([]string) string {
	return "[kylling rev korn ---\\ \\_korn_/ _________________/---]"
}

//må ha "put" funksjon for alle items, men for å demonstrere, har jeg valgt å ikke gøre det



/*
func ViewState() string {
	return "[kylling rev korn person ---\\ \\__/ _________________/---]"
}

func Takeout(item) string {

}

func crossriver()

func takeout(item)


func setup()

func reset()
*/