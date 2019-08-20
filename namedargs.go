package main

import "fmt"

// NamedArg は、名前指定で初期化しないといけないようになっている構造体
type NamedArg struct {
	_namedFieldsRequired struct{} // 非公開で struct{} にしておけばメモリサイズもゼロ

	Value1 int
	Value2 string
}

// NamedFieldsRequired は、名前指定で初期化しないといけない構造体の作り方サンプル
// REFERENCES:: http://bit.ly/2ZfkiIg
func NamedFieldsRequired() {
	// OK
	n := &NamedArg{
		Value1: 1,
		Value2: "hoge",
	}

	fmt.Println(n)

	// NG
	// n2 := &NamedArg{ 1, "hoge" }
	// OK
	n2 := &NamedArg{struct{}{}, 1, "hoge"}
	fmt.Println(n2)

	// OK
	n3 := new(NamedArg)
	n3.Value1 = 1
	n3.Value2 = "hoge"
	fmt.Println(n3)
}
