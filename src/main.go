package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	CSharp = iota
	CPP
	Go
	Java
	JavaScript
	C
	TypeScript
	Swift
	Python
	Dart
	Rust
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Amount of repos
	repo_amount := get_amount_of_repos(reader)

	// Lang
	lang_repo := get_lang_for_repo(reader)

	// .bat
	make_batch(lang_repo, repo_amount)

	run_batch()
}

func get_amount_of_repos(reader *bufio.Reader) int {
	fmt.Println("How many repositories you want to create?: ")

	count, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading amount of repositories! :", err)
	}

	count = strings.TrimSpace(count)
	count_int, err := strconv.Atoi(count)

	if err != nil {
		fmt.Println("Error converting string to int in repositories count! :", err)
	}

	return count_int
}

func get_lang_for_repo(reader *bufio.Reader) int {
	fmt.Println("Which PL you want to display in repositories?: ")
	fmt.Println("0. C#")
	fmt.Println("1. С++")
	fmt.Println("2. Go")
	fmt.Println("3. Java")
	fmt.Println("4. JavaScript")
	fmt.Println("5. C")
	fmt.Println("6. TypeScript")
	fmt.Println("7. Swift")
	fmt.Println("8. Python")
	fmt.Println("9. Dart")
	fmt.Println("10. Rust")

	lang, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading language! :", err)
	}

	lang = strings.TrimSpace(lang)
	lang_int, err := strconv.Atoi(lang)

	if err != nil {
		fmt.Println("Error converting string to int in getting language! :", err)
	}

	return lang_int
}

func make_random_name(lenght int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, lenght)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func make_batch(lang_int, amount int) {
	file, err := os.Create("Generator.bat")

	if err != nil {
		fmt.Println("Error creating .bat file! :", err)
	}

	defer file.Close()

	file.WriteString("set BASEDIR=%CD% \n")
	for i := 0; i < amount; i++ {
		folder_and_repo_name := make_random_name(10)
		folder_and_repo_msg := fmt.Sprintf("md %s \n", folder_and_repo_name)
		file.WriteString(folder_and_repo_msg)
		go_to_folder_msg := fmt.Sprintf("cd %s \n", folder_and_repo_name)
		file.WriteString(go_to_folder_msg)

		file.WriteString("git init \n")

		err = make_lang_file(file, lang_int)
		if err != nil {
			fmt.Println("Error with language! :", err)
		}

		file.WriteString("git add . \n")
		file.WriteString("git commit -m \"Initial commit\" \n")

		create_clone_msg := fmt.Sprintf("gh repo create %s --private --source=. --remote=origin --push \n", folder_and_repo_name)
		file.WriteString(create_clone_msg)

		delete_folder_msg := fmt.Sprintf("RD /S /Q %s \n", folder_and_repo_name)
		file.WriteString("cd %BASEDIR% \n")
		file.WriteString(delete_folder_msg)
		fmt.Println(".bat file done!")
	}
}

func make_lang_file(file *os.File, lang_int int) error {
	switch lang_int {
	case CSharp:
		file.WriteString("echo using System; > SomeFile.txt \n")
		file.WriteString("echo public class Program >> SomeFile.txt \n")
		file.WriteString("echo { >> SomeFile.txt \n")
		file.WriteString("echo public static void Main(string[] args) >> SomeFile.txt \n")
		file.WriteString("echo { >> SomeFile.txt \n")
		file.WriteString("echo Console.WriteLine(\"Hello, World!\"); >> SomeFile.txt \n")
		file.WriteString("echo } >> SomeFile.txt \n")
		file.WriteString("echo } >> SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.cs \n")
		return nil
	case CPP:
		// Some problems with <iostream> in batch commands, so just return 0
		file.WriteString("echo int main()  > SomeFile.txt \n")
		file.WriteString("echo { >> SomeFile.txt \n")
		file.WriteString("echo return 0; >> SomeFile.txt \n")
		file.WriteString("echo } >> SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.cpp \n")
		return nil
	case Go:
		file.WriteString("echo package main > SomeFile.txt \n")
		file.WriteString("echo import \"fmt\"  >> SomeFile.txt \n")
		file.WriteString("echo func main() { >> SomeFile.txt \n")
		file.WriteString("echo fmt.Println(\"Hello, World!\") >> SomeFile.txt \n")
		file.WriteString("echo } >> SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.go \n")
		return nil
	case Java:
		file.WriteString("echo public class HelloWorld { > SomeFile.txt \n")
		file.WriteString("echo public static void main(String[] args) {  >> SomeFile.txt \n")
		file.WriteString("echo System.out.println(\"Hello, World!\"); >> SomeFile.txt \n")
		file.WriteString("echo } >> SomeFile.txt \n")
		file.WriteString("echo } >> SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.java \n")
		return nil
	case JavaScript:
		file.WriteString("echo console.log(\"Hello, World!\"); > SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.js \n")
		return nil
	case C:
		// Also with C libraries
		file.WriteString("echo int main()  > SomeFile.txt \n")
		file.WriteString("echo { >> SomeFile.txt \n")
		file.WriteString("echo return 0; >> SomeFile.txt \n")
		file.WriteString("echo } >> SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.c \n")
		return nil
	case TypeScript:
		file.WriteString("echo console.log(\"Hello, World!\"); > SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.ts \n")
		return nil
	case Swift:
		file.WriteString("echo print(\"Hello, World!\")  > SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.swift \n")
		return nil
	case Python:
		file.WriteString("echo print(\"Hello, World!\")  > SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.py \n")
		return nil
	case Dart:
		file.WriteString("echo void main() {  > SomeFile.txt \n")
		file.WriteString("echo print(\"Hello, World!\");  >> SomeFile.txt \n")
		file.WriteString("echo }  >> SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.dart \n")
		return nil
	case Rust:
		file.WriteString("echo fn main() { > SomeFile.txt \n")
		file.WriteString("echo println!(\"Hello, World!\"); >> SomeFile.txt \n")
		file.WriteString("echo } >> SomeFile.txt \n")
		file.WriteString("ren SomeFile.txt SomeFile.rs \n")
		return nil
	default:
		return fmt.Errorf("Unknown language!")
	}
}

func run_batch() {
	cmd := exec.Command("cmd.exe", "/K", "Generator.bat")

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error with cmd! :", err)
	}

	err = os.Remove("Generator.bat")
	if err != nil {
		fmt.Println("Error with deleting .bat! :", err)
	}
}
