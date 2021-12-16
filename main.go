package main

import (
	"constraints"
	"fmt"
)

type Pattern string

func Do[V any](v V) (V, error) {
	return v, nil
}

func comp[C comparable](l, r C) bool {
	return l == r
}

func comp2[C constraints.Ordered](l, r C) bool {
	return l > r
}

type Hoe interface {
	int | string
}

func doHoe[H Hoe](h H) H {
	return h
}

type S1 interface {
	string
}

func doS1[H S1](h H) H {
	return h
}

type Foo interface {
	Hoe | S1
}

func doFoo[H Foo](h H) H {
	return h
}

type SAny interface {
	~string
}

func doSAny[H SAny](h H) H {
	return h
}

type IntOrString interface {
	string | int
}

type Into[T any] interface {
	into() T
}

type Int struct {
	value int
}

func (i Int) into() int { return i.value }

type Uint struct {
	value uint
}

func (i Uint) into() uint { return i.value }

func add[I constraints.Integer, II Into[I]](l, r II) I {
	return l.into() + r.into()
}

func from[To any, From Into[To]](v From) To {
	return v.into()
}

func gen[I constraints.Integer]() I {
	return 1
}

func main() {
	// p := Pattern("a")
	fmt.Println(Do("a"))
	{
		s := "a"
		h := doHoe(s)
		h = doS1(h)
		h = doFoo(h)
		h = doSAny(h)
	}
	{
		// ~string と 定義することで、base type が string なものを指定できる
		p := Pattern("a")
		// h := doHoe(p)
		// h = doS1(h)
		// h = doFoo(h)
		// h = doSAny(h)
		h := doSAny(p)
		fmt.Println(h)
	}
	{
		// 戻り値では推論されない
		x := gen[int]()
		fmt.Println(x)
	}
	{
		i := Int{10}
		var sum int = add[int](i, i)
		var x int = from[int](i)
		fmt.Println(sum, x)
	}
}
