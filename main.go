package main

import "fmt"

type Car struct {
	name         string
	body         body
	yearOFissue  int
	engine       engine
	transmissuin string
	suspension   string
	price        price
}

type body struct {
	typeBody string
	color    string
}

type engine struct {
	typeEgine string
	volume    float32
}

type price struct {
	sum     int
	currens string
}

func main() {
	var car1 Car
	car1.name = "Lada Priora"
	car1.body = body{typeBody: "Седан", color: "Цвета баклажана"}
	car1.yearOFissue = 2016
	car1.engine = engine{typeEgine: "Бензиновый двигатель объемом", volume: 1.6}
	car1.transmissuin = "Анальная коробка передач"
	car1.suspension = "Подвеска говно"
	car1.price = price{sum: 25000, currens: "Тенге"}

	var car2 Car
	car2.name = "BMW M5"
	car2.body = body{typeBody: "Седан", color: "Синий"}
	car2.yearOFissue = 2019
	car2.engine = engine{typeEgine: "Бензиновый двигатель объемом", volume: 5.5}
	car2.transmissuin = "Автоматическая коробка передач"
	car2.suspension = "Задний привод"
	car2.price = price{sum: 1500000, currens: "Тенге"}

	var car3 Car
	car3.name = "BMW X5 M"
	car3.body = body{typeBody: "Внедорожник", color: "Черный"}
	car3.yearOFissue = 2018
	car3.engine = engine{typeEgine: "Бензиновый двигатель объемом", volume: 5.5}
	car3.transmissuin = "Автоматическая коробка передач"
	car3.suspension = "Полный привод"
	car3.price = price{sum: 4, currens: "Человеческой почки"}

	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car3)
}
