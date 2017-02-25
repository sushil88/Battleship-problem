package main

import (
	"os"
	"bufio"
	"log"
	"strconv"
	"strings"
	"errors"
)

type playerInfo struct {
	ShipPositions [][]string
	AttackMoves [][]string
	playerDamage int
}

func doBattle(battleGroundSize int, shipPositions [][]string, attackPositions [][]string) ([][]string, int) {
	totalDamage := 0
	for i := 0; i < battleGroundSize; i++ {
		for j := 0; j < battleGroundSize; j++ {
			if attackPositions[i][j] == "A" && shipPositions[i][j] == "B" {
				shipPositions[i][j] = "X"
				totalDamage++
			} else if attackPositions[i][j] == "A" {
				shipPositions[i][j] = "O"
			}
		}
	}
	return shipPositions, totalDamage
}

func getShipPositions(battleGroundSize int, numberOfShips int, input string)([][]string, error) {
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

func getAttackPositions(battleGroundSize int, totalMissiles int, input string)([][]string, error) {
	attackPositions := make([][]string,battleGroundSize)
	for i:=0; i < battleGroundSize; i++ {
		attackPositions[i] = make([]string, battleGroundSize)
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

func printBattlePositions(battleGroundSize int, playerPositions [][]string)  string {
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Get input values form file
	if inputFileOpen, err := os.Open("input.txt"); err == nil {
		var battleGroundSize int
		var numberOfShips int
		var totalMissiles int
		var err error
		var player1 playerInfo
		var player2 playerInfo
		// make sure it gets closed
		defer inputFileOpen.Close()

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(inputFileOpen)
		if scanner.Scan() { // Get size of battle ground
			battleGroundSize, err = strconv.Atoi(scanner.Text())
			check(err)
		} else {
			check(errors.New("No battleship ground size available"))
		}

		if scanner.Scan() { // Get number of ships
			numberOfShips, err = strconv.Atoi(scanner.Text())
			check(err)
		} else {
			check(errors.New("No ships size available"))
		}

		if scanner.Scan() { // Get player 1 ship positions
			player1.ShipPositions, err = getShipPositions(battleGroundSize, numberOfShips, scanner.Text())
		} else {
			check(errors.New("No ship positions available for palyer 1"))
		}

		if scanner.Scan() { // Get player 2 ship positions
			player2.ShipPositions, err = getShipPositions(battleGroundSize, numberOfShips, scanner.Text())
			check(err)
		} else {
			check(errors.New("No ship positions available for palyer 2"))
		}

		if scanner.Scan() { // Get total number of missiles player have
			totalMissiles, err = strconv.Atoi(scanner.Text())
			check(err)
		} else {
			check(errors.New("No missile count available"))
		}

		if scanner.Scan() { // Get player 1 missile moves
			player1.AttackMoves, err = getAttackPositions(battleGroundSize, totalMissiles, scanner.Text())
			check(err)
		} else {
			check(errors.New("No attack positions available for palyer 1"))
		}
		if scanner.Scan() { // Get player 2 missile moves
			player2.AttackMoves, err = getAttackPositions(battleGroundSize, totalMissiles, scanner.Text())
			check(err)
		} else {
			check(errors.New("No attack positions available for player 2"))
		}

		// Start battle between players
		player1.ShipPositions, player1.playerDamage = doBattle(battleGroundSize, player1.ShipPositions, player2.AttackMoves)
		player2.ShipPositions, player2.playerDamage = doBattle(battleGroundSize, player2.ShipPositions, player1.AttackMoves)

		// Write battle results to the file
		outputFileWrite, err := os.Create("output.txt")
		check(err)
		defer outputFileWrite.Close()
		_, err = outputFileWrite.WriteString("Player1\n")
		check(err)
		_, err = outputFileWrite.WriteString(printBattlePositions(battleGroundSize, player1.ShipPositions))
		check(err)
		_, err = outputFileWrite.WriteString("Player2\n")
		check(err)
		_, err = outputFileWrite.WriteString(printBattlePositions(battleGroundSize, player2.ShipPositions))
		check(err)
		_, err = outputFileWrite.WriteString("P1:"+strconv.Itoa(player1.playerDamage)+"\n")
		check(err)
		_, err = outputFileWrite.WriteString("P2:"+strconv.Itoa(player2.playerDamage)+"\n")
		gameResult := ""
		if player1.playerDamage == player2.playerDamage {
			gameResult = "It is a draw\n"
		} else if player1.playerDamage < player2.playerDamage {
			gameResult = "Player 1 wins"
		} else {
			gameResult = "Player 2 wins"
		}
		_, err = outputFileWrite.WriteString(gameResult)
		check(err)
		outputFileWrite.Sync()
	} else {
		log.Fatal(err)
	}
}
