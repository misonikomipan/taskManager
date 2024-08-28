package main

import (
	"fmt"
	"os"
)

func main() {
	// コマンドライン引数が1つもない場合の処理
	if len(os.Args) < 2 {
		fmt.Println("No command provided. Use 'help' to see available commands.")
		return
	}

	// コマンドライン引数を処理
	handleArgs()
}
