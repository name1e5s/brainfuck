package Brainfucker

import (
	"fmt"
	"log"
)

type fucker struct {
	Cells		[30000]byte
	Pointer		int
	Code		string
	Codepos		int
	Codelen		int
}

func NewFucker(code string) *fucker {
	var cell [30000]byte
	fucker := fucker{cell,0,code,0,len(code)}
	return &fucker
}

func (fucker *fucker)Check() {
	count := 0
	for i := 0 ; i < fucker.Codelen; i++ {
		if fucker.Code[i] == '[' {
			count ++
		}
		if fucker.Code[i] == ']' {
			count --
		}
	}
	if count != 0 {
		log.Fatal("Check - ERROR: Mismatch '[' and ']'\n")
	} else {
		log.Println("Check - OK.")
	}
}

func (fucker *fucker)Run() {
	for fucker.Codepos < fucker.Codelen {
		switch fucker.Code[fucker.Codepos] {
		case '>':
			fucker.Pointer++
		case '<':
			fucker.Pointer--
		case '+':
			fucker.Cells[fucker.Pointer]++
		case '-':
			fucker.Cells[fucker.Pointer]--
		case ',':
			fmt.Scanf("%c",&fucker.Cells[fucker.Pointer])
		case '.':
			fmt.Printf("%c",fucker.Cells[fucker.Pointer])
		case ']':
			if fucker.Cells[fucker.Pointer] != 0 {
				count := 0
				for i := fucker.Codepos - 1; i >= 0; i -- {
					if fucker.Code[i] == ']' {
						count ++
					}
					if fucker.Code[i] == '[' {
						if count == 0 {
							fucker.Codepos = i
							break
						}
						count --
					}
				}
			}
		}
		fucker.Codepos++
	}
}