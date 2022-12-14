# Лабораторная работа №2: Создание приложения с метео-информацией

## Задача:
> Вывести в текущем и недельном прогнозе скорость ветра и видимость.

## Описание работы программы:

1. Начало работы всех программ Go начинается с функции main пакета main.
2. В функции main происходит вызов функции getCurrentWeather, которая возвращает структуру Weather.
3. Функция getCurrentWeather создает URL для запроса к API и, используя функцию makeRequest, получает ответ, после чего передает его parseWeather, которая возвращает готовую структуру Weather, и отдает ее в main.
4. В функции main происходит вызов функции showCurrentWeather, которая выводит в консоль текущую погоду.
5. В функции main происходит вызов функции getForecast, которая возвращает структуру Forecast.
6. Функция getForecast создает URL для запроса к API и, используя функцию makeRequest, получает ответ, после чего передает его parseForecast, которая возвращает готовую структуру Forecast, и отдает ее в main.
7. В функции main происходит вызов функции showForecast, которая выводит в консоль прогноз погоды на неделю.

## Пример работы программы:

```bash
root@oustrix: ./main
Погода сейчас:
Температура: 3°C
Ощущается как: 2°C
Минимальная температура: 2°C
Максимальная температура: 4°C
Давление: 1024 hPa
Влажность: 78%
Видимость: 10000 м
Скорость ветра: 1 м/с
Направление ветра: 205°
Облачность: 86%


Погода сегодня в 21:00:
Температура: 3°C
Ощущается как: 3°C
Минимальная температура: 3°C
Максимальная температура: 4°C
Давление: 1024 hPa
Влажность: 78%
Видимость: 10000 м
Скорость ветра: 1 м/с
Направление ветра: 209°
Облачность: 85%

....

Погода 01.11.2022 в 18:00:
Температура: 0°C
Ощущается как: -2°C
Минимальная температура: 0°C
Максимальная температура: 0°C
Давление: 1021 hPa
Влажность: 78%
Видимость: 10000 м
Скорость ветра: 2 м/с
Направление ветра: 349°
Облачность: 14%
```

## Запуск и билд программы из исходного кода:
### С использованием Makefile:
```bash
make build && make run
```
### Без использования Makefile
```bash
go build -o bin/main main.go && cd bin && ./main
```
### Простой запуск программы, без сохранения бинарного файла
```bash
go run main.go
```

## Примечание:
+ при использовании утилиты make необходимо наличие утилиты make
+ при сборке программы с помощью make, бинарный файл будет находиться в папке bin
+ готовый файл main находится в папке bin
+ при создании программы была использована внешняя библиотека [zakaria-chahboun/cute](github.com/zakaria-chahboun/cute) для более легкой обработки ошибок. Для установки библиотеки необходимо выполнить командду `go get github.com/zakaria-chahboun/cute`, либо `go mod tidy`
