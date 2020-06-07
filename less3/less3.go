package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type SmallCar struct {
	Brand 		string
	Year		int
	Volume		int
	EngineStart bool
	WindowsOpen bool
	TrunkFilled int
}

type BigCar struct {
	Brand		string
	Year		int
	Volume		int
	EngineStart bool
	WindowsOpen bool
	TrunkFilled int
}

var oldSmall 	SmallCar
var modernSmall SmallCar

var oldBig		BigCar
var modernBig 	BigCar

var Queue[21]int
var QueuePos int

func cars(){

	oldSmall.Brand		 = "Opel"
	oldSmall.Year		 = 1984
	oldSmall.Volume 	 = 120
	oldSmall.EngineStart = false
	oldSmall.WindowsOpen = false
	oldSmall.TrunkFilled = 20

	modernSmall.Brand 	 	= "Mercedes"
	modernSmall.Year		= 2019
	modernSmall.Volume	 	= 100
	modernSmall.EngineStart = true
	modernSmall.WindowsOpen = true
	modernSmall.TrunkFilled = 5

	oldBig.Brand		= "Tatra"
	oldBig.Year			= 2003
	oldBig.Volume 		= 6000
	oldBig.EngineStart  = false
	oldBig.WindowsOpen  = false
	oldBig.TrunkFilled  = 55

	modernBig.Brand 	  = "MAN"
	modernBig.Year		  = 2017
	modernBig.Volume 	  = 100
	modernBig.EngineStart = true
	modernBig.WindowsOpen = true
	modernBig.TrunkFilled = 86

	fmt.Println(oldSmall)

	// Вывести разницу в возрасте больших машин
	fmt.Println("Автомобиль",modernBig.Brand, "моложе", oldBig.Brand, "на", modernBig.Year - oldBig.Year, "лет")


	// Сравнить машины целиком
	if oldBig != modernBig {
		fmt.Println("У машин", oldBig.Brand, "&", modernBig.Brand, "разные параметры.")
	}

	// Если багажник заполнен менее 50% - дополнить +20%
	if oldSmall.TrunkFilled < 50 {
		fmt.Println("В машине", oldSmall.Brand,"багажник заполнен менее половины")

		oldSmall.TrunkFilled += 20 
		fmt.Println("В машине", oldSmall.Brand,"багажник дополнен до", oldSmall.TrunkFilled ,"%")
	}
}

// qPush - Add element -> Queue
func qPush (res int) (err int) {

	if len(Queue)-1 > QueuePos {
		QueuePos++
		Queue[QueuePos] = res
		return 0
	}
	return 1
	
}

// qPop - Get element <- Queue
func qPop () (res int, err int) {

	if QueuePos > 0 {
		res = Queue[QueuePos]
		QueuePos--
		return res, 0
	}
	return 0,1
	
}


func main()  {
	
	cars()

	// Очередь FIFO - Push
	for i:=1; i<22; i++{

		if err := qPush(i); err > 0 {
			log.Println("Error - переполнение очереди", err)
			break
		}
	}
	fmt.Println(Queue)

	// Очередь FIFO - Pop
	for i:=1; i<25; i++{

		res, err := qPop();
		if  err > 0 {
			log.Println("Error - окончание очереди", err)
			break
		}

		fmt.Println("Элемент ",i,"=", res)
	}


	// Телефонная книга
	phoneBook :=make(map[string]int)
	
	phoneBook["Alex"] 	= 79031456767
	phoneBook["Anna"] 	= 79651750000
	phoneBook["Anton"] 	= 79044322222
	phoneBook["Michel"] = 79031111111

	fmt.Println(phoneBook)

	serialized, err := json.Marshal(phoneBook)	
	if err!=nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(serialized))

	err = ioutil.WriteFile("phoneBook", serialized, 0644)
	if err != nil {
		log.Fatal(err)
	}

	delete(phoneBook, "Anton")
	fmt.Println("Удалили запись Anton" )
	fmt.Println(phoneBook)

	// Восстановить phoneBook из файла
	content, err := ioutil.ReadFile("phoneBook")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	fmt.Println(phoneBook)

	phone := &phoneBook
	if  err := json.Unmarshal(content, phone); err != nil {
		log.Fatal(err)
	}
	fmt.Println(phoneBook)
}

