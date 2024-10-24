package main

import "strings"

func Parser(str string) Resultado {

	splitLines := strings.Split(str, "\n")
	var cleanLines []string
	for _, line := range splitLines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			cleanLines = append(cleanLines, trimmedLine)
		}
	}

	println(cleanLines[0], strings.Trim(cleanLines[len(cleanLines)-2], ","), cleanLines[len(cleanLines)-1])
	return Resultado{
		Quote:  cleanLines[0],
		Author: strings.Trim(cleanLines[len(cleanLines)-2], ","),
		Book:   cleanLines[len(cleanLines)-1],
	}
}
