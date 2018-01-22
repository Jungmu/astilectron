package main

import (
	"encoding/json"
	"strconv"
	"fmt"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	// var result float64
	
	switch m.Name {
	case "sendNum":
		var s string
		if err := json.Unmarshal(m.Payload, &s); err != nil {
			astilog.Error(errors.Wrap(err, "json.Unmarshal failed"))
		}
		num, err := strconv.ParseFloat(s, 64)
		if err == nil {
			numbers = append(numbers,num)
		}
	case "sendSymbol":
		var s string
		if err := json.Unmarshal(m.Payload, &s); err != nil {
			astilog.Error(errors.Wrap(err, "json.Unmarshal failed"))
		}
		if err == nil {
			symbols = s
		}
	case "clear":
		numbers = numbers[:0]
		symbols = ""
	case "result":
		switch symbols {
		case "+":
			fmt.Println(numbers[0])
			fmt.Println(numbers[1])
			result := numbers[0] + numbers[1]
			bootstrap.SendMessage(w, "resResult", strconv.FormatFloat(result,'f',6,64)) 
		case "-":
			fmt.Println(numbers[0])
			fmt.Println(numbers[1])
			result := numbers[0] - numbers[1]
			bootstrap.SendMessage(w, "resResult", strconv.FormatFloat(result,'f',6,64)) 
		case "x":
			fmt.Println(numbers[0])
			fmt.Println(numbers[1])
			result := numbers[0] * numbers[1]
			bootstrap.SendMessage(w, "resResult", strconv.FormatFloat(result,'f',6,64)) 
		case "/":
			fmt.Println(numbers[0])
			fmt.Println(numbers[1])
			result := numbers[0] / numbers[1]
			bootstrap.SendMessage(w, "resResult", strconv.FormatFloat(result,'f',6,64)) 
		}
		// for _,value := range numbers {
		// 	result += value
		// }
		// if err := bootstrap.SendMessage(w, "about", strconv.FormatFloat(result,'f',6,64), func(m *bootstrap.MessageIn) {
		
		// }); err != nil {
		// }
	default:
	}
	return
}
