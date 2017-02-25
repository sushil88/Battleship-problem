package main

import (
	"os"
	"bufio"
	"log"
	"strconv"
	"errors"
	"Battleship-problem/battle"
	"Battleship-problem/ioHelpers"
)

type playerInfo struct {
	ShipPositions [][]string
	AttackMoves [][]string
	playerDamage int
}



func check(e error) {
	if e != nil {
		log.Fatal(e)
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
			player1.ShipPositions, err = ioHelpers.PrepareShipPositions(battleGroundSize, numberOfShips, scanner.Text())
		} else {
			check(errors.New("No ship positions available for palyer 1"))
		}

		if scanner.Scan() { // Get player 2 ship positions
			player2.ShipPositions, err = ioHelpers.PrepareShipPositions(battleGroundSize, numberOfShips, scanner.Text())
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
			player1.AttackMoves, err = ioHelpers.PrepareAttackPositions(totalMissiles, scanner.Text())
			check(err)
		} else {
			check(errors.New("No attack positions available for palyer 1"))
		}
		if scanner.Scan() { // Get player 2 missile moves
			player2.AttackMoves, err = ioHelpers.PrepareAttackPositions(totalMissiles, scanner.Text())
			check(err)
		} else {
			check(errors.New("No attack positions available for player 2"))
		}

		// Start battle between players
		player1.ShipPositions, player1.playerDamage = battle.DoBattle(battleGroundSize, player1.ShipPositions, player2.AttackMoves)
		player2.ShipPositions, player2.playerDamage = battle.DoBattle(battleGroundSize, player2.ShipPositions, player1.AttackMoves)

		// Write battle results to the file
		outputFileWrite, err := os.Create("output.txt")
		check(err)
		defer outputFileWrite.Close()
		_, err = outputFileWrite.WriteString("Player1\n")
		check(err)
		_, err = outputFileWrite.WriteString(ioHelpers.PrintBattlePositions(battleGroundSize, player1.ShipPositions))
		check(err)
		_, err = outputFileWrite.WriteString("Player2\n")
		check(err)
		_, err = outputFileWrite.WriteString(ioHelpers.PrintBattlePositions(battleGroundSize, player2.ShipPositions))
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
