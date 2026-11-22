package main

// "flag"
// "fmt"
// "os"

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Аргументов менее 2х")
	// 	os.Exit(1)
	// }
	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer f.Close()
	// buffer := bufio.NewScanner(f)
	// numberString := 1
	// for buffer.Scan() {
	// 	text := buffer.Text()
	// 	fmt.Printf("%d: %s\n", numberString, text)
	// 	numberString++
	// }

	// 3 задача
	// input := flag.String("input", "file.text", "Источник данных")
	// output := flag.String("output", "file.text", "Файл для записи")
	// flag.Parse()

	// f1, err := os.Open(*input)
	// f2, err := os.OpenFile(*output, os.O_APPEND|os.O_WRONLY, 0644)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// defer f1.Close()
	// defer f2.Close()

	// buffer := bufio.NewScanner(f1)
	// currentNumber := 0
	// for buffer.Scan() {
	// 	text := buffer.Text()
	// 	currentNumber++
	// 	fmt.Fprintf(f2, "%d: %s\n", currentNumber, text)
	// }

	// if err := buffer.Err(); err != nil {
    //     fmt.Printf("ошибка чтения: %v\n", err)
    //     os.Exit(1)
    // }

	// fmt.Printf("Готово: %d строки записано в %s\n", currentNumber, *output)

	// Задача 4
	// data := flag.String("src", "", "Источник")
	// copy := flag.String("dst", "", "Копия")
	// flag.Parse()

	// dataFile, errData := os.ReadFile(*data)
	// if errData != nil {
	// 	fmt.Println(errData)
	// }

	// errCopy := os.WriteFile(*copy, dataFile, 0644)
	// if errCopy != nil {
	// 	fmt.Println(errCopy)
	// 	os.Exit(1)
	// }

	// fmt.Printf("Скопировано: %d байт\n", len(dataFile))
}