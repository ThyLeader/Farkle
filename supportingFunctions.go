package main

import (
        "fmt"
	    "log" 
        "github.com/fatih/color"
		"time"
        "math/rand"
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

func rollDice(len int) ([]int, int) {
    toReturn := make([]int, len) 
    for i := range toReturn {
        r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
        d := r.Intn(6) + 1
        toReturn[i] = d
    }
    return toReturn, len
}

func addAmt(amt int, p int) {
    fmt.Println("Adding", amt, "points to player", p)
}

func (c *KeepScore) givePoints(player int, points int, i int, e int) {
    switch player {
        default: fmt.Println("KEK")
        case 1:
            c.p1 = c.p1 + points
            fmt.Println("Added", points, "points")
            fmt.Println("You scored", i, e)
            
        case 2:
            c.p2 = c.p2 + points
        case 3:
            c.p3 = c.p3 + points
        case 4:
            c.p4 = c.p4 + points
        case 5:
            c.p5 = c.p5 + points
        case 6:
            c.p6 = c.p6 + points
    }
}

func (c *KeepScore) calcPoints(dice []int, player bool) {
    s1, s2, s3, s4, s5, s6 := 0, 0, 0, 0, 0, 0
    before := c.p1
    for _, e := range dice {
        switch e {
            default: fmt.Println("KEK")
            case 1: s1++
            case 2: s2++
            case 3: s3++
            case 4: s4++
            case 5: s5++
            case 6: s6++
        }
        //fmt.Println(s1, s2, s3, s4, s5, s6)
    }
    num := []int{s1, s2, s3, s4, s5, s6}
    for i, e := range num {
        i++
        if i == 1 {
            switch e {
                default: fmt.Println("No points scored")
                case 1:
                    c.givePoints(1, 100, i, e)
                case 2:
                    c.givePoints(1, 200, i, e)
                case 3:
                    c.givePoints(1, 1000, i, e)
                case 4:
                    c.givePoints(1, 2000, i, e)
                case 5:
                    c.givePoints(1, 3000, i, e)
                case 6:
                    c.givePoints(1, 4000, i, e)
            }
        }
        if i == 2 {
            switch e {
                default: fmt.Println("No points scored")
                case 3:
                    c.givePoints(1, 200, i, e)
                case 4:
                    c.givePoints(1, 400, i, e)
                case 5:
                    c.givePoints(1, 600, i, e)
                case 6:
                    c.givePoints(1, 800, i, e)
            }
        }
        if i == 3 {
            switch e {
                default: fmt.Println("No points scored")
                case 3:
                    c.givePoints(1, 300, i, e)
                case 4:
                    c.givePoints(1, 600, i, e)
                case 5:
                    c.givePoints(1, 900, i, e)
                case 6:
                    c.givePoints(1, 1200, i, e)
            }
        }
        if i == 4 {
            switch e {
                default: fmt.Println("No points scored")
                case 3:
                    c.givePoints(1, 400, i, e)
                case 4:
                    c.givePoints(1, 800, i, e)
                case 5:
                    c.givePoints(1, 1200, i, e)
                case 6:
                    c.givePoints(1, 1600, i, e)
            }
        }
        if i == 5 {
            switch e {
                default: fmt.Println("No points scored")
                case 1:
                    c.givePoints(1, 50, i, e)
                case 2:
                    c.givePoints(1, 100, i, e)
                case 3:
                    c.givePoints(1, 500, i, e)
                case 4:
                    c.givePoints(1, 1000, i, e)
                case 5:
                    c.givePoints(1, 1500, i, e)
                case 6:
                    c.givePoints(1, 2000, i, e)
            }
        }
        if i == 6 {
            switch e {
                default: fmt.Println("No points scored")
                case 3:
                    c.givePoints(1, 600, i, e)
                case 4:
                    c.givePoints(1, 1200, i, e)
                case 5:
                    c.givePoints(1, 1800, i, e)
                case 6:
                    c.givePoints(1, 2400, i, e)
            }
        }
    }
    c.turn++
    diff := c.p1 - before
    fmt.Println("You scored", diff, "points on round", c.turn)
}

func (c *KeepScore) printScore() {
    fmt.Println(c.p1/*, c.p2*/, c.turn)
}