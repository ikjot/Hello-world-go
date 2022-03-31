package main

// ------------------------ 1 -----------------------
// Package == Project == Workspace.
// Very first line.
// Two types of packages
//	1) Executable: Generates a file that we can run.
//	2) Reusable: Helpers/Libraries/Dependency.

// Package main -> go build -> main.exe
// "Main" -> is used to generate an executable/runnable. Any other package name -> is resusable.
// "package blahblah" -> go build -> nothing!

// ------------------- 2 ----------------------------
// fmt -> standard library (format)
import "fmt"

// ------------------- 3 ----------------------------
// func -> function.
func main() {
	fmt.Println("Hi there!")
}

// Go CLI
//	go build - compile
// 	go run - compiles + execute one or two files
//	go fmt - formats all code
//	go install - compiles + installs a package
//	go get - download raw source code of someone else's package
// 	go test - Run tests

// ------------------- 4 ----------------------------
// File organization
// package
// imports
// functions
 