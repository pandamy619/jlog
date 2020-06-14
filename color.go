package jlog

const (
	resetColor = "\u001B[0m"

	infoColor = "\u001B[34;1m"
	warningColor = "\u001B[33m"
	errorColor = "\u001B[31m"

)

const (
	pkgPrefix = "\u001B[32mpackage" + resetColor
	funcPrefix = "\u001B[33mfunc" + resetColor
	detailPrefix = "\u001b[35mdetail" + resetColor

	infoPrefix = infoColor + "[INFO]" + resetColor
	warningPrefix = warningColor + "[WARN]" + resetColor
	errorPrefix = errorColor + "[ERROR]" + resetColor
)