package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	goPs "github.com/mitchellh/go-ps"
)
func CdCommand(cd []string) (error) {
	switch len(cd){
	case 1:
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		err = os.Chdir(homeDir)
		if err != nil {
		
			return err
		}
	case 2:
		err := os.Chdir(cd[1])
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("too many arguments")
	}

	return nil
}
func PwdCommand(pwd []string) (string, error) {
	switch len(pwd) {
	case 1:
		path, err := os.Getwd()
		if err != nil {
			return "", err
		} else {
			return path, nil
		}
	default:
		return "", fmt.Errorf("too many arguments")
	}
}
func EchoCommand(echo []string) (string, error) {
	var res string
	for i := 1; i < len(echo); i++ {
		res += echo[i] + " "
	}
	return res, nil
}
func KillCommand(kill []string) (error) {
	switch len(kill){
	case 1:
		return fmt.Errorf("using kill <proc number>")
	case 2:
		pid, err := strconv.Atoi(kill[1])
		if err != nil {
			return err
		}
		process, err := os.FindProcess(pid)
		if err != nil {
			return err
		}
		err = process.Kill()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("too many arguments")
	}

	return nil
}
func PsCommand(ps []string) ([]string, error) {
	result := make([]string, 0, len(ps))
	
	switch len(ps){
	case 1:
		sliceProc, _ := goPs.Processes()

		for _, proc := range sliceProc {
			str := "Process name: " + proc.Executable()+ " process id: " + strconv.Itoa(proc.Pid()) + "\n"
			result = append(result, str)

		}
	default:
		return nil, fmt.Errorf("too many arguments")
	}
	
	return result, nil
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		if str[0]== "\\quit"{
			break
		}	
		switch str[0] {
		case "cd":
			err:=CdCommand(str)
			if err != nil {
				fmt.Println(err)
			}
			
		case "pwd":
			strPrint, err:=PwdCommand(str)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(strPrint)
			
		case "echo":
			strPrint, err:=EchoCommand(str)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(strPrint)
		case "kill":
			err:=KillCommand(str)
			if err != nil {
				fmt.Println(err)
			}
		case "ps":
			strPrint, err:=PsCommand(str)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(strPrint)
		
		default:
			fmt.Printf("command not found: %s\n", str[0])
		}

	}
}
