package main

import (
	"control/pkg/card"
	"control/pkg/types"
	"fmt"
)

func main() {
	temps := []*types.Card{
		{
		  	Balance: 16,
	     	Currency: "TJS",
		    Activity: false,
		},
		{
		 	Balance: 17,
	     	Currency: "TJS1",
		    Activity: false,
		},
	}    
     card.ChangeActiveCard(temps)	 
	 for i:=0; i<len(temps); i++{
         fmt.Println(*&temps[i].Activity)
	 }      	
}