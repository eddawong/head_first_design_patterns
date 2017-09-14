/*
 设计模式第一章的内容。
 实现一个简单的策略模式。
 相比较原书中java实现， golang实现要直白简单的多。
 毕竟号称静态的duck type。
 */

package chap1

import "fmt"

type Flyer interface {
	Fly()
}

type FlyWithWings struct {
}

func (f *FlyWithWings) Fly() {
	fmt.Println("I'm Flying!")
}

type FlyNoWay struct {
}

func (f *FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}

type Quacker interface {
	Quack()
}

type Quack struct {
}

func (q *Quack) Quack() {
	fmt.Println("Quack!")
}

type MuteQuack struct {
}

func (q *MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

type Squeak struct {
}

func (q *Squeak) Quack() {
	fmt.Println("Squeak!")
}

// Ducker 鸭子接口。
// 鸭子应该实现 Fly 和 Quack 方法。
type Ducker interface {
	Flyer
	Quacker
}

// 具体的鸭子类型

type Duck struct {
	FlyBehavior   Flyer
	QuackBehavior Quacker
}

func (d *Duck) Fly() {
	d.FlyBehavior.Fly()
}

func (d *Duck) Quack() {
	d.QuackBehavior.Quack()
}

func PerformDuck(d Ducker) {
	d.Fly()
	d.Quack()
}
