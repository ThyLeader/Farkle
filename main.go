package main

import (
        "fmt"
	    "log" 
        "github.com/faith/color"
		"math/rand"
		"time"
        "strings"
        "sort"
		//"net/http"
)

func main() {
    //d := color.New(color.FgRed, color.Bold)
    mode := printWelcome()
    if mode == 1{
        singlePlayer()
    }
}

func rollDice(len int) ([]int, int) {
    toReturn := make([]int, len) 
    for i, _ := range toReturn {
        r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
        d := r.Intn(6) + 1
        toReturn[i] = d
    }
    return toReturn, len
}

func singlePlayer() {
    var mode string
    d := color.New(color.FgRed, color.Bold)
    d.Println("Singleplayer initiated")
    d.Println("Type \"Rules\" to see the rules, otherwise type anything else")
    if _, err := fmt.Scanln(&mode); err != nil {
        log.Fatal(err)
        d.Println(mode)
    }
    if strings.ToLower(mode) == "rules" {
        printScoring()
        printRules()
    }
    dice, amt := rollDice(6)
    sort.Ints(dice)
    d.Println("\nYou rolled", amt, "dice")
    d.Println(dice)
}
