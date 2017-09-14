package chap1

import "testing"

func TestPerformDuck(t *testing.T) {

	mallardDuck := &Duck{
		FlyBehavior:   &FlyWithWings{},
		QuackBehavior: &Quack{},
	}
	PerformDuck(mallardDuck)

	rubberDuck := &Duck{
		FlyBehavior:   &FlyNoWay{},
		QuackBehavior: &MuteQuack{},
	}

	PerformDuck(rubberDuck)
}
