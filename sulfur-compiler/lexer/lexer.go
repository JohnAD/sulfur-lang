package lexer

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"sulfur-compiler/context"
)

type Lexer struct {
	state         TokenType
	currentToken  Token
	currentLine   int
	currentOffset int
	runeCategory  RuneCategory
	previousRune  rune
	tokenList     []Token
	fileId        int
}

func NewToken(plex *Lexer, init TokenType) Token {
	t := Token{
		TokenType:    init,
		SourceFile:   plex.fileId,
		SourceLine:   plex.currentLine,
		SourceOffset: plex.currentOffset,
		Content:      "",
		PostBound:    false,
	}
	return t
}

func NewTokenWithRune(plex *Lexer, init TokenType, ch rune) Token {
	t := Token{
		TokenType:    init,
		SourceFile:   plex.fileId,
		SourceLine:   plex.currentLine,
		SourceOffset: plex.currentOffset,
		Content:      string(ch),
		PostBound:    false,
	}
	return t
}

func lexStandingSymbolBegin(plex *Lexer, ch rune) {
	plex.state = TT_STANDING_SYMBOL
	plex.currentToken = NewTokenWithRune(plex, TT_STANDING_SYMBOL, ch)
}
func lexStandingSymbol(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
		if plex.currentToken.Content == "#" {
			lexCommentBegin(plex, ch)
			return
		}
		lexStandingSymbolEnd(plex)
		lexWhitespaceBegin(plex, 0)
	case RC_LETTER:
		lexSyntaxErrorSwitch(plex, ch, "TBD LETTER")
	case RC_NUMBER:
		lexSyntaxErrorSwitch(plex, ch, "TBD NUMBER")
	case RC_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_OPEN_PUNCTUATION:
		lexSyntaxErrorSwitch(plex, ch, "TBD OPEN P")
	case RC_CLOSE_PUNCTUATION:
		lexSyntaxErrorSwitch(plex, ch, "TBD CLOSE P")
	case RC_LINE_END:
		lexStandingSymbolEnd(plex)
		lexIndentLineBegin(plex, ch)
	case RC_QUOTE:
		lexSyntaxErrorSwitch(plex, ch, "TBD QUOTE")
	case RC_FORBIDDEN:
		lexSyntaxErrorSwitch(plex, ch, "forbidden rune")
	}
}
func lexStandingSymbolEnd(plex *Lexer) {
	plex.tokenList = append(plex.tokenList, plex.currentToken)
}

func lexOpenSymbolBegin(plex *Lexer, ch rune) {
	plex.state = TT_OPEN_SYMBOL
	plex.currentToken = NewTokenWithRune(plex, TT_OPEN_SYMBOL, ch)
}
func lexOpenSymbol(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
		lexOpenSymbolEnd(plex)
		lexWhitespaceBegin(plex, 0)
	case RC_LETTER:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_NUMBER:
		lexSyntaxErrorSwitch(plex, ch, "TBD NUMBER")
	case RC_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_OPEN_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_CLOSE_PUNCTUATION:
		lexSyntaxErrorSwitch(plex, ch, "TBD CLOSE P")
	case RC_LINE_END:
		lexOpenSymbolEnd(plex)
		lexIndentLineBegin(plex, ch)
	case RC_QUOTE:
		lexSyntaxErrorSwitch(plex, ch, "TBD QUOTE")
	case RC_FORBIDDEN:
		lexSyntaxErrorSwitch(plex, ch, "forbidden rune")
	}
}
func lexOpenSymbolEnd(plex *Lexer) {
	plex.tokenList = append(plex.tokenList, plex.currentToken)
}

func lexCloseSymbolBegin(plex *Lexer, ch rune) {
	plex.state = TT_CLOSE_SYMBOL
	plex.currentToken = NewTokenWithRune(plex, TT_CLOSE_SYMBOL, ch)
}
func lexCloseSymbol(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
		lexCloseSymbolEnd(plex)
		lexWhitespaceBegin(plex, 0)
	case RC_LETTER:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_NUMBER:
		lexSyntaxErrorSwitch(plex, ch, "TBD NUMBER")
	case RC_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_OPEN_PUNCTUATION:
		lexSyntaxErrorSwitch(plex, ch, "TBD OPEN P")
	case RC_CLOSE_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_LINE_END:
		lexCloseSymbolEnd(plex)
		lexIndentLineBegin(plex, ch)
	case RC_QUOTE:
		lexSyntaxErrorSwitch(plex, ch, "TBD QUOTE")
	case RC_FORBIDDEN:
		lexSyntaxErrorSwitch(plex, ch, "forbidden rune")
	}
}
func lexCloseSymbolEnd(plex *Lexer) {
	plex.tokenList = append(plex.tokenList, plex.currentToken)
}

func lexOpenBindSymbolBegin(plex *Lexer, ch rune) {
	plex.state = TT_OPEN_BIND_SYMBOL
	plex.currentToken = NewTokenWithRune(plex, TT_OPEN_BIND_SYMBOL, ch)
}
func lexOpenBindSymbol(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
		lexOpenBindSymbolEnd(plex)
		lexWhitespaceBegin(plex, 0)
	case RC_LETTER:
		lexSyntaxErrorSwitch(plex, ch, "TBD LETTER")
	case RC_NUMBER:
		lexSyntaxErrorSwitch(plex, ch, "TBD NUMBER")
	case RC_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_OPEN_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_CLOSE_PUNCTUATION:
		lexSyntaxErrorSwitch(plex, ch, "TBD CLOSE P")
	case RC_LINE_END:
		lexOpenBindSymbolEnd(plex)
		lexIndentLineBegin(plex, ch)
	case RC_QUOTE:
		lexSyntaxErrorSwitch(plex, ch, "TBD QUOTE")
	case RC_FORBIDDEN:
		lexSyntaxErrorSwitch(plex, ch, "forbidden rune")
	}
}
func lexOpenBindSymbolEnd(plex *Lexer) {
	plex.tokenList = append(plex.tokenList, plex.currentToken)
}

func lexBindingSymbolBegin(plex *Lexer, ch rune) {
	plex.state = TT_BINDING_SYMBOL
	plex.currentToken = NewTokenWithRune(plex, TT_BINDING_SYMBOL, ch)
}
func lexBindingSymbol(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
		lexSyntaxErrorSwitch(plex, ch, "symbol unbound to any identifier or literal after")
	case RC_LETTER:
		lexBindingSymbolEnd(plex)
		lexIdentBegin(plex, ch)
	case RC_NUMBER:
		lexBindingSymbolEnd(plex)
		lexNumStrBegin(plex, ch)
	case RC_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_OPEN_PUNCTUATION:
		lexSyntaxErrorSwitch(plex, ch, "a plain binding symbol cannot contain opening punctuation")
	case RC_CLOSE_PUNCTUATION:
		lexSyntaxErrorSwitch(plex, ch, "a plain binding symbol cannot contain closing punctuation")
	case RC_LINE_END:
		lexSyntaxErrorSwitch(plex, ch, "symbol unbound to any identifier or literal after")
	case RC_QUOTE:
		lexSyntaxErrorSwitch(plex, ch, "TBD QUOTE")
	case RC_FORBIDDEN:
		lexSyntaxErrorSwitch(plex, ch, "forbidden rune")
	}
}
func lexBindingSymbolEnd(plex *Lexer) {
	plex.tokenList = append(plex.tokenList, plex.currentToken)
}

func lexIdentBegin(plex *Lexer, ch rune) {
	plex.state = TT_IDENT
	plex.currentToken = NewTokenWithRune(plex, TT_IDENT, ch)
}
func lexIdent(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
		lexIdentEnd(plex, false)
		lexWhitespaceBegin(plex, 0)
	case RC_LETTER:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_NUMBER:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_PUNCTUATION:
		lexIdentEnd(plex, true)
		lexBindingSymbolBegin(plex, ch)
	case RC_OPEN_PUNCTUATION:
		lexIdentEnd(plex, true)
		lexOpenBindSymbolBegin(plex, ch)
	case RC_CLOSE_PUNCTUATION:
		lexSyntaxErrorSwitch(plex, ch, "TBD CLOSE P")
	case RC_LINE_END:
		lexIdentEnd(plex, false)
		lexIndentLineBegin(plex, ch)
	case RC_QUOTE:
		lexSyntaxErrorSwitch(plex, ch, "TBD QUOTE")
	case RC_FORBIDDEN:
		lexSyntaxErrorSwitch(plex, ch, "forbidden rune")
	}
}
func lexIdentEnd(plex *Lexer, postBound bool) {
	plex.currentToken.PostBound = postBound
	plex.tokenList = append(plex.tokenList, plex.currentToken)
}

func lexStrLitBegin(plex *Lexer, ch rune) {
	plex.state = TT_STR_LIT
	plex.currentToken = NewToken(plex, TT_STR_LIT)
}
func lexStrLit(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_LETTER:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_NUMBER:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_OPEN_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_CLOSE_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_LINE_END:
		lexSyntaxErrorSwitch(plex, ch, "small string did not finish on same line")
	case RC_QUOTE:
		lexStrLitEnd(plex)
		lexWhitespaceBegin(plex, ch) // TODO: this is wrong, but I leave it as a hack for now
	case RC_FORBIDDEN:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	}
}
func lexStrLitEnd(plex *Lexer) {
	plex.tokenList = append(plex.tokenList, plex.currentToken)
}

func lexNumStrBegin(plex *Lexer, ch rune) {
	plex.state = TT_NUMSTR_LIT
	plex.currentToken = NewTokenWithRune(plex, TT_NUMSTR_LIT, ch)
}
func lexNumStr(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
		lexNumStrEnd(plex)
		lexWhitespaceBegin(plex, 0)
	case RC_LETTER:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_NUMBER:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_PUNCTUATION:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_OPEN_PUNCTUATION:
		lexSyntaxErrorSwitch(plex, ch, "TBD OPEN P")
	case RC_CLOSE_PUNCTUATION:
		lexSyntaxErrorSwitch(plex, ch, "TBD CLOSE P")
	case RC_LINE_END:
		lexNumStrEnd(plex)
		lexIndentLineBegin(plex, ch)
	case RC_QUOTE:
		lexSyntaxErrorSwitch(plex, ch, "TBD QUOTE")
	case RC_FORBIDDEN:
		lexSyntaxErrorSwitch(plex, ch, "forbidden rune")
	}
}
func lexNumStrEnd(plex *Lexer) {
	plex.tokenList = append(plex.tokenList, plex.currentToken)
}

func lexSyntaxErrorSwitch(plex *Lexer, ch rune, msg string) {
	errToken := NewToken(plex, TT_SYNTAX_ERROR)
	errToken.Content = fmt.Sprintf(
		"while parsing rune '%s'(%d) for '%v' on line %d col %d found error: %s. characters lexed so far: %s",
		string(ch),
		int(ch),
		plex.currentToken.TokenType,
		plex.currentLine,
		plex.currentOffset,
		msg,
		plex.currentToken.Content,
	)
	plex.state = TT_SYNTAX_ERROR
	plex.currentToken = errToken
	plex.tokenList = append(plex.tokenList, plex.currentToken)
}

// no need for lexSyntaxErrorBegin
// no need for lexSyntaxError
// no need for lexSyntaxErrorEnd

func lexCommentBegin(plex *Lexer, ch rune) {
	plex.state = TT_COMMENT
	plex.currentToken = NewToken(plex, TT_COMMENT)
}
func lexComment(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
	case RC_LETTER:
	case RC_NUMBER:
	case RC_PUNCTUATION:
	case RC_OPEN_PUNCTUATION:
	case RC_CLOSE_PUNCTUATION:
	case RC_LINE_END:
		lexIndentLineBegin(plex, ch)
	case RC_QUOTE:
	case RC_FORBIDDEN:
	}
}

// no need for lexCommentEnd

func lexWhiteSpaceSwitch(plex *Lexer, ch rune) {
	plex.state = TT_WHITESPACE
	lexWhitespace(plex, ch)
}
func lexWhitespaceBegin(plex *Lexer, ch rune) {
	plex.state = TT_WHITESPACE
	plex.currentToken = NewToken(plex, TT_WHITESPACE)
}
func lexWhitespace(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
		// ingnore ch
	case RC_LETTER:
		lexIdentBegin(plex, ch)
	case RC_NUMBER:
		lexNumStrBegin(plex, ch)
	case RC_PUNCTUATION:
		lexStandingSymbolBegin(plex, ch)
	case RC_OPEN_PUNCTUATION:
		lexOpenSymbolBegin(plex, ch)
	case RC_CLOSE_PUNCTUATION:
		lexCloseSymbolBegin(plex, ch)
	case RC_LINE_END:
		if len(plex.tokenList) > 0 && plex.tokenList[len(plex.tokenList)-1].TokenType != TT_LINE { // this prevents too many duplicates
			lexIndentLineBegin(plex, ch)
		}
	case RC_QUOTE:
		lexStrLitBegin(plex, ch)
	case RC_FORBIDDEN:
		lexSyntaxErrorSwitch(plex, ch, "forbidden rune")
	}
}

// no need for lexWhitespaceEnd

func lexIndentLineBegin(plex *Lexer, ch rune) {
	plex.state = TT_LINE
	plex.currentToken = NewToken(plex, TT_LINE)
}
func lexIndentLine(plex *Lexer, ch rune) {
	switch plex.runeCategory {
	case RC_WHITESPACE:
		plex.currentToken.Content = plex.currentToken.Content + string(ch)
	case RC_LETTER:
		lexIndentLineEnd(plex)
		lexWhiteSpaceSwitch(plex, ch)
	case RC_NUMBER:
		lexIndentLineEnd(plex)
		lexWhiteSpaceSwitch(plex, ch)
	case RC_PUNCTUATION:
		lexIndentLineEnd(plex)
		lexWhiteSpaceSwitch(plex, ch)
	case RC_OPEN_PUNCTUATION:
		lexIndentLineEnd(plex)
		lexWhiteSpaceSwitch(plex, ch)
	case RC_CLOSE_PUNCTUATION:
		lexIndentLineEnd(plex)
		lexWhiteSpaceSwitch(plex, ch)
	case RC_LINE_END:
		plex.currentToken.Content = ""
	case RC_QUOTE:
		lexIndentLineEnd(plex)
		lexWhiteSpaceSwitch(plex, ch)
	case RC_FORBIDDEN:
		lexIndentLineEnd(plex)
		lexSyntaxErrorSwitch(plex, ch, "forbidden rune")
	}
}
func lexIndentLineEnd(plex *Lexer) {
	plex.currentToken.Indent = len(plex.currentToken.Content) / 2
	plex.currentToken.Content = ""
	lastIndex := len(plex.tokenList) - 1
	if lastIndex >= 0 {
		if plex.tokenList[lastIndex].TokenType == TT_LINE {
			plex.tokenList[lastIndex] = plex.currentToken
			return
		}
	}
	plex.tokenList = append(plex.tokenList, plex.currentToken)
}

func lex(plex *Lexer, ch rune) error {
	plex.runeCategory = GetRuneCategory(ch)
	switch plex.state {
	case TT_STANDING_SYMBOL:
		lexStandingSymbol(plex, ch)
	case TT_OPEN_SYMBOL:
		lexOpenSymbol(plex, ch)
	case TT_CLOSE_SYMBOL:
		lexCloseSymbol(plex, ch)
	case TT_OPEN_BIND_SYMBOL:
		lexOpenBindSymbol(plex, ch)
	case TT_BINDING_SYMBOL:
		lexBindingSymbol(plex, ch)
	case TT_IDENT:
		lexIdent(plex, ch)
	case TT_STR_LIT:
		lexStrLit(plex, ch)
	case TT_NUMSTR_LIT:
		lexNumStr(plex, ch)
	case TT_SYNTAX_ERROR:
		// should not reach here
	case TT_COMMENT:
		lexComment(plex, ch)
	case TT_WHITESPACE:
		lexWhitespace(plex, ch)
	case TT_LINE:
		lexIndentLine(plex, ch)
	}
	if plex.state == TT_SYNTAX_ERROR {
		return fmt.Errorf("error found: %v", plex.currentToken)
	}
	if plex.runeCategory == RC_LINE_END {
		plex.currentLine += 1
		plex.currentOffset = 1
	} else {
		plex.currentOffset += 1
	}
	plex.previousRune = ch
	return nil
}

func LexString(source string) (error, []Token) {
	var err error
	plex := Lexer{
		state:         TT_WHITESPACE,
		currentLine:   1,
		currentOffset: 1,
		tokenList:     []Token{},
		previousRune:  0,
		fileId:        0,
	}
	plex.currentToken = NewToken(&plex, TT_LINE)
	for _, ch := range []rune(source) {
		err = lex(&plex, ch)
		if err != nil {
			break
		}
	}
	return err, plex.tokenList
}

func LexFile(cc *context.CompilerContext, target string) (error, []Token) {

	fileName := cc.RootDir + "/" + target + ".sulfur"
	fileHandle, err := os.Open(fileName)
	if err != nil {
		return err, nil
	}
	defer fileHandle.Close()

	fileId := context.SetNewFile(cc, target)

	plex := Lexer{
		state:         TT_WHITESPACE,
		currentLine:   1,
		currentOffset: 1,
		tokenList:     []Token{},
		previousRune:  0,
		fileId:        fileId,
	}
	plex.currentToken = NewToken(&plex, TT_LINE)

	reader := bufio.NewReader(fileHandle)
	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		err = lex(&plex, ch)
		if err != nil {
			break
		}
	}
	if cc.SaveLexedFlag {
		listCopy := plex.tokenList
		yamlData, err := yaml.Marshal(&listCopy)
		if err != nil {
			return err, nil
		}
		outPath := context.GetLexParseRoundPath(cc)
		yamlFilePath := outPath + "/" + fmt.Sprintf("file-%04d.token.yaml", fileId)
		err = os.WriteFile(yamlFilePath, yamlData, 0644)
		if err != nil {
			return err, nil
		}
	}
	return nil, plex.tokenList
}
