package jlog

const (
	brightRed = "\u001b[31;1m"
	brightGreen = "\u001b[32m"
	magenta = "\u001B[35m"

	resetColor = "\u001B[0m"

	infoColor = "\u001B[34;1m"
	warningColor = "\u001B[33m"
	errorColor = "\u001B[31m"
)

const (
	pkgPrefix = brightRed + "package" + resetColor
	funcPrefix = brightRed + "func" + resetColor
	detailPrefix = magenta + "detail" + resetColor

	infoPrefix = infoColor + "[INFO]" + resetColor
	warningPrefix = warningColor + "[WARN]" + resetColor
	errorPrefix = errorColor + "[ERROR]" + resetColor
)