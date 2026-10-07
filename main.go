package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

/* func main() {
	const filename = "users.csv"

	// Открываем csv файл
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Не удалось открыть файл %s: %v", filename, err)
	}
	defer file.Close()

	// Создаем новый reader
	reader := csv.NewReader(file)

	// Пропускаем заголовок
	_, err = reader.Read()

	if err != nil {
		log.Fatalf("Ошибка чтения заголовка: %v", err)
	}

	// Читаем все оставшиеся записи
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Ошибка чтения csv файла: %v", err)
	}

	// Выводим записи на экран
	for _, record := range records {
		fmt.Printf("ID: %s, Name: %s, Email: %s\n", record[0], record[1], record[2])
	}
}*/

/*type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	// Исходные данные
	products := []Product{
		{ID: 1, Name: "Монитор", Price: 21990},
		{ID: 2, Name: "Видеокарта", Price: 69900},
		{ID: 3, Name: "Ноутбук", Price: 79500},
	}

	fmt.Println("Запись данных в csv файл...")

	// Создаем csv файл
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalf("Не удалось создать csv файл: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush() // Очищаем буфер в конце работы

	// Записываем заголовок
	header := []string{"ID", "Name", "Price"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("Ошибка записи заголовка: %v", err)
	}

	// Проходим по срезу продуктов и записываем каждую структуру в файл
	for _, product := range products {
		id := strconv.Itoa(product.ID)
		price := strconv.FormatFloat(product.Price, 'f', 2, 64)

		record := []string{id, product.Name, price}
		err := writer.Write(record)
		if err != nil {
			log.Fatalf("Ошибка записи строки: %v", err)
		}
	}

	fmt.Println("Файл products.csv успешно создан.")
}
*/

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	// Открываем файл с пользователями
	file, err := os.Open("users.csv")
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Читаем все данные сразу
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Ошибка чтения данных из csv файла: %v", err)
	}

	// Создаем срез структур для сохранения загруженных данных
	loadedUsers := []User{}

	// Пропускаем заголовок, срезая первую строку
	if len(records) > 0 {
		records = records[1:]
	}

	// Проходимся в цикле по записям и парсим типы вручную
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalf("Ошибка чтения ID %s: %v", record[0], err)
		}

		user := User{
			ID:    id,
			Name:  record[1],
			Email: record[2],
		}

		loadedUsers = append(loadedUsers, user)
	}

	// Выводим структурированные данные на экран
	fmt.Println("Успешно загружены структуры User:")
	for _, user := range loadedUsers {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}
