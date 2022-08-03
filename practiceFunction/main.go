package main

import (
	"fmt"

)
/*
	1. 메서드는 반드시 메서드명이 있어야 합니다.
	2. 매개변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없다.
	3. 인터페이스에서는 메서드 구현을 포함하지 않음.


type Stringer interface {
	String() string //메서드
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string { //Student String() 메서드 구현
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
	// 문자열 만들기
}
*/
func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")

	for i := 0; i< 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for i := 0; i<3; i++ {
		f[i]()
	}
}
func CaptureLoop2() {
	f := make([]func (), 3) 
	fmt.Println("ValueLoop2")
	for i := 0; i<3; i++ {
		v := i 
		f[i] = func() {fmt.Println(v)}
		
	}
	for i := 0; i<3; i++ {
		f[i]()
	}
}
func main() {/*
	student := Student{"철수", 12} // student 타입
	var stringer Stringer        // Stringer 타입

	stringer = student // stringer값으로 student 대입

	fmt.Printf("%s\n", stringer.String()) // stringer의 String() 메서드 호출
	*/
	CaptureLoop()
	CaptureLoop2()
}
