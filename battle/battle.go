package battle

func DoBattle(battleGroundSize int, shipPositions [][]string, attackPositions [][]string) ([][]string, int) {
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


