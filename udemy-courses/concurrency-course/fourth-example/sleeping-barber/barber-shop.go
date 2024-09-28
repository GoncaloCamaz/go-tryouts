package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity       int
	HairCutDuration    time.Duration
	NumberOfBarbers    int
	BarbersDoneChannel chan bool
	ClientsChannel     chan string
	IsOpen             bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			// if there are no clients, the barber falls asleep
			if len(shop.ClientsChannel) == 0 {
				color.Yellow("There is nothing to do, %s falls asleep.", barber)
				isSleeping = true
			}

			// when we have more than one barber, we might have more than one go routine trying to read the open property of modifying the struct barber shop
			client, shopOpen := <-shop.ClientsChannel

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes up and starts cutting %s's hair.", barber, client)
					isSleeping = false
				}

				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, so send the barber home and close the go routine
				shop.sendBarberHome(barber)
				// close go routine
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s finished cutting %s's hair.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	// we need to decrement the number of barbers
	shop.BarbersDoneChannel <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day")

	close(shop.ClientsChannel)
	shop.IsOpen = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		//this will block until every barbers send the value of true
		<-shop.BarbersDoneChannel
	}

	close(shop.BarbersDoneChannel)
	color.Green("The barbershop is now closed for the day, everyone has gone home.")
	color.Green("-----------------------------------------------------------------")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("*** %s enters the barbershop. ***", client)
	// if the shop is open, we can add the client to the waiting room
	if shop.IsOpen {
		select {
		case shop.ClientsChannel <- client:
			color.Yellow("%s is waiting in the waiting room.", client)
		default:
			// the buffer is full, so the client leaves
			color.Red("The waiting room is full, so %s leaves...", client)
		}
	} else {
		color.Red("%s leaves the barbershop because it is closed.", client)
	}
}
