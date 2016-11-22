package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Location struct {
	Description string
	Transitions []string
	Events      []string
}

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

type Character struct {
	Name    string
	Health  int
	Evasion int
	Alive   bool
	Speed   int
	Weap    int
	Npc     bool
}

type Players []Character

var Out *os.File
var In *os.File

var evts = map[string]*Event{
	"alienAttack": {
		Chance:      20,
		Description: "An alien beams in front of you and shoots you with a ray gun.",
		Health:      -50,
		Evt:         "doctorTreatment"},
	"doctorTreatment": {
		Chance:      10,
		Description: "The doctor rushes in and inject you with a health boost.",
		Health:      +30,
		Evt:         ""},
	"android": {
		Chance:      50,
		Description: "Data is in the turbo lift and says hi to you",
		Health:      0,
		Evt:         ""},
	"relaxing": {
		Chance:      100,
		Description: "In the lounge you are so relaxed that your health improves.",
		Health:      +10,
		Evt:         ""},
}

var locationMap = map[string]*Location{
	"Bridge": {
		"You are on the bridge of a spaceship sitting in the Captain's chair.",
		[]string{"Ready Room", "Turbo Lift"},
		[]string{"alienAttack"}},
	"Ready Room": {
		"The Captain's ready room.",
		[]string{"Bridge"},
		[]string{}},
	"Turbo Lift": {
		"A Turbo Lift that takes you anywhere in the ship.",
		[]string{"Bridge", "Lounge", "Engineering"},
		[]string{"android"}},
	"Engineering": {"You are in engineering where you see the star drive",
		[]string{"Turbo Lift"}, []string{"alienAttack"}},
	"Lounge": {"You are in the lounge, you feel very relaxed",
		[]string{"Turbo Lift"}, []string{"relaxing"}},
}
var enemies = map[int]*Character{
	1: {Name: "Klingon", Health: 50, Alive: true, Weap: 2},
	2: {Name: "Romulan", Health: 55, Alive: true, Weap: 3},
}

var Weaps = map[int]*Weapon{
	1: {Name: "Phaser", minAtt: 5, maxAtt: 15},
	2: {Name: "Klingon Disruptor", minAtt: 1, maxAtt: 15},
	3: {Name: "Romulan Disruptor", minAtt: 3, maxAtt: 12},
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

func (p *Character) Equip(w int) {
	p.Weap = w
}

func (p *Character) Attack() int {
	return Weaps[p.Weap].Fire()
}

func (slice Players) Len() int {
	return len(slice)
}

func (slice Players) Less(i, j int) bool {
	// Sort Descending
	return slice[i].Speed > slice[j].Speed
	// Sort ascending
	// return slice[i].Speed < slice[j].Speed
}

func (slice Players) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
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

func GetUserInput(i *int) {
	fmt.Fscan(In, i)
}

func init() {
	Out = os.Stdout
	In = os.Stdin
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	DisplayInfo("Welcome to Conkistador!")
	// Players
	p1 := new(Character)
	p1.Name = "Paul"
	p1.Speed = 1 + rand.Intn(100)
	p1.Health = 100
	p1.Alive = true
	p1.Weap = 1

	p2 := new(Character)
	*p2 = *enemies[1+rand.Intn(2)]
	p2.Npc = true
	p2.Speed = 1 + rand.Intn(100)

	players := Players{*p1, *p2}

	sort.Sort(players)

	DisplayInfo(players[0])
	DisplayInfo(players[1])
	DisplayInfo(players)

	round := 1

	g := &Game{Health: 100, Welcome: "Welcome to conkistdor!", CurrentLocation: "Bridge"}
	g.Play()
}
