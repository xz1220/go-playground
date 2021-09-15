package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var workstation []int
	var sterille []int
	var sterilleRange []int
	var numSterille int

	var strWorkstation string
	var strSterille string

	in := bufio.NewReader(os.Stdin)
	strWorkstation, _ = in.ReadString('\n')
	strSterille, _ = in.ReadString('\n')

	workstations := strings.Split(strWorkstation[0:len(strWorkstation)-1], " ")
	sterilles := strings.Split(strSterille[0:len(strSterille)-1], " ")
	for _, w := range workstations {
		num, _ := strconv.Atoi(w)
		workstation = append(workstation, num)
	}
	sort.Ints(workstation)
	for _, w := range sterilles {
		num, _ := strconv.Atoi(w)
		sterille = append(sterille, num)

		// sterilleRange = append(sterilleRange, num)
		// sterilleRange = append(sterilleRange, num)
	}
	sort.Ints(sterille)
	sort.Ints(sterilleRange)
	numSterille = len(sterille)

	// 挑选候选集合
	var ansCandidates []int
	for _, w := range workstation {
		for _, s := range sterille {
			dis := math.Abs(float64(w - s))
			ansCandidates = append(ansCandidates, int(dis))
		}
	}
	sort.Ints(ansCandidates)

	// 集合筛选
	var pre int = math.MaxInt64 - 1
	var ans int
	var minIndex int = 0
	for _, candidate := range ansCandidates {
		if pre == candidate {
			continue
		}

		// 更新range
		for index := 0; index < numSterille; index++ {
			lower := sterille[index] - (candidate - pre)
			higher := sterille[index] + (candidate + pre)

			if pre == math.MaxInt64-1 {
				lower = sterille[index] - candidate
				higher = sterille[index] + candidate
			}

			for ; minIndex < len(workstation) && workstation[minIndex] >= lower && workstation[minIndex] <= higher; minIndex++ {
			}
		}

		if minIndex == len(workstation) {
			ans = candidate
			break
		}

		pre = candidate
	}

	fmt.Println(ans)
}
