package servercore

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

// Vegetable struct represents a vegetable.
type Vegetable struct {
	Name string 
	Amount, Price float64
}

// Vegetables struct represents a map of vegetable.
type Vegetables struct {
	store map[string]Vegetable // private
}

// GetAll methods returns available vegetable name list.
func (v *Vegetables) GetAll(payload string, vegNames *[]string) error {
	
	keys := make([]string, 0, len(v.store))
    	for k := range v.store {
        	keys = append(keys, k)
    	}
	
	// set vegetable List
	*vegNames = keys

	// return `nil` error
	return nil
}

// PricePerKilo methods returns price per kilo of a vegetable.
func (v *Vegetables) PricePerKilo(vegName string, price *float64) error {
	
	// Get vegetable in the store by name
	result, ok := v.store[vegName]

	// check if vegetable exists in the store
	if !ok {
		return fmt.Errorf(" Vegetable, '%s' is not available in store", vegName)
	}

	// set price
	*price = result.Price

	// return `nil` error
	return nil
}

// Amount methods returns available amount in kilos of a vegetable.
func (v *Vegetables) Amount(vegName string, amount *float64) error {
	
	// Get vegetable in the store by name
	result, ok := v.store[vegName]

	// check if vegetable exists in the store
	if !ok {
		return fmt.Errorf(" Vegetable, '%s' is not available in store", vegName)
	}

	// set amount
	*amount = result.Amount

	// return `nil` error
	return nil
}

//Add a new vegetable.
func (v *Vegetables) Add(newVeg Vegetable, createdVeg *Vegetable) error {

	// check if vegetable already exists in the store
	if _, ok := v.store[newVeg.Name]; ok {
		return fmt.Errorf("Vegetable with name '%s' already exists", newVeg.Name)
	}

	// add new vegetable to the store
	v.store[newVeg.Name] = newVeg
	
	f, err := os.OpenFile("vegetables.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	
	if err != nil {
        	log.Fatalf("failed to open vegetables.txt file")
    	}
    	
    	defer f.Close()
    	
    	price := fmt.Sprintf("%.2f", newVeg.Price)
    	amount := fmt.Sprintf("%.3f", newVeg.Amount)
    		
	n3, err := f.WriteString(newVeg.Name+";"+amount+";"+price+"\n")
	
	if err != nil {
        	log.Fatalf("failed to write to vegetables.txt file")
    	}
    	
    	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()
	
	// set reply value
	*createdVeg = newVeg

	// return `nil` error
	return nil
}

func (v *Vegetables) Update(veg Vegetable, updatedVeg *Vegetable) error {

	// check if vegetable not exists in the store
	if _, ok := v.store[veg.Name]; !ok {
		return fmt.Errorf("Vegetable with name '%s' doesn't exists", veg.Name)
	}

	// add new vegetable to the store
	v.store[veg.Name] = veg
	
	f, err := os.OpenFile("vegetables.txt", os.O_WRONLY|os.O_CREATE, 0600)
	
	if err != nil {
        	log.Fatalf("failed to open vegetables.txt file")
    	}
    	
    	defer f.Close()
	
	for k := range v.store {
	
		price := fmt.Sprintf("%.2f", v.store[k].Price)
    		amount := fmt.Sprintf("%.3f", v.store[k].Amount)
    	
        	n3, err := f.WriteString(v.store[k].Name+";"+amount+";"+price+"\n")
        	if err != nil {
        		log.Fatalf("failed to write to vegetables.txt file")
    		}
    		fmt.Printf("wrote %d bytes\n", n3)
    	}
	
	f.Sync()
	
	// set reply value
	*updatedVeg = veg

	// return `nil` error
	return nil
}

// NewVegetables function returns a new instance of Vegetables (pointer).
func NewVegetables() *Vegetables {

	file, err := os.Open("vegetables.txt")
  
    	if err != nil {
        	log.Fatalf("failed to open vegetables.txt file")
    	}
    	
    	scanner := bufio.NewScanner(file)
    	scanner.Split(bufio.ScanLines)
    	
    	vegetablesMap := make(map[string]Vegetable)
    	
    	for scanner.Scan() {
        	nextLine := strings.Split(scanner.Text(), ";")
        	price, _ := strconv.ParseFloat(nextLine[2], 64)
        	amount, _ := strconv.ParseFloat(nextLine[1], 64)
		vegetablesMap[nextLine[0]] = Vegetable{
			Name: nextLine[0],
			Amount: amount,
			Price: price,
		}
		//fmt.Println(vegetablesMap[nextLine[0]])
    	}
    	
    	file.Close()
    	
	return &Vegetables{
		store: vegetablesMap,
	}
}
