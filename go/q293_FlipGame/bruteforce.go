package q293_FlipGame

func generatePossibleNextMoves(currentState string) []string {
	var result []string

	for i := range currentState[:len(currentState)-1] {
		if currentState[i:i+2] == "++" {
			next := currentState[:i] + "--" + currentState[i+2:]
			result = append(result, next)
		}
	}

	return result
}	