package main

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func NewCar(speed, batteryDrain int) Car {
	defaultBattery := 100
	defaultDistance := 0
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      defaultBattery,
		distance:     defaultDistance,
	}
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

func Drive(car Car) Car {
	remainingBattery := car.battery - car.batteryDrain
	if remainingBattery < 0 {
		return car
	}

	car.battery = remainingBattery
	car.distance = car.distance + car.speed
	return car
}

func CanFinish(car Car, track Track) bool {
	calDistance := (car.battery/car.batteryDrain)*car.speed + car.distance
	if calDistance >= track.distance {
		return true
	}
	return false

}

func main() {

}
