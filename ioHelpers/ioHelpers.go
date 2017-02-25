package ioHelpers

import (
	"strings"
	"errors"
	"strconv"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func PrepareShipPositions(battleGroundSize int, numberOfShips int, input string)([][]string, error) {
	shipPositions := make([][]string,battleGroundSize)
	for i:=0; i < battleGroundSize; i++ {
		shipPositions[i] = make([]string, battleGroundSize)
	}
	var err error
	playerShips := strings.Split(strings.TrimSpace(input), ",")
	if len(playerShips) != numberOfShips {
		err = errors.New("Number of ship postions not matching with total ship count")
		return shipPositions, err
	}
	for _, shipPosition := range playerShips {
		ship := strings.Split(strings.TrimSpace(shipPosition),":")
		if len(ship) != 2 {
			err = errors.New("Invalid ship position")
			return shipPositions, err
		}
		X, err := strconv.Atoi(strings.TrimSpace(ship[0]))
		check(err)
		Y, err := strconv.Atoi(strings.TrimSpace(ship[1]))
		check(err)
		shipPositions[X][Y] = "B"
	}
	return shipPositions, err
}

func PrepareAttackPositions(totalMissiles int, input string)([][]string, error) {
	attackPositions := make([][]string,totalMissiles)
	for i:=0; i < totalMissiles; i++ {
		attackPositions[i] = make([]string, totalMissiles)
	}
	var err error
	playerAttacks := strings.Split(strings.TrimSpace(input), ":")
	if len(playerAttacks) != totalMissiles {
		err = errors.New("Number of attack positions not matching with total missiles")
		return attackPositions, err
	}
	for _, attackPosition := range playerAttacks {
		ship := strings.Split(strings.TrimSpace(attackPosition),",")
		if len(ship) != 2 {
			err = errors.New("Invalid attack position")
			return attackPositions, err
		}
		X, err := strconv.Atoi(strings.TrimSpace(ship[0]))
		check(err)
		Y, err := strconv.Atoi(strings.TrimSpace(ship[1]))
		check(err)
		attackPositions[X][Y] = "A"
	}
	return attackPositions, err
}

func PrintBattlePositions(battleGroundSize int, playerPositions [][]string)  string {
	matrixView := ""
	for i := 0; i < battleGroundSize; i++ {
		row := ""
		for j := 0; j < battleGroundSize; j++ {
			if playerPositions[i][j] == "" {
				row = row+"_"
			} else {
				row = row+playerPositions[i][j]
			}
		}
		matrixView = matrixView+row+"\n"
	}
	return matrixView
}
