package main

//Finna trying to make a somewhate relevant slice, pray for me
//Watch as I attempt to reinvent the wheel

//if yall want to try and mess aroudn with come shit just add some cool lookup functions or user input to these two files
//I believe in you, sorry I had to duck for the day but this would be really good practice working with slices
//If you dont feel ready to tackle it yet then no worries but feel free to mess wit it all you want
//I'll see you guys on Moday!
//Matthew

func idLookup(o []Order, s string) string {
	for _, id := range o {
		if s == id.ID {
			id.returnInfo()
			return "\n\nFound your order, inserted above.\n"
		} //end of if
	} //end of for range loop
	return "\nThere were no orders found with that ID.\n"
} //end of idLookup

func main() {
	var db []Order
	a := Order{"72APL6", "Verizon Edge 9", "Verizon", "Insert really long url here", 1.00, 2} //initializes an order
	db = append(db, a)                                                                        //adds a new order to the database
}
