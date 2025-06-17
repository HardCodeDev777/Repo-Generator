package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	// 0
	CSharp = iota
	// 1
	CPP
	// 2
	Go
	// 3
	Java
	// 4
	JavaScript
	// 5
	C
	// 6
	TypeScript
	// 7
	Swift
	// 8
	Python
	// 9
	Dart
	// 10
	Rust
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	repoAmount := flag.Int("c", 0, "Amount of repositories to create")

	langRepo := flag.Int("lang", 2, "Language for repositories(see github doc for indexes - https://github.com/HardCodeDev777/Repo-Generator)")

	flag.Parse()
	// .bat
	makeBatch(*langRepo, *repoAmount)

	runBatch()
}

func makeRandomName(lenght int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, lenght)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func makeBatch(langInt, amount int) {
	file, err := os.Create("Generator.bat")

	if err != nil {
		fmt.Println("Error creating .bat file! :", err)
	}

	defer file.Close()

	file.WriteString("set BASEDIR=%CD% \n")
	for i := 0; i < amount; i++ {
		folderAndRepoName := makeRandomName(10)
		folderAndRepoMsg := fmt.Sprintf("md %s \n", folderAndRepoName)
		file.WriteString(folderAndRepoMsg)
		goToFolderMsg := fmt.Sprintf("cd %s \n", folderAndRepoName)
		file.WriteString(goToFolderMsg)

		file.WriteString("git init \n")

		err = makeLangFile(file, langInt)
		if err != nil {
			fmt.Println("Error with language! :", err)
		}

		file.WriteString("git add . \n")
		file.WriteString("git commit -m \"Initial commit\" \n")

		createCloneMsg := fmt.Sprintf("gh repo create %s --public --source=. --remote=origin --push \n", folderAndRepoName)
		file.WriteString(createCloneMsg)

		deleteFolderMsg := fmt.Sprintf("RD /S /Q %s \n", folderAndRepoName)
		file.WriteString("cd %BASEDIR% \n")
		file.WriteString(deleteFolderMsg)
	}
}

func makeLangFile(file *os.File, langInt int) error {
	switch langInt {
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

func runBatch() {
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
