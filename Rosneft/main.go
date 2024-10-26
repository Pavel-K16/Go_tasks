package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	log_file := CreateLogfile()
	defer log_file.Close()

	var nums []float64
	URL := "https://developer.mozilla.org/ru/docs/Web/HTTP/Status"

	fmt.Println("Для ввода массива с клавиатуры введите k, а для считывания массива из JSON файла f.")
	var symbol rune
	symbol, _, _ = bufio.NewReader(os.Stdin).ReadRune()

	switch symbol {
	case 'k':
		nums, err := Input()
		if err == nil {
			log.Println("Maccив успешно считан из терминала.")
			log.Println("Массив:", nums)
			Sum(nums)
		}
	case 'f':
		if err := Decoding(&nums); err == nil {
			log.Println("Массив чисел успешно считан из JSON файла.")
			log.Println("Массив:", nums)
			Sum(nums)
		}
	default:
		log.Println("Неверный формат выбора ввода для считывания массива.")
	}
	ResponceStatus(&URL)
	file_info, _ := os.ReadFile("log.txt")
	fmt.Println(string(file_info))
}
func Input() ([]float64, error) {
	var nums []float64
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Println("Не удалось считать числа:", err)
		log.Println("Не удалось посчитать сумму чисел:", err)
		return nums, err
	}
	text = strings.TrimSpace(text)
	numbers := strings.Split(text, " ")
	nums = make([]float64, len(numbers))
	for i, val := range numbers {
		nums[i], _ = strconv.ParseFloat(val, 64)
		if _, err = strconv.ParseFloat(val, 64); err != nil {
			log.Println("Некорректный ввод чисел:", err)
			log.Println("Не удалось посчитать сумму чисел:", err)
			break
		}
	}
	return nums, err
}

func CreateLogfile() *os.File {
	log_file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal("Ошибка при создании log.txt файла.", err)
	}
	if _, err = os.Open("log.txt"); err != nil {
		log.Fatal("Ошибка при открытии log.txt файла.", err)
	}
	log.SetOutput(log_file)
	return log_file
}
func Decoding(nums *[]float64) error {
	file, err := os.Open("test.json")
	if err != nil {
		log.Println("Ошибка при открытии json файла:", err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	data, err := rd.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Println("Ошибка при считывании из json файла", err)
	}
	if err = json.Unmarshal([]byte(data), nums); err != nil {
		log.Println("Ошибка при декодировании:", err)
	}

	return err
}
func Sum(nums []float64) {
	sum := 0.0
	for _, val := range nums {
		sum += val
	}
	log.Println("Посчитанная сумма всех чисел в массиве:", sum)
}
func ResponceStatus(URL *string) {
	resp, err := http.Get(*URL)

	if err != nil {
		log.Println("Ошибка при Get запросе:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Статус ответа: 200;", "URL:", *URL)
	} else {
		log.Println("Неожиданный статус ответа:", resp.StatusCode, "Ожидаемый статус ответа: 200.", "URL:", *URL)
	}
}
