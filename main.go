package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/fatih/color"
)

var Out *os.File
var In *os.File
var player Character

type Event struct {
	Type        string
	Chance      int
	Description string
	Health      int
	Evt         string
}

type Game struct {
	Welcome         string
	Health          int
	CurrentLocation string
}

func (e *Event) ProcessEvent() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	if e.Chance >= r1.Intn(100) {
		hp := e.Health
		if e.Type == "Combat" {
			fmt.Println("Combat Event")
		}

		fmt.Printf("\t%s\n", e.Description)
		if e.Evt != "" {
			hp = hp + evts[e.Evt].ProcessEvent()
		}
		return hp
	}
	return 0
}

func (g *Game) Play() {
	fmt.Println(g.Welcome)

	for {
		// Where are you
		fmt.Println(locationMap[g.CurrentLocation].Description)

		// Did anything happened here
		g.ProcessEvents(locationMap[g.CurrentLocation].Events)
		if g.Health <= 0 {
			fmt.Println("You are dead, GAME OVER!!!")
			return
		}

		// Print current health information
		fmt.Printf("Health: %d\n", g.Health)
		fmt.Println("You can go to these palces:")

		for index, loc := range locationMap[g.CurrentLocation].Transitions {
			fmt.Printf("\t%d - %s\n", index+1, loc)
		}

		i := 0
		for i < 1 || i > len(locationMap[g.CurrentLocation].Transitions) {
			// What would you like to do?
			fmt.Printf("%s%d%s\n", "Where do you want to go (0 - to quit), [1...",
				len(locationMap[g.CurrentLocation].Transitions), "]: ")

			fmt.Scan(&i)
		}

		newLoc := i - 1
		g.CurrentLocation = locationMap[g.CurrentLocation].Transitions[newLoc]
	}
}

func (g *Game) ProcessEvents(events []string) {
	for _, evtName := range events {
		g.Health += evts[evtName].ProcessEvent()
	}
}

func RunBattle(players Players) {
	sort.Sort(players)

	round := 1
	numALive := players.Len()
	for {
		DisplayInfo("Combat round ", round, "begin...")

		for x := 0; x < players.Len(); x++ {
			if players[x].Alive != true {
				continue
			}

			tgt := selectTarget(players, x)

			if tgt != -1 {
				DisplayInfo("Player: ", x, "Target: ", tgt)
				attp1 := players[x].Attack()
				players[tgt].Health = players[tgt].Health - attp1

				if players[tgt].Health <= 0 {
					players[tgt].Alive = false
					numALive--
				}

				DisplayInfo(players[x].Name+"attacks and does ",
					attp1,
					"points of damage with his",
					Weaps[players[x].Weap].Name,
					"to the enemy.")

			}
		}
		if endBattle(players) {
			break
		} else {
			DisplayInfo(players)
			round++
		}
	}
}

func endBattle(players []Character) bool {
	count := make([]int, 2)
	count[0] = 0
	count[1] = 0

	for _, playr := range players {
		if playr.Alive {
			if playr.Npc == false {
				count[0]++
			} else {
				count[1]++
			}
		}
	}

	if count[0] == 0 || count[1] == 0 {
		return true
	}

	return false
}

func selectTarget(players []Character, x int) int {
	y := x
	for {
		y = y + 1
		if y >= len(players) {
			y = 0
		}

		if (players[y].Npc != players[x].Npc) && players[y].Alive {
			return y
		}

		if y == x {
			return -1
		}
	}
	return -1
}

func DisplayInfoInterface(format string, args ...interface{}) {
	fmt.Fprintf(Out, format, args...)
}

func DisplayInfo(args ...interface{}) {
	fmt.Fprintln(Out, args...)
}

func UserInput(i *int) {
	fmt.Fscan(In, i)
}

func UserInputln() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n >>> ")
	text, _ := reader.ReadString('\n')
	return text
}

func init() {
	Out = os.Stdout
	In = os.Stdin
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	// Players
	player := *new(Character)
	player.Name = "Paul"
	player.Speed = 1 + rand.Intn(100)
	player.Health = 100
	player.Alive = true
	player.Weap = 1
	player.CurrentLocation = "Bridge"
	player.Play()
}

func Output(c string, args ...interface{}) {
	s := fmt.Sprint(args...)
	colr := color.WhiteString
	switch c {
	case "green":
		colr = color.GreenString
	case "red":
		colr = color.RedString
	case "blue":
		colr = color.BlueString
	case "yellow":
		colr = color.YellowString
	}
	fmt.Fprintln(Out, colr(s))
}

func Outputf(c string, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	Output(c, s)
}
