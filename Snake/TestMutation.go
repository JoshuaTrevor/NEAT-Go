package snake

import "fmt"

type SnakeTest struct {
	Head Coords
	Body []Coords
}

func DoTest() {
	test := SnakeTest{
		Head: Coords{0, 0},
		Body: []Coords{{1, 1}},
	}
	fmt.Println("Before mutation method", test)
	testMutate(&test)
	fmt.Println("After mutation method", test)
}

func testMutate(test *SnakeTest) {
	test.Head = Coords{2, 2}
	test.Body[0] = Coords{3, 3}
	fmt.Println("In mutation method", test)
}
