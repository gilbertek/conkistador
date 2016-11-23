package main

import "errors"

type Item struct {
	Name		string
	Action		string
	ItemForUser int
	Contains	[]int
}

func FindItemByName(name string) (error, int, *Item) {
	for index, _, itm := range Items {
		if itm.Name == name {
			return nil, index, itm
		}
	}

	return errors.New("Item not found!"), -1, nil
}

func OpenItem(plyr *Character, itemName string) {
	loc := LocationMap[plyr.CurrentLocation]
}
