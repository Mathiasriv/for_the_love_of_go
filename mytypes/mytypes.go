package mytypes

import "strings"

// Twice multiplica su receptor por 2 y devuelve
// el resultado.

type MyInt int
type MyString string

// type MyBuilder strings.Builder
type MyBuilder struct {
	Contents strings.Builder
}

type StringUpperCase struct {
	ContentsStringUpperCase strings.Builder
}
// metodo de MyInt
func (i MyInt) Twice() MyInt {
	return i * 2
}
// metodo de MyString
func (s MyString) Len() int {

	return len(s)

}

func (s MyBuilder) Hello() string {
	return "Hello, Gophers!"

}

func (st StringUpperCase) ToUpper() string {
	return strings.ToUpper(st.ContentsStringUpperCase.String())

}

func(mi *MyInt)Double() {
	*mi *= 2
}
