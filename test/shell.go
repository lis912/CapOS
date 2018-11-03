package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"os"
	// "os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("capos> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if input[0] == 10 {
			continue
		}
		// fmt.Printf("%T, %s, %d", input, input, len(input))
		// for _, i := range input {
		// 	fmt.Println(i)
		// }
		// Handle the execution of the input.
		err = execInput(input)
		if err != nil {
			fmt.Println("错了")
			continue
			// fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// ErrNoPath is returned when 'cd' was called without a second argument.var
	ErrNoPath := errors.New("path required")
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")
	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")
	// Cmd(args)
	// Check for built-in commands.
	switch args[0] {
	// 'cd' to home with empty path not yet supported.
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}
		err := os.Chdir(args[1])
		if err != nil {
			return err // Stop further processing.
		}
		return nil
	case "exit":
		// return nil
		os.Exit(0) // Prepare the command to execute.
	}
	// Cmd(args)
	cmd := exec.Command(args[0], args[1:]...) // Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and save it's output.
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func Cmd(cmdline []string) {
	//实例化cli
	app := cli.NewApp()
	//Name可以设定应用的名字
	app.Name = "capos"
	// Version可以设定应用的版本号
	app.Version = "2.0.0"
	// Commands用于创建命令
	app.Commands = []cli.Command{
		{
			// 命令的名字
			Name: "print",
			// 命令的缩写，就是不输入language只输入lang也可以调用命令
			Aliases: []string{"p"},
			// 命令的用法注释，这里会在输入 程序名 -help的时候显示命令的使用方法
			Usage: "print format-type",
			// 命令的处理函数
			Action: func(c *cli.Context) error {
				// format := c.Args().First()
				format := cmdline[2]
				if format == "txt" {
					fmt.Println("format is txt")
				} else if format == "docx" {
					fmt.Println("format is docx")
				}
				return nil
			},
		},
	}
	// 接受os.Args启动程序
	// app.Run(os.Args)
	app.Run(cmdline)
}
