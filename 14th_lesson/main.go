package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//file, err := os.Create("myfile.txt")
	//if err != nil {
	//	fmt.Println("Error creating file:", err)
	//	return
	//}
	//defer file.Close()
	//
	//file.WriteString("Hello World")
	//fmt.Println("File created successfully")
	//
	//buf := make([]byte, 1024)
	//n, err := file.Read(buf)
	//if err != nil {
	//	fmt.Println("Error reading file:", err)
	//	return
	//}
	//fmt.Printf("Read %d bytes: %s\n", n, buf[:n])

	//file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	//if err != nil {
	//	fmt.Println("Error opening file:", err)
	//	return
	//
	//}
	//defer file.Close()
	//
	//fmt.Println("File opened successfully")
	//
	//_, err = file.Write([]byte("Hello, Go!"))
	//if err != nil {
	//	fmt.Println("Error writing to file:", err)
	//	return
	//}
	//
	//fmt.Println("Data written successfully")

	//file, err := os.Open("example.txt")
	//if err != nil {
	//	fmt.Println("Error opening file:", err)
	//	return
	//}
	//defer file.Close()
	//
	//// Чтение файла и вывод его содержимого по частям
	//buffer := make([]byte, 5) // Чтение по 1024 байта
	//for {
	//	n, err := file.Read(buffer)
	//	if err == io.EOF {
	//		break // Достигнут конец файла
	//	}
	//	if err != nil {
	//		fmt.Println("Error reading file:", err)
	//		return
	//	}
	//	// Вывод прочитанной части файла
	//	time.Sleep(5 * time.Second)
	//	fmt.Print(string(buffer[:n]))
	//}

	//var input string
	//fmt.Print("Enter your name: ")
	//fmt.Fscan(os.Stdin, &input)
	//fmt.Printf("Hello, %s!\n", input)

	fileName := "example.txt"
	count, err := CountCharacters(fileName)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Количество символов в файле %s: %d\n", fileName, count)
	}

	countLines, err := CountLines(fileName)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Count of string words: ", countLines)
	}

	str := "This is test string, so you can skip"
	countWords := CountWords(str)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Count of words: ", countWords)
	}

	writeToFile := WriteStringToFile("example.txt", "Hi Fira")
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Count of words: ", writeToFile)
	}

	firstStr, err := ReadFirstLine(fileName)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("First line: ", firstStr)
	}

}

//1 Описание: Напишите программу, которая считывает файл и подсчитывает количество символов в нём.
//Методы или функции:
//func countCharacters(fileName string) (int, error)

func CountCharacters(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil

}

//2 Напишите программу, которая считает количество строк в текстовом файле.
//Методы или функции:
//func countLines(fileName string) (int, error)

func CountLines(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}

//3 Напишите функцию, которая считает количество слов в строке.
//Методы или функции:
//func countWords(s string) int

func CountWords(s string) int {
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0
	}
	return count

}

//4 Напишите программу, которая записывает заданную строку в файл.
//Методы или функции:
//func writeStringToFile(fileName, content string) error

func WriteStringToFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}
	return writer.Flush()
}

//5 Напишите программу, которая читает первую строку из текстового файла.
//Методы или функции:
//func readFirstLine(fileName string) (string, error)

func ReadFirstLine(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}

//6 Напишите программу, которая копирует содержимое одного файла в другой.
//Методы или функции:
//func copyFile(src, dst string) error

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

//7 Напишите программу, которая читает строку с консоли и записывает её в файл.
//Методы или функции:
//func readAndWriteToFile(fileName string) error

func ReadAndWriteToFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строку:")
	if scanner.Scan() {
		_, err := writer.WriteString(scanner.Text())
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return writer.Flush()
}

//8 Напишите программу, которая читает файл с конца в начало и выводит его содержимое на экран.
//Методы или функции:
//func reverseReadFile(fileName string) (string, error)

func ReverseReadFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(content), "\n")
	var result strings.Builder
	for i := len(lines) - 1; i >= 0; i-- {
		result.WriteString(lines[i])
		result.WriteString("\n")
	}
	return result.String(), nil
}

//9 Напишите программу, которая объединяет содержимое двух файлов и записывает результат в новый файл.
//Методы или функции:
//func concatenateFiles(file1, file2, outputFile string) error

func ConcatenateFiles(file1, file2, outputFile string) error {
	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()

	for _, f := range []string{file1, file2} {
		in, err := os.Open(f)
		if err != nil {
			return err
		}
		_, err = io.Copy(out, in)
		in.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//10 Напишите функцию, которая проверяет, существует ли файл с заданным именем.
//Методы или функции:
//func fileExists(fileName string) bool

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

//11 Напишите программу, которая читает файл и подсчитывает количество уникальных слов.
//Методы или функции:
//func countUniqueWords(fileName string) (int, error)

func CountUniqueWords(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	uniqueWords := make(map[string]bool)
	for scanner.Scan() {
		uniqueWords[strings.ToLower(scanner.Text())] = true
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(uniqueWords), nil
}

//12 Напишите программу, которая заменяет все вхождения определенного слова в файле на другое слово.
//Методы или функции:
//func replaceWordInFile(fileName, oldWord, newWord string) error

func ReplaceWordInFile(fileName, oldWord, newWord string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	newContent := strings.ReplaceAll(string(content), oldWord, newWord)
	return os.WriteFile(fileName, []byte(newContent), 0644)
}

//13 Напишите программу, которая подсчитывает частоту каждого слова в файле.
//Методы или функции:
//func wordFrequency(fileName string) (map[string]int, error)

func WordFrequency(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	counts := make(map[string]int)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return counts, nil
}

//14 Напишите программу, которая переворачивает строки в файле и записывает их в новый файл.
//Методы или функции:
//func reverseLines(fileName, outputFile string) error

func ReverseLines(fileName, outputFile string) error {
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	lines := strings.Split(string(input), "\n")
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	return os.WriteFile(outputFile, []byte(strings.Join(lines, "\n")), 0644)
}

//15 Напишите программу, которая удаляет все пустые строки из файла.
//Методы или функции:
//func removeEmptyLines(fileName, outputFile string) error

func removeEmptyLines(fileName, outputFile string) error {
	input, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) > 0 {
			fmt.Fprintln(writer, line)
		}
	}
	return scanner.Err()
}
