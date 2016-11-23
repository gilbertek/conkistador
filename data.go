package main

var Items = map[int]*Item{
	1: {Name: "Key"},
	2: {Name: "Chext", ItemForUse: 1, Contains: []int{3}},
	3: {Name: "MEdal"},
}

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

var LocationMap = map[string]*Location{
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
		[]string{"Turbo Lift"},
		[]string{"alienAttack"}},
	"Lounge": {"You are in the lounge, you feel very relaxed",
		[]string{"Turbo Lift"},
		[]string{"relaxing"}},
}
var Enemies = map[int]*Character{
	1: {Name: "Klingon", Health: 50, Alive: true, Weap: 2},
	2: {Name: "Romulan", Health: 55, Alive: true, Weap: 3},
}

var Weaps = map[int]*Weapon{
	1: {Name: "Phaser", minAtt: 5, maxAtt: 15},
	2: {Name: "Klingon Disruptor", minAtt: 1, maxAtt: 15},
	3: {Name: "Romulan Disruptor", minAtt: 3, maxAtt: 12},
}
