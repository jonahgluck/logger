package log

import (
    "fmt"
    "strings"
    "errors"
)

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

// func Log(message string, color string) {
//     settings := style{message, color, ""}
//     fmt.Println(string(colors[settings.color]), message, string(Reset))
// }

func Log(message string, color string, config string) {
    color = strings.ToLower(color)
    if color, ok := colors[color]; ok { 
        fmt.Println(string(color), message, string(Reset))
    } else {
        err := errors.New("invalid color")
        fmt.Println(err)
    }
}

