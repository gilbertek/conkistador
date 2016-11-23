package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

var Out *os.File
var In *os.File
var player Character

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
	player := new(Character)
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
