package main

import (
	"MTUCI-VvIT-labs/lab-2/entities"
	"encoding/json"
	"fmt"
	"github.com/zakaria-chahboun/cute"
	"io"
	"net/http"
	"strings"
	"time"
)

const ApiKey = "07b2af16f7fbfe8e13ef2910ed28ee85"

func main() {
	// получение прогноза погоды на данный момент и вывод его на экран
	weather := getCurrentWeather()
	showCurrentWeather(weather)

	// получение прогноза погоды на 5 дней и вывод его на экран
	forecast := getForecast()
	showForecast(forecast)
}

func getCurrentWeather() entities.Weather {
	url := "http://api.openweathermap.org/data/2.5/weather?q=Moscow,RU&units=metric&lang=ru&APPID=" + ApiKey // создание строки запроса
	resp := makeRequest(url)                                                                                 // получаем запрос
	defer resp.Body.Close()                                                                                  // закрываем запрос после завершения работы функции

	return parseWeather(resp) // возращам готовую структуру в main
}

func makeRequest(url string) *http.Response {
	resp, err := http.Get(url)                    // отпрввляем http запрос к url
	cute.Check("Error while making request", err) // провеяем на наличие ошибка, если err === nil -> программа продолжает работу, иначе -> программа завершается с ошибкой

	if resp.StatusCode != http.StatusOK { // проверяем статус ответа, если он не равен 200 -> программа завершается с ошибкой
		cute.Check("Status code is not OK", fmt.Errorf("status code is %d", resp.StatusCode))
	}
	return resp // возвращаем http ответ
}

func parseWeather(resp *http.Response) (weather entities.Weather) {
	bodyBytes, err := io.ReadAll(resp.Body)              // считываем тело ответа
	cute.Check("Error while reading response body", err) // проверяем на наличие ошибки

	err = json.Unmarshal(bodyBytes, &weather)         // декодируем json в структуру Weather, которую создана в entities\weather.go, переменная weather определена в определении функции
	cute.Check("Error while unmarshalling json", err) // проверяем на наличие ошибки
	return weather                                    // возвращаем структуру Weather
}

// функция нужна для вывода строки "Погода сейчас:", т.к. функция showWeather() используется и для вывода недельного прогноза погоды
func showCurrentWeather(weather entities.Weather) {
	fmt.Println("Погода сейчас:")
	showWeather(weather)
}

// функция выводит всю необходимую информацию из структуры Weather
func showWeather(weather entities.Weather) {
	// использование fmt.Sprintf внутри fmt.Println - говнокод, но изначально я хотел сделать немного по-другому,
	// но потом передумал и теперь мне лень переписывать, главное, что работает, может быть, потом переделаю
	// TODO: заменить fmt.Println(fmt.Sprintf()) на fmt.Printf()
	fmt.Println(fmt.Sprintf("Температура: %.0f°C", weather.Main.Temp))
	fmt.Println(fmt.Sprintf("Ощущается как: %.0f°C", weather.Main.FeelsLike))
	fmt.Println(fmt.Sprintf("Минимальная температура: %.0f°C", weather.Main.TempMin))
	fmt.Println(fmt.Sprintf("Максимальная температура: %.0f°C", weather.Main.TempMax))
	fmt.Println(fmt.Sprintf("Давление: %d hPa", weather.Main.Pressure))
	fmt.Println(strings.ReplaceAll(fmt.Sprintf("Влажность: %d%", weather.Main.Humidity), "!(NOVERB)", "")) // заменяем "!(NOVERB)" на пустую строку, т.к. в ответе от сервера зачем-то присутствует эта строка
	fmt.Println(fmt.Sprintf("Видимость: %d м", weather.Visibility))
	fmt.Println(fmt.Sprintf("Скорость ветра: %.0f м/с", weather.Wind.Speed))
	fmt.Println(fmt.Sprintf("Направление ветра: %d°", weather.Wind.Deg))
	fmt.Println(strings.ReplaceAll(fmt.Sprintf("Облачность: %d%", weather.Clouds.All), "!(NOVERB)", "")) // заменяем "!(NOVERB)" на пустую строку, т.к. в ответе от сервера зачем-то присутствует эта строка
}

func getForecast() entities.Forecast {
	url := "http://api.openweathermap.org/data/2.5/forecast?q=Moscow,RU&units=metric&lang=ru&APPID=" + ApiKey // формируем url запроса
	resp := makeRequest(url)                                                                                  // делаем запрос
	defer resp.Body.Close()                                                                                   // закрываем соединение

	return parseForecast(resp) // возвращаем структуру Forecast
}

func parseForecast(resp *http.Response) (forecast entities.Forecast) {
	bodyBytes, err := io.ReadAll(resp.Body)              // считываем тело ответа
	cute.Check("Error while reading response body", err) // проверяем на наличие ошибки

	err = json.Unmarshal(bodyBytes, &forecast)        // декодируем json в структуру Forecast, forecast определена в определении функции
	cute.Check("Error while unmarshalling json", err) // проверяем на наличие ошибки
	return forecast                                   // возвращаем структуру Forecast
}

func showForecast(forecast entities.Forecast) {
	/*
		Формат JSONa:
		list {
			0: {...},
			1: {...},
			2: {...},
			...
		}

		for перебирает все элементы в списке list
		_ - индекс элемента в списке, т.к. мы его не используем, обознаем его как _
		weather - элемент списка list
	*/
	for _, weather := range forecast.List {
		makeIndent() // делаем отступ от предыдущего прогноза в 2 строки

		date, err := time.Parse("2006-01-02 15:04:05", weather.DtTxt) // парсим дату из строки в формате "2006-01-02 15:04:05"
		cute.Check("Error while parsing date", err)                   // проверяем на наличие ошибки
		switch {                                                      // в зависимости от даты, выводим разные сообщения
		case date.Day() == time.Now().Day(): // если день совпадает с текущим
			fmt.Printf("Погода сегодня в %s: \n", date.Format("15:04"))

		case date.Day() == time.Now().AddDate(0, 0, 1).Day(): // если день совпадает с завтрашним
			fmt.Printf("Погода завтра в %s: \n", date.Format("15:04"))

		case date.Day() == time.Now().AddDate(0, 0, 2).Day(): // если день совпадает с послезавтрашним
			fmt.Printf("Погода послезавтра в %s: \n", date.Format("15:04"))

		default: // остальные случаи
			fmt.Printf("Погода %s в %s: \n", date.Format("02.01.2006"), date.Format("15:04"))
		}

		showWeather(weather) // выводим погоду
	}
}

// функция делает отступ в 2 строки
func makeIndent() {
	fmt.Println()
	fmt.Println()
}
