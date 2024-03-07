package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium)-1; i++ {
		for j := 0; j < len(podium); j++ {
			podium[i], podium[j] = podium[j], podium[i]
		}
	}
	return podium
}
