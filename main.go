package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type GopherSay struct {
	art string
}

func NewGopherSay(artPath string) (*GopherSay, error) {
	art, err := os.ReadFile(artPath)
	if err != nil {
		return nil, err
	}
	return &GopherSay{art: string(art)}, nil
}

func (g *GopherSay) Say(message string) string {
	border := strings.Repeat("-", len(message)+2)
	return fmt.Sprintf(`
 %s 
< %s >
 %s 
%s`, border, message, border, g.art)
}

func getRandomGopherArt(directory string) (string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return "", err
	}

	var gopherArts []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".goph") {
			gopherArts = append(gopherArts, directory+"/"+file.Name())
		}
	}

	if len(gopherArts) == 0 {
		return "", fmt.Errorf("no .goph files found in directory")
	}

	rand.Seed(time.Now().UnixNano())
	return gopherArts[rand.Intn(len(gopherArts))], nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gophersay <message>")
		os.Exit(1)
	}

	message := os.Args[1]
	artPath, err := getRandomGopherArt("gopher")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	gopher, err := NewGopherSay(artPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(gopher.Say(message))
}
