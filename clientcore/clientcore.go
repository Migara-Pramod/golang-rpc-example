package clientcore

import (
	"fmt"
	"net/rpc"
	"strconv"
)

type Vegetable struct {
	Name string 
	Amount, Price float64
}

// GetAll methods returns available vegetable name list.
func GetAllVegetables(client *rpc.Client) {

	var vegList []string
	
	if err := client.Call("Vegetables.GetAll", "", &vegList); err != nil {
		fmt.Println("\nError | Vegetables.GetAll | ", err)
	} else {
		fmt.Printf("Available vegetables in store | '%s'\n", vegList)
	}
}

// Get price of one kilo of a vegetable
func GetPrice(client *rpc.Client) {

	name, _ := GetUserInput("Enter vegetable name: ", "s")
    	
    	var price float64
    	
        // get price
	if err := client.Call("Vegetables.PricePerKilo", name, &price); err != nil {
		fmt.Println("\nError | Vegetables.PricePerKilo |", err)
	} else {
		fmt.Printf("\nPrice of one kilo of '%s' | '%.2f'\n", name, price)
	}
}

// Get available amount of a vegetables
func GetAmount(client *rpc.Client) {

	name, _ := GetUserInput("Enter vegetable name: ", "s")
    			
    	var amount float64
    			
        // get amount
	if err := client.Call("Vegetables.Amount", name, &amount); err != nil {
		fmt.Println("\nError | Vegetables.Amount | ", err)
	} else {
		fmt.Printf("\nAvailable amount of '%s' in kilos | '%.3f'\n", name, amount)
	}
}

// Add a vegetable
func AddVegetable(client *rpc.Client) {
    	
    	name, _ := GetUserInput("Enter vegetable name: ", "s")		
    	_, price := GetUserInput("Enter vegetable price per kilo: ", "f")
    	_, amount := GetUserInput("Enter vegetable amount in kilos: ", "f")
    		
    	var veg Vegetable
    			    			    			    	
	//add a new vegetable
	if err := client.Call("Vegetables.Add", Vegetable{
		Name: name,
		Amount: amount,
		Price:  price,
	}, &veg); err != nil {
		fmt.Println("\nError | Vegetables.Add | ", err)
	} else {
		fmt.Printf("\nVegetable '%s' created\n", name)
	}
}

// Update a vegetable
func UpdateVegetable(client *rpc.Client) {

    	name, _ := GetUserInput("Enter vegetable name: ", "s")		
    	_, price := GetUserInput("Enter vegetable price per kilo: ", "f")
    	_, amount := GetUserInput("Enter vegetable amount in kilos: ", "f")
    			
    	var veg Vegetable

	//update a vegetable
	if err := client.Call("Vegetables.Update", Vegetable{
		Name: name,
		Amount: amount,
		Price:  price,
	}, &veg); err != nil {
		fmt.Println("\nError | Vegetables.Update | ", err)
	} else {
		fmt.Printf("\nVegetable '%s' updated\n", name)
	}    
}

//Get user input and validate
func GetUserInput(message string, typ string) (string, float64){
    	for {
    		var ins string
    		fmt.Println(message)
       	_, err := fmt.Scanln(&ins)
       	
       	if err != nil || ins == "" {
        		fmt.Println("!!! User input validation failed !!!")
        		continue
        	}
       	
       	//if output type is string return
       	//if output type is float execute next steps
       	if(typ == "s"){
       		return ins, 0
       	}
       	
       	var inf float64
       	inf, err = strconv.ParseFloat(ins, 64)
        	if err != nil {
        		fmt.Println("!!! User input validation failed !!!")
        	} else {
        		return "",inf
        	}
    	}
}
