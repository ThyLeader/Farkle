package main

import (
        "fmt"
	    "log" 
        "github.com/fatih/color"
		//"math/rand"
		//"time"
        "strings"
        "sort"
)

type KeepScore struct {
    p1 int
    p2 int
}

func main() {
    mode := printWelcome()
    if mode == 1{
        singlePlayer()
    }
}

func singlePlayer() {
    c := new(KeepScore)
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
    c.addPoints(dice, true)
    c.printScore()
}
