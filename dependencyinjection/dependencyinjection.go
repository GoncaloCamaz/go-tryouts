package dependencyinjection

import (
	"fmt"
	"time"
)

type GunService interface {
	fireGun()
}

type WarShip struct {
	gs GunService
}

type MachineGun struct {
	rounds     int
	reloadTime int
}

func (mg MachineGun) fireGun() {
	for i := 0; i < mg.rounds; i++ {
		fmt.Printf("Firing round from machine gun %d\n", i)
		time.Sleep(time.Duration(mg.reloadTime) * 1000)
	}
}

type CanonGun struct {
	rounds     int
	reloadTime int
}

func (mg CanonGun) fireGun() {
	for i := 0; i < mg.rounds; i++ {
		fmt.Printf("Firing round from canon gun %d\n", i)
		time.Sleep(time.Duration(mg.reloadTime) * 1000)
	}
}

type AntiSubmarineGun struct {
	rounds     int
	reloadTime int
}

func (mg AntiSubmarineGun) fireGun() {
	for i := 0; i < mg.rounds; i++ {
		fmt.Printf("Firing round from anti submarine gun %d\n", i)
		time.Sleep(time.Duration(mg.reloadTime) * 1000)
	}
}

func newWarShip(gs GunService) *WarShip {
	return &WarShip{
		gs: gs,
	}
}
