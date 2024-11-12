package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "errors"
)

// Карта для преобразования римских чисел в арабские
var romanToIntMap = map[string]int{
    "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
    "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Карта для преобразования арабских чисел в римские
var intToRomanMap = map[int]string{
    1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
    6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Введите выражение:")
    scanner.Scan()
    input := scanner.Text()
    
    result, err := calculate(input)
    if err != nil {
        fmt.Println("Ошибка:", err)
    } else {
        fmt.Println("Результат:", result)
    }
}

// Основная функция для вычисления
func calculate(input string) (string, error) {
    parts := strings.Fields(input)
    if len(parts) != 3 {
        return "", errors.New("введите два операнда и один оператор, например: '1 + 2'")
    }

    a, b := parts[0], parts[2]
    operator := parts[1]

    isRoman := isRomanNumeral(a) && isRomanNumeral(b)
    isArabic := isArabicNumeral(a) && isArabicNumeral(b)
    
    if !(isRoman || isArabic) {
        return "", errors.New("операнды должны быть либо оба римскими, либо оба арабскими")
    }

    if isRoman {
        return calculateRoman(a, b, operator)
    }
    
    return calculateArabic(a, b, operator)
}

// Проверка, является ли строка римским числом
func isRomanNumeral(s string) bool {
    _, exists := romanToIntMap[s]
    return exists
}

// Проверка, является ли строка арабским числом от 1 до 10
func isArabicNumeral(s string) bool {
    num, err := strconv.Atoi(s)
    return err == nil && num >= 1 && num <= 10
}

// Функция для вычисления римских чисел
func calculateRoman(a, b, operator string) (string, error) {
    num1 := romanToIntMap[a]
    num2 := romanToIntMap[b]
    
    result, err := performOperation(num1, num2, operator)
    if err != nil {
        return "", err
    }

    if result < 1 {
        return "", errors.New("в римской системе нет отрицательных чисел или нуля")
    }

    return intToRomanMap[result], nil
}

// Функция для вычисления арабских чисел
func calculateArabic(a, b, operator string) (string, error) {
    num1, _ := strconv.Atoi(a)
    num2, _ := strconv.Atoi(b)
    
    result, err := performOperation(num1, num2, operator)
    if err != nil {
        return "", err
    }

    return strconv.Itoa(result), nil
}

// Выполнение арифметической операции
func performOperation(a, b int, operator string) (int, error) {
    switch operator {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        if b == 0 {
            return 0, errors.New("деление на ноль")
        }
        return a / b, nil
    default:
        return 0, errors.New("недопустимая операция")
    }
}
