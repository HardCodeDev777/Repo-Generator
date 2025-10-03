![Go](https://img.shields.io/badge/Go-%2300ADD8?logo=Go&logoColor=%23FFFFFF)
![License](https://img.shields.io/github/license/HardCodeDev777/Repo-Generator?color=%2305991d)
![Last commit](https://img.shields.io/github/last-commit/HardCodeDev777/Repo-Generator?color=%2305991d)
![Top lang](https://img.shields.io/github/languages/top/HardCodeDev777/Repo-Generator)

# 💻 Repo Generator

---

## 🚀 Overview

**Repo Generator** is a CLI tool that allows you to generate repositories right into your github profile. Also you can choose any language(see available languages below) and it will be in every repository!

---

## 📦 Installation

1. Download .exe from Realeses
2. Download GitHub CLI
3. Login in GitHub CLI

---

> [!IMPORTANT]
>  Always run RepoGenerator.exe in cmd.

## ⌨️ Usage

```bat
RepoGenerator.exe -c 5 -lang 2
```

### Explanation

`-c` stands for count and `-lang` for language. You can choose how many repositories you want to create and which language should be in them.

| Index  | Language      |
|----|------------|
| 0  | `C#` |
| 1  | `C++`         |
| 2  | `Go`     |
| 3  | `Java`       |
| 4  | `JavaScript`       |
| 5  | `C`       |
| 6  | `TypeScript`       |
| 7  | `Swift`       |
| 8  | `Python`       |
| 9  | `Dart`       |
| 10  | `Rust`       |

I will add more languages in future(maybe).

---

## ⚙️ How it works

For example, you typed `-c 2 -lang 4`, here's how it will work:

### For people who know BatchFile:

It will generate this .bat file:

```bat
set BASEDIR=%CD% 

md QoYVntuyMX 
cd QoYVntuyMX 
git init 
echo console.log("Hello, World!"); > SomeFile.txt 
ren SomeFile.txt SomeFile.js 
git add . 
git commit -m "Initial commit" 
gh repo create QoYVntuyMX --public --source=. --remote=origin --push 
cd %BASEDIR% 
RD /S /Q QoYVntuyMX 

md PCRNDGGYJc 
cd PCRNDGGYJc 
git init 
echo console.log("Hello, World!"); > SomeFile.txt 
ren SomeFile.txt SomeFile.js 
git add . 
git commit -m "Initial commit" 
gh repo create PCRNDGGYJc --public --source=. --remote=origin --push 
cd %BASEDIR% 
RD /S /Q PCRNDGGYJc
```

P.S: "PCRNDGGYJc" and "QoYVntuyMX" are randomly generated names

### For people who don't know Batchfile:

It will generate .bat file, that will:

1. Generate random name
2. Create folder with this name
3. Init git and create SomeFile.js file with `console.log("Hello, World!);` in this folder
4. Push commit with name "Initial commit"
5. Create repository in your github profile via GitHub CLI
6. Delete this folder
7. Will do it again but with another name


### Result:

You have two public repositories in your GitHub profile with JavaScript language, SomeFile.js with `console.log("Hello, World!);` in both of them; no folders or files in your computer.

---

## 📄 License

This project is licensed under the **MIT License**.  
See LICENSE for full terms.

---

> 💬 Got feedback, found a bug, or want to contribute? Open an issue or fork the repo!
