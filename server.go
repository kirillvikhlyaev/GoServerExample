package main

import (
	"encoding/json" // Работа с json
	"fmt"           // Вывод инфы в консоль
	"log"           // Логи
	"net/http"      // Работа с сетью

	"github.com/gorilla/mux" // основная библиотека для обработки веб-запросов
)

const port = ":8080" // Указываем порт, общепринятый для серверов - 8080

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)                                 // Определяем ссылку и фукнцию на главную страницу
	router.HandleFunc("/X-Dumbledore-Mode", getStats).Methods("GET") // Получение статистики для Дамблдора
	router.HandleFunc("/houses", getHouses).Methods("GET")           // Ссылка для получения всех факультетов
	router.HandleFunc("/houses/{id}", getHouse).Methods("GET")       // Ссылка для получения факультета по ID
	router.HandleFunc("/houses/{id}", updateHouse).Methods("PUT")    // Обновляем очки факултета по ID

	fmt.Println("Serving @ http://127.0.0.1" + port) // Для проверки запустился ли наш друг

	log.Fatal(http.ListenAndServe(port, router))
}

// Корневая страница
func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is root page")) // Выводим на главной странице текст
}

// Получение всех факультетов
func getHouses(w http.ResponseWriter, r *http.Request) {
	fetchCount := len(houseList) // Количество факультетов

	jsonList, err := json.Marshal(houseList[0:fetchCount]) // Форматируем в json с 0 объекта до последнего

	if err != nil { // Проверка на ошибку и выводим json
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonList)
	}
}

// Обновление данных факультета по ID
func updateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jsoon")
	params := mux.Vars(r) // параметры - это тело запроса

	for index, item := range houseList { // for in по массиву факультетов
		if item.Id == params["id"] { // находим нужный факультет по id
			houseList = append(houseList[:index], houseList[index+1:]...) // убираем этот факультет,
			var house House                                               // создаем новый объект и наполняем данными из
			_ = json.NewDecoder(r.Body).Decode(&house)                    // тела запроса
			house.Id = params["id"]
			houseList = append(houseList, house)
			json.NewEncoder(w).Encode(house)
			return
		}
	}
	json.NewEncoder(w).Encode(houseList)
}

// Получение факультета по ID
func getHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range houseList {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&House{})
}

// Получение статистики для Дамблдора
func getStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonBody, err := json.Marshal(info1)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(jsonBody)
	}
}

type House struct { // Класс House - факультеты Хогвартса
	Id    string `json: "id"`
	Name  string `json: "name"`
	Score string `json: "score"`
}

type Info struct { // Информация для Дамблдора
	Id         string `json: "id"`
	DeviceInfo string `json: "deviceInfo"`
}

var info1 = Info{"131341", "Android 12, SF-313"}

var houseList = []House{ // Список факультетов Хогвартса
	House{"0", "Slytherin", "15"},
	House{"2", "Gryffindor", "42"},
	House{"3", "Hufflepuff", "32"},
	House{"1", "Ravenclaw", "5"},
}

//..>cd GoServer/
//../GoServer>go run server.go
