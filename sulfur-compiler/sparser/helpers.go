package sparser

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

const debugParse = false

func debug(cursor *parseCursor, procRole string) {
	if debugParse {
		_, file, line, ok := runtime.Caller(1)
		var refName string
		if ok {
			fileName := filepath.Base(file)
			coreName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			if strings.HasPrefix(fileName, "parse_ast_") {
				corePart := strings.TrimPrefix(coreName, "parse_ast_")
				astPart := strings.ToUpper(corePart)
				refName = fmt.Sprintf("%s:%s:%d", astPart, procRole, line)
			} else {
				refName = fmt.Sprintf("%s:%s:%d", coreName, procRole, line)
			}
		} else {
			refName = fmt.Sprintf("-:%s:-", procRole)
		}
		fmt.Printf("[DEBUG %s] %s against %s `%s`\n", refName, cursor.currentNode.Kind, cursor.currentNode.src.TokenType, cursor.currentNode.src.Content)
	}
}

func debugNext(cursor *parseCursor) {
	if debugParse {
		fmt.Printf("== %s token '%s' ==\n", cursor.currentNode.Kind, cursor.src.Content)
	}
}

func parseError(cursor *parseCursor, procRole string, msg string) error {
	_, file, line, ok := runtime.Caller(1)
	var refName string
	if ok {
		fileName := filepath.Base(file)
		coreName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		if strings.HasPrefix(fileName, "parse_ast_") {
			corePart := strings.TrimPrefix(coreName, "parse_ast_")
			astPart := strings.ToUpper(corePart)
			refName = fmt.Sprintf("%s:%s:%d", astPart, procRole, line)
		} else {
			refName = fmt.Sprintf("%s:%s:%d", coreName, procRole, line)
		}
	} else {
		refName = fmt.Sprintf("-:%s:-", procRole)
	}
	return fmt.Errorf("[%s] token `%s` at line %d col %d: %s", refName, cursor.src.Content, cursor.src.SourceLine, cursor.src.SourceOffset, msg)
}
