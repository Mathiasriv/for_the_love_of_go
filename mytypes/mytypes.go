package mytypes

// Twice multiplica su receptor por 2 y devuelve
// el resultado.

type MyInt int
type MyString string

func (i MyInt) Twice() MyInt{
    return i * 2
}


func (s MyString)Len() int{

  return len(s)

}
