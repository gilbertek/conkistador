package main

import (
	"os"
	"strings"
)

func ProcessCommands(player *Character, input string) {
	Output("yellow", "===============================")

	tokens := strings.Fields(input)
	tokensLength := len(tokens)

	if tokensLength == 0 {
		Output("red", "No command received.")
		return
	}

	command := strings.ToLower(tokens[0])
	itemName := ""
	if tokensLength > 1 {
		itemName = tokens[1]
	}

	loc := LocationMap[player.CurrentLocation]
	switch command {
	case "go":
		fallthrough
	case "goto":
		if loc.CanGoTo(strings.ToLower(itemName)) {
			locName, err := FindLocationName(strings.ToLower(itemName))

			if err != nil {
				Output("red", "Can't go to "+itemName+" from here!")
			} else {
				player.CurrentLocation = locName
			}
		} else {
			Output("red", "Can't go to "+itemName+" from here!")
		}
	case "get":
		err, index, itm := FindItemByName(itemName)
		if err == nil && itm.ItemInRoom(loc) && !itm.ItemOnPlayer(player) {
			player.Items = append(player.Items, index)
			itm.RemoveItemFromRoom(loc)
		} else {
			Output("red", "Could not get "+itemName)
		}

	case "open":
		OpenItem(player, itemName)
	case "inv":
		Output("yellow", "Your Inventory: ")
		for _, itm := range player.Items {
			Output("yellow", "\t"+Items[itm].Name)
		}
	case "help":
		Output("blue", "Commands:")
		Output("blue", "\tgo <Location Name> - Move to the new location")
		Output("blue", "\tattack - Attack opponent(s)")
		Output("blue", "\tblock - Block incoming attack")
		Output("blue", "\trun - Escape attack")
		Output("blue", "\tget <Item Name> - Pick up item")
		Output("blue", "\topen <Item Name> - Open an item if it can be opened")
		Output("blue", "\tinv - Show what you are carrying\n\n")
	case "quit":
		Output("green", "Goodbye...")
		os.Exit(0)
	default:
	}
}
