package main

import "fmt"

func main() {

	rawLiteral := "아리랑 \n"

	/*
		아래와 같이 +를 사용하여 두 라인에 걸쳐 사용할 수도 있다.
		interLiteral := "아리랑 아리랑 \n" + "아라리요"
	*/

	interLiteral := "아리랑아리랑\n아라리요"

	fmt.Println(rawLiteral)

	fmt.Println()

	fmt.Println(interLiteral)

}
