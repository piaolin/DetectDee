package utils

func AddTarget(targetList [][]string, targets chan []string) {
	for i := 0; i < len(targetList); i++ {
		targets <- targetList[i]
	}
}
