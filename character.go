package main

type Character struct {
	Name    string
	Health  int
	Evasion int
	Alive   bool
	Speed   int
	Weap    int
	Npc     bool
	Items   []int

	Welcome         string
	CurrentLocation string
}

type Players []Character

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

func (p *Character) Play() {
	Output(p.Welcome)

	for {
		Output("blue", LocationMap[p.CurrentLocation].Description)
		p.ProcessEvents(LocationMap[p.CurrentLocation].Events)

		if p.Health <= 0 {
			Output("white", "You are dead! GAME OVER!!!")
			return
		}

		Output("\nblue", "Health: ", p.Health)

		if len(LocationMap[p.CurrentLocation].Items) > 0 {
			Output("yellow", "You can see:")
			for _, itm := range LocationMap[p.CurrentLocation].Items {
				Outputf("yellow", "\t%s", Items[itm].Name)
			}
		}

		Output("green", "You can go to these places:")
		for _, loc := range LocationMap[p.CurrentLocation].Transitions {
			Outputf("yellow", "\t%s", loc)
		}

		cmd := UserInputln()
		ProcessCommands(p, cmd)
	}
}

func (p *Character) ProcessEvents(events []string) {
	for _, evnt := range events {
		p.Health += evts[evnt].ProcessEvent(p)
	}
}
