package logger

import (
	"errors"
	"fmt"
	"os"
	"strings"
    "log"
)

// reset color to default
var Reset string = "\033[0m"

// color of text appearance
var colors map[string]string = map[string]string{
    "black": "\033[30m",
    "red": "\033[31m",
    "green": "\033[32m",
    "yellow": "\033[33m",
    "blue": "\033[34m",
    "magenta": "\033[35m",
    "cyan": "\033[36m",
    "white": "\033[37m",
    "reset": "\033[0m",
}

// settings for text appearance
var config map[string]string = map[string]string{
    "none": "",
    "bold": "\033[1m",
    "underline": "\033[4m",
    "blink": "\033[5m",
    "background": "\033[7m",
    "hidden": "\033[8m",
    "dim": "\033[2m",
}

func Log(message string, color string, formatting []string) {
    var format string
    color = strings.ToLower(color)
    if color, ok := colors[color]; ok {
        for i := 0; i < len(formatting); i++ {
            if formatting[i], ok = config[formatting[i]]; ok {
                format += formatting[i]
            }else {
                err := errors.New("invalid type formatting")
                fmt.Println(err)
                message=""
            }
        }
        
        if message != "" {
            fmt.Println(string(color), string(format), message, string(Reset))
        }
    } else {
        err := errors.New("invalid color")
        fmt.Println(err)
    }
}

/* 
    Logging functions
*/
func Logs(fileName string){
    file, err := os.OpenFile(fileName + ".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    log.SetOutput(file)
}