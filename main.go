package main

import (
        "fmt"
	    "log" 
        "github.com/faith/color"
		"math/rand"
		"time"
		//"net/http"
)

func main() {
    mode := printWelcome()
    dice := rollDice(6)
    fmt.Println(dice, mode)
}

func printWelcome() int{
    mode := 0
    d := color.New(color.FgRed, color.Bold)
    d.Println("Welcome to Farkle, a fairly convoluted dice game that may be fun.")
    d.Println("Would you like to play against ai or a friend? (1 for by yourself, 2 for ai, 3 for friends):")
    d.Println("Hint: you can currently only play by yourself. Deal with it")
    for mode == 0 {
        if _, err := fmt.Scanln(&mode); err != nil {
            log.Fatal(err)
        }
        if mode == 1 {
            break
        } else {
            d.Println("Please choose 1")
            mode = 0
        }
    }

    return mode
}

func rollDice(len int) []int {
    toReturn := make([]int, len) 
    for i, _ := range toReturn {
        r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
        d := r.Intn(6) + 1
        toReturn[i] = d
    }
    return toReturn
}