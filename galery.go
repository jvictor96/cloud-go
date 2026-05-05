package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type ArtWork struct {
	Width    int
	Height   int
	Filename string
	Content  []string
}

type Galery struct {
	ArtWorks []ArtWork
	Path     string
}

func (g *Galery) LoadArt() {
	g.ArtWorks = []ArtWork{}
	files, _ := os.ReadDir(g.Path)
	for _, file := range files {
		data, _ := os.ReadFile(fmt.Sprintf("%s%s", g.Path, file.Name()))
		content := strings.Split(string(data), "\n")
		g.ArtWorks = append(g.ArtWorks, ArtWork{
			Content:  content,
			Filename: file.Name(),
			Height:   len(content),
			Width:    utf8.RuneCountInString(content[0]),
		})
	}
}

func (g *Galery) Route(args []string) {
	command := args[0]
	if command == "add" {
		g.AddImage(args[1])
		return
	}
	if command == "remove" {
		g.RemoveImage(args[1])
		return
	}
	g.ListImages()
}

func (g *Galery) AddImage(Path string) {
	Name := strings.Split(Path, "/")
	os.Rename(Path, fmt.Sprintf("%s%s", g.Path, Name))
}

func (g *Galery) RemoveImage(Name string) {
	os.Remove(fmt.Sprintf("%s%s", g.Path, Name))
}

func (g *Galery) ListImages() {
	files, _ := os.ReadDir(g.Path)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
