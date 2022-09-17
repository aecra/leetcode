package problem850

import "sort"

func rectangleArea(rectangles [][]int) int {
	// Divide the space into 2n subintervals
	verticalLines := make([]int, 0, 2*len(rectangles))
	for _, rectangle := range rectangles {
		verticalLines = append(verticalLines, rectangle[0], rectangle[2])
	}
	sort.Ints(verticalLines)
	ans := 0
	// Calculate the area of each interval
	for i := 1; i < len(verticalLines); i++ {
		a, b, HorizontalLength := verticalLines[i-1], verticalLines[i], verticalLines[i]-verticalLines[i-1]
		if HorizontalLength == 0 {
			continue
		}
		// Get the top and bottom edges of the covered rectangle
		lines := [][]int{}
		for _, rectangle := range rectangles {
			if rectangle[0] <= a && b <= rectangle[2] {
				lines = append(lines, []int{rectangle[1], rectangle[3]})
			}
		}
		// Calculate the vertical length of the covered rectangle
		sort.Slice(lines, func(i, j int) bool {
			if lines[i][0] != lines[j][0] {
				return lines[i][0] < lines[j][0]
			}
			return lines[i][1] < lines[j][1]
		})
		verticalLength, bottow, top := 0, -1, -1
		for _, line := range lines {
			if line[0] > top {
				verticalLength += top - bottow
				bottow, top = line[0], line[1]
			} else if line[1] > top {
				top = line[1]
			}
		}
		verticalLength += top - bottow
		// Add the area of the interval to the answer
		singalArea := HorizontalLength * verticalLength % (1e9 + 7)
		ans = (ans + singalArea) % (1e9 + 7)
	}
	return ans
}
