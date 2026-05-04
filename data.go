package main

import (
	"fmt"
	"math/rand"
)

type ArtWork struct {
	Width    int
	Height   int
	Filename string
	Lines    []string
	Status   bool
	PosY     int
	MinDif   int
	Fuzz     int
}

type Engine struct {
	Buffer      []string
	Dimensions  []ArtWork
	FinalBuffer []string
	Columns     int
	Lines       int
	MaxLines    int
	Spacing     int
}

func (e *Engine) PlaceImages() bool {
	modified := false
	lastPrint := 0

	for i := range e.Dimensions {
		art := &e.Dimensions[i]
		status := false
		height := 0
		pos := 0
		minDif := 0
		startingPoint := lastPrint + 1 + e.Spacing

		for cursor, line := range e.Buffer {
			// Em Go, len(line) lida com caracteres, não precisamos de ghost_bytes do wc
			lineLen := len(line)

			if (e.Columns-lineLen > art.Width) && (cursor >= startingPoint) {
				height++
				if minDif < lineLen {
					minDif = lineLen
				}
				if height > art.Height && !status {
					status = true
					modified = true
					lastPrint = pos + art.Height + e.Spacing
				}
			} else if !status {
				pos = cursor
				minDif = lineLen
				height = 0
			}
		}

		art.Status = status
		art.PosY = pos
		art.MinDif = minDif
		if e.Columns-minDif-art.Width-1 > 0 {
			art.Fuzz = rand.Intn(e.Columns - minDif - art.Width - 1)
		}
	}
	return modified
}

func (e *Engine) ManipulateBuffer(frame int) {
	e.FinalBuffer = []string{}
	cursor := 0

	for _, art := range e.Dimensions {
		if !art.Status {
			continue
		}

		// Preenche as linhas antes da imagem
		for cursor < art.PosY {
			e.FinalBuffer = append(e.FinalBuffer, e.Buffer[cursor])
			cursor++
		}

		// Desenha a imagem com o efeito de frame (scroll-in)
		for cursor < (art.PosY + art.Height) {
			relativeLine := cursor - art.PosY
			threshold := art.Height - frame
			if threshold < 0 {
				threshold = 0
			}

			artLine := ""
			if relativeLine >= threshold {
				artLine = art.Lines[relativeLine-threshold]
			}

			// Montagem da linha: Buffer original + Padding + Arte
			padding := art.MinDif + art.Fuzz
			line := fmt.Sprintf("%-*s%s", padding, e.Buffer[cursor], artLine)
			e.FinalBuffer = append(e.FinalBuffer, line)
			cursor++
		}
	}

	// Adiciona o restante do buffer original
	for cursor < len(e.Buffer) {
		e.FinalBuffer = append(e.FinalBuffer, e.Buffer[cursor])
		cursor++
	}
}
