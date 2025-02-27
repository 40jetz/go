package main
import (
	"fmt" 
)

func main() {

	var conferencetName string  = "Go conference"
	const conferenceTickets uint = 50 // uint is whole number (0 - >0)
	fmt.Printf("Conference : %v \n" , conferencetName)
	fmt.Printf("Total Tickets : %v \n", conferenceTickets) 

	// ! Data type : %T
	// fmt.Printf("conferenceTicket : %T \n" , conferenceTickets)
	
	var username string
	fmt.Println("Enter your username")
	fmt.Scan(&username)
	fmt.Println(username)
	// fmt.Println(&conferencetName) // &var (poiter/specialvar) gives memory location


}