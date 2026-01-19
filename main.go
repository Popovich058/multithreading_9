package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements генерирует случайные элементы
func generateRandomElements(size int) []int {
	// Обрабатываем крайние случаи
	if size <= 0 {
		return []int{}
	}
	
	// Инициализируем генератор случайных чисел текущим временем
	rand.Seed(time.Now().UnixNano())
	
	// Создаём слайс заданного размера
	result := make([]int, size)
	
	// Заполняем слайс случайными числами
	for i := 0; i < size; i++ {
		result[i] = rand.Int()
	}

	return result
}

// maximum возвращает максимальное количество элементов
func maximum(data []int) int {
	// Обрабатываем крайние случаи
	if data == nil || len(data) == 0 {
		return 0
	}

	// Инициализируем максимум первым элементом
	currentMax := data[0]

	// Проходим по оставшимся элементам
	for i := 1; i < len(data); i++ {
		if data[i] > currentMax {
			currentMax = data[i]
		}
	}

	return currentMax
}


// maxChunks возвращает максимальное количество элементов
func maxChunks(data []int) (int, error) {
	// Обрабатываем крайние случаи
	if data == nil {
		return 0, fmt.Errorf("the slice is equal to nil")
	}
	if len(data) == 0 {
		return 0, fmt.Errorf("the slice is empty")
	}

	// Если слайс меньше или равен количеству частей, ищем максимум обычным способом
	if len(data) <= CHUNKS {
		return maximum(data), nil
	}

	// Создаём слайс для хранения максимумов
	chunkMaxs := make([]int, CHUNKS)

	// Инициализируем WaitGroup
	var wg sync.WaitGroup

	// Размер одной части
	chunkSize := len(data) / CHUNKS

	// Запускаем горутины
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		// Вычисляем границы и создаём чанк
		start := i * chunkSize
		end := start + chunkSize

		// Для последней части захватываем оставшиеся элементы
		if i == CHUNKS-1 {
			end = len(data)
		}

		// Создаём чанк
		chunk := data[start:end]

		// Запускаем горутину, передавая готовый чанк
		go func(chunkIndex int, chnk []int) {
			defer wg.Done()

			// Находим максимум 
			chunkMaxs[chunkIndex] = maximum(chnk)
		}(i, chunk)
	}

	// Ждём завершения всех горутин
	wg.Wait()

	// Находим максимум среди максимумов
	finalMax := maximum(chunkMaxs)
	return finalMax, nil
}



func main() {
	fmt.Printf("Генерируем %d целых чисел...\n", SIZE)
	data := generateRandomElements(SIZE)

	// Поиск максимума в один поток
	fmt.Println("Ищем максимальное значение в один поток...")
	start := time.Now()
	max1 := maximum(data)  
	elapsed1 := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение (1 поток): %d\nВремя поиска: %d мкс\n\n", max1, elapsed1)

	// Поиск максимума в 8 потоков
	fmt.Printf("Ищем максимальное значение в %d потоков...\n", CHUNKS)
	start = time.Now()
	max2, err := maxChunks(data)  
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	elapsed2 := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение (%d потоков): %d\nВремя поиска: %d мкс\n", CHUNKS, max2, elapsed2)
}

