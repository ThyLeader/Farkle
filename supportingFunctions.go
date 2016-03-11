package main

import (
        "fmt"
	    "log" 
        "github.com/faith/color"
		"time"
		//"net/http"
)

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

func printRules() {
    var isDone string
    d := color.New(color.FgRed, color.Bold)
    ruleArray := []string{"At the beginning of each turn, the player throws all the dice at once from a cup.", "After each throw, one or more scoring dice must be set aside.", "The player may then either end their turn and bank the score accumulated so far, or continue to throw the remaining dice.", "If the player has scored all six dice, they have \"hot dice\" and may continue their turn with a new throw of all six dice, adding to the score they have already accumulated. There is no limit to the number of \"hot dice\" a player may roll in one turn.", "If none of the dice score in any given throw, the player has \"farkled\" and all points for that turn are lost.", "At the end of the player's turn, the dice are handed to the next player in succession (usually in clockwise rotation), and they have their turn."}
    
    for _, e := range ruleArray {
        d.Println(e)
        time.Sleep(150 * time.Millisecond)
    }
    d.Println("\nType anything when you have finished reading the rules to continue")
    if _, err := fmt.Scanln(&isDone); err != nil {
        log.Fatal(err)
    }
}

func printScoring() {
    //d := color.New(color.FgRed, color.Bold)
    fmt.Println(`
==============================
| Dice combination |  Score  |
|    Each 1        |  100    |
|    Each 5        |  50     |
|    Three 1s      |  1000   |
|    Three 2s      |  200    |
|    Three 3s      |  300    |
|    Three 4s      |  400    |
|    Three 5s      |  500    | 
|    Three 6s      |  600    |
==============================`)
}