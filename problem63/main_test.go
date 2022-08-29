package problem63

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	// test case 1
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	expected := 2
	actual := uniquePathsWithObstacles(obstacleGrid)
	if actual != expected {
		t.Errorf("TestUniquePathsWithObstacles failed, expected: %v, actual: %v", expected, actual)
	}

	// test case 2
	obstacleGrid = [][]int{
		{0, 1},
		{0, 0},
	}
	expected = 1
	actual = uniquePathsWithObstacles(obstacleGrid)
	if actual != expected {
		t.Errorf("TestUniquePathsWithObstacles failed, expected: %v, actual: %v", expected, actual)
	}
}
