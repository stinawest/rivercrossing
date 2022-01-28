package event

import "fmt"



func Start() {
	fmt.Println("Dette er Rivercrossing")

}

// array av items som skal krysse elven
var item [4]string = [4]string { "kylling", "korn", "rev", "person"}

//slices av array
var kylling []string = item[:1]
var korn []string = item[1:2]
var rev []string = item[2:3]
var person []string = item[3:4]



func PutKyll([]string) string {
	return "[kylling rev korn person ---\\ \\_kylling_/ ______________/---]"
}

//må ha "put" funksjon for alle items, men dette er bare for å demonstrere




func CrossRiverFor() string {
	// "CrossRiverFor" -- Forward
	return "[kylling rev korn person ---\\____________________ \\____//---]"
}



//flere mulige events: func Takeout(item), func CrossRiverBck(), func setup(), func reset()
