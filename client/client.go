package main

import (
	"fmt"
	"net/rpc"
	"rpc.example.com/clientcore"
)

func main() {

	// get RPC client by dialing at `rpc.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`
	
	fmt.Println("\n******** Welcome to online vegetable store **********")
	
	Loop:for {
	
		fmt.Println("\nYou have following options\n* Press 'a' and Enter to Exit\n* Press 'b' and Enter to List all vegetables\n* Press 'c' and Enter to Get price of one kilo of a vegetable\n* Press 'd' and Enter to Get available amount of a vegetable in kilos\n* Press 'e' and Enter to Add a new vegetable\n* Press 'f' and Enter to Update a vegetable\n")
	
		fmt.Println("Enter Your Option: ")
    		var option string
    		fmt.Scanln(&option)
    		
    		fmt.Println("\n******************************************************")
    		
		switch option {
		case "a":
			//Exit
			break Loop
    		case "b":
        		// Get all vegetables
			clientcore.GetAllVegetables(client)
		
    		case "c":
        		// Get price of one kilo of a vegetable
			clientcore.GetPrice(client)
					
    		case "d":
        		// Get available amount of a vegetables
			clientcore.GetAmount(client)
			
    		case "e":
        		// Add a vegetable
			clientcore.AddVegetable(client)
			
    		case "f":
        		// Update a vegetable
			clientcore.UpdateVegetable(client)
			
		default:
			fmt.Println("!!! Invalid option !!!")
			
    		}
    		
    		fmt.Println("******************************************************")
	}

}
