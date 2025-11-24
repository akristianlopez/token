package token

import (
	"fmt"
)

/**
 * This class implements a node of the list that returned by the scanner after it performing.
 * This node is used by the lexer during its performing
 * @author Christian Lopez ATOUBA
 */

type TokenType string

type Token struct {
	CurLine     int // current line
	CurCol     int // current column ;
	// CurPos    int //Current cursor position
	// WLen    int //Word length
	Token TokenType
	Value string
}

func (t *Token) SetPos(x, y int) {
	t.CurLine = x
	t.CurCol = y
}
func traitsChar(s string) string {
	switch {
	case s == "\r":
		return "\\r"
	case s == "\n":
		return "\\n"
	case s == "":
		return "''"
	}
	return s
}
func traitTv(s TokenType) TokenType {
	switch {
	case s == "\r":
		return "\\r"
	case s == "\n":
		return "\\n"
	case s == "":
		return "''"
	}
	return s
}
func (t *Token) ToString() string {
	return fmt.Sprintf("%1s [%2s]", traitsChar(t.Value), traitTv(t.Token))
}

const (
	UNDEFINED = "UNDEFINED" //Unkown token
	EOF       = "EOF"       //End of file
	EOL       = "\r"        // End of Line
	STAT_CONT = "_"         //Statement continue to the line below

	// Identifiers + literals
	IDENT  = "IDENT"   // add, foobar, x, y, ...
	NUMBER = "NUMBER"  // 1343456
	STRING = "STRING"  //' "mom is in the kitchen"
	BOOL   = "BOOLEAN" // true or false
	DATE   = "DATE"
	TIMES  = "TIME"

	COMMENT = "//" // Is being to delete later

	// Operators
	ASSIGN        = "="  //assignment
	PLUS          = "+"  //addition
	TIME          = "*"  //multiplication
	DIVIDE        = "/"  //division
	MINUS         = "-"  //Substraction
	GREATER       = ">"  //Greater
	GREATER_EQUAL = ">=" //Greater or equal
	LOWER         = "<"  // lower
	LOWER_EQUAL   = "<=" // lower or equal
	EQUAL         = "==" //equal
	NOT_EQUAL     = "<>" // not equal
	AND           = "AND"
	OR            = "OR"
	NOT           = "NOT"
	DOT           = "." // Dot
	LEFT          = "?" //Left join in the Fetch statement
	RENAME        = "@"
	IN            = "IN"
	BETWEEN       = "BETWEEN"
	OF            = "OF"
	RARROW        = "->"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	COLON     = ":"
	LBRAKET   = "["
	RBRAKET   = "]"

	// LBRACE    = "{"
	// RBRACE    = "}"

	// Keywords
	FUNC   = "FUNCTION"
	LET    = "LET"
	TYPE   = "TYPE"
	RECORD = "RECORD"
	IF     = "IF"
	THEN   = "THEN"
	ELSE   = "ELSE"
	END    = "END"
	STOP   = "STOP"
	ACTION = "ACTION"
	START  = "START"
	CONST  = "CONST"
	PARAM  = "PARAMETER"
	FOR    = "FOR"
	WHILE  = "WHILE"
	SELECT = "SELECT"
	BROWSE = "BROWSE"
	BEGIN  = "BEGIN"
	WHERE  = "WHERE"
	BREAK  = "BREAK"
	RETURN = "RETURN"
	REMOVE = "REMOVE"
	ADD    = "ADD"
	ALTER  = "ALTER"
	NULL   = "NULL"
	DO     = "DO"
	FROM   = "FROM"
	TAKE   = "TAKE"
)




