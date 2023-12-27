package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// generateMatrix Генерирует случайную матрицу 4х4 (1 и 0)
func generateMatrix() [4][4]int {
	var matrix [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matrix[i][j] = rand.Intn(2)
		}
	}
	return matrix
}

// generateKey Генерирует случайный ключ длиной 16 символов (1 и 0)
func generateKey() [16]int {
	var key [16]int
	for i := 0; i < 16; i++ {
		key[i] = rand.Intn(2)
	}
	return key
}

// countSentence Считает количество вхождений переданной последовательности из 4 элементов в ключе
func countSentence(sentence [4]int, key [16]int) int {
	var correct bool
	var keyIndex int
	var count int
	for j := keyIndex; j < 13; j++ {
		for i := 0; i < 4; i++ {
			correct = true
			if sentence[i] != key[j+i] {
				correct = false
				break
			}
		}
		if correct {
			count++
		}
		correct = false
		keyIndex++
	}
	return count
}

// watchMatrix Подсчет количества совпадений каждой последовательности ключа с каждым строкой/столбцом и главной диагональю матрицы
func watchMatrix(matrix [4][4]int, key [16]int) bool {
	var row, col [4]int
	var countRow, countCol, countDiag int
	diag := findDiagMatrix(matrix)
	countDiag = countSentence(diag, key)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			row[i] = matrix[i][j]
			col[i] = matrix[j][i]
		}
		countRow += countSentence(row, key)
		countCol += countSentence(col, key)
		if countRow+countCol+countDiag == 16 {
			return true
		}
	}
	return false
}

// findDiagMatrix находит главную диагональ матрицы
func findDiagMatrix(matrix [4][4]int) (diag [4]int) {
	for i := 0; i < 4; i++ {
		diag[i] = matrix[i][3-i]
	}
	return
}

// bruteForce Проводит перебор сгенерированных ключей, до нахождения действительного
func bruteForce(matrix [4][4]int) string {
	key := generateKey()
	ok := watchMatrix(matrix, key)
	if ok {
		return stringKey(key)
	}
	return bruteForce(matrix)
}

// printMatrix Выводит матрицу
func printMatrix(matrix [4][4]int) {
	fmt.Print("Матрица:\n")
	for i := range matrix {
		for j := 0; j < len(matrix); j++ {
			fmt.Printf(" %d ", matrix[i][j])
		}
		fmt.Print("\n")
	}
}

// stringKey Делает ключ строкой
func stringKey(key [16]int) string {
	var str string
	for i := range key {
		str += strconv.Itoa(key[i])
	}
	return str
}

func main() {
	rand.Seed(time.Now().UnixNano())

	matrix := generateMatrix()
	printMatrix(matrix)

	// Одна поптыка генерации ключа (согласно условию задачи)
	anyKey := generateKey()
	fmt.Println("Cгенерированный ключ: ", stringKey(anyKey))
	if watchMatrix(matrix, anyKey) {
		fmt.Println("...Ключ доступа действителен")
	} else {
		fmt.Println("...Ключ доступа недействителен")
	}

	// Генерация действительного ключа доступа методом перебора
	// ? Как ограничить runtime программы или отследить "stack overflow" ?
	correctKey := bruteForce(matrix)
	fmt.Println("Cгенерированный ключ: ", correctKey)
	fmt.Println("...Ключ доступа действителен")
}
