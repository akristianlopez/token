package token

/**
 * This class implements a node of the list that returned by the scanner after it performing.
 * This node is used by the lexer during its performing
 * @author Christian Lopez ATOUBA
 */

// Token représente un token lexical
type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

// TokenType représente le type de token
type TokenType int

const (
	// Tokens spéciaux
	TOKEN_EOF TokenType = iota
	TOKEN_NEWLINE
	TOKEN_ERROR

	// Mots-clés principaux
	TOKEN_ACTION
	TOKEN_START
	TOKEN_STOP
	TOKEN_END

	// Structures de contrôle
	TOKEN_IF
	TOKEN_ELSE
	TOKEN_WHILE
	TOKEN_FOR
	TOKEN_FOREACH
	TOKEN_SWITCH
	TOKEN_CASE
	TOKEN_DEFAULT
	TOKEN_DO
	TOKEN_THEN

	// Déclarations
	TOKEN_LET
	TOKEN_CONST
	TOKEN_STRUCT
	TOKEN_FUNCTION
	TOKEN_RETURN
	TOKEN_TYPE

	// SQL et bases de données
	TOKEN_SELECT
	TOKEN_FROM
	TOKEN_WHERE
	TOKEN_ORDER
	TOKEN_BY
	TOKEN_ASC
	TOKEN_DESC
	TOKEN_LIMIT
	TOKEN_OFFSET
	TOKEN_AS
	TOKEN_INTO
	TOKEN_OBJET
	TOKEN_INNER
	TOKEN_LEFT
	TOKEN_RIGHT
	TOKEN_JOIN
	TOKEN_ON
	TOKEN_GROUP
	TOKEN_HAVING

	// Opérateurs SQL
	TOKEN_LIKE
	TOKEN_BETWEEN
	TOKEN_IS
	TOKEN_NULL
	TOKEN_UNION
	TOKEN_ALL
	TOKEN_WITH
	TOKEN_RECURSIVE
	TOKEN_DISTINCT

	// DDL/DML
	TOKEN_CREATE
	TOKEN_DROP
	TOKEN_DELETE
	TOKEN_UPDATE
	TOKEN_SAVE
	TOKEN_SET
	TOKEN_VIEW
	TOKEN_INDEX
	TOKEN_SEQUENCE
	TOKEN_NOT
	TOKEN_RARROW
	TOKEN_LARROW
	// TOKEN_DEFAULT

	// Types de base
	TOKEN_INTEGER
	TOKEN_BOOLEAN
	TOKEN_FLOAT_TYPE
	TOKEN_STRING_TYPE
	TOKEN_DATE
	TOKEN_TIME
	TOKEN_DATETIME
	TOKEN_DURATION
	TOKEN_ARRAY

	// Contraintes de types
	TOKEN_DIGITS
	TOKEN_LENGTH
	TOKEN_MIN
	TOKEN_MAX
	TOKEN_RANGE
	TOKEN_INTEGER_DIGITS
	TOKEN_DECIMAL_DIGITS

	// Structures récursives
	TOKEN_OPTIONAL
	TOKEN_SELF

	// Unités de durée
	TOKEN_YEARS
	TOKEN_MONTHS
	TOKEN_DAYS
	TOKEN_HOURS
	TOKEN_MINUTES
	TOKEN_SECONDS
	TOKEN_MILLISECONDS

	// Valeurs littérales
	TOKEN_TRUE
	TOKEN_FALSE

	// Opérateurs logiques
	TOKEN_AND
	TOKEN_OR
	TOKEN_NOT_OP

	// Autres mots-clés
	TOKEN_IN

	// Littéraux
	TOKEN_IDENTIFIER
	TOKEN_NUMBER
	TOKEN_FLOAT
	TOKEN_STRING
	TOKEN_DATE_LITERAL
	TOKEN_TIME_LITERAL
	TOKEN_DURATION_LITERAL

	// Opérateurs
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_MULTIPLY
	TOKEN_DIVIDE
	TOKEN_MODULO
	TOKEN_ASSIGN
	TOKEN_EQUAL
	TOKEN_NOT_EQUAL
	TOKEN_LESS
	TOKEN_LESS_EQUAL
	TOKEN_GREATER
	TOKEN_GREATER_EQUAL

	// Délimiteurs
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_LBRACE
	TOKEN_RBRACE
	TOKEN_LBRACKET
	TOKEN_RBRACKET
	TOKEN_SEMICOLON
	TOKEN_COMMA
	TOKEN_COLON
	TOKEN_DOT
	TOKEN_LESS_ANGLE
	TOKEN_GREATER_ANGLE
)

// String retourne la représentation textuelle du token
func (t TokenType) String() string {
	return tokenNames[t]
}

var tokenNames = map[TokenType]string{
	TOKEN_EOF:     "EOF",
	TOKEN_NEWLINE: "NEWLINE",
	TOKEN_ERROR:   "ERROR",

	TOKEN_ACTION: "ACTION",
	TOKEN_START:  "START",
	TOKEN_STOP:   "STOP",
	TOKEN_END:    "END",

	TOKEN_IF:      "IF",
	TOKEN_ELSE:    "ELSE",
	TOKEN_WHILE:   "WHILE",
	TOKEN_FOR:     "FOR",
	TOKEN_FOREACH: "FOREACH",
	TOKEN_SWITCH:  "SWITCH",
	TOKEN_CASE:    "CASE",
	TOKEN_DEFAULT: "DEFAULT",
	TOKEN_DO:      "DO",
	TOKEN_THEN:    "THEN",
	TOKEN_LARROW:  "<-",
	TOKEN_RARROW:  "->",

	TOKEN_LET:      "LET",
	TOKEN_CONST:    "CONST",
	TOKEN_STRUCT:   "STRUCT",
	TOKEN_FUNCTION: "FUNCTION",
	TOKEN_RETURN:   "RETURN",
	TOKEN_TYPE:     "type",

	TOKEN_SELECT: "SELECT",
	TOKEN_FROM:   "FROM",
	TOKEN_WHERE:  "WHERE",
	TOKEN_ORDER:  "ORDER",
	TOKEN_BY:     "BY",
	TOKEN_ASC:    "ASC",
	TOKEN_DESC:   "DESC",
	TOKEN_LIMIT:  "LIMIT",
	TOKEN_OFFSET: "OFFSET",
	TOKEN_AS:     "AS",
	TOKEN_INTO:   "INTO",
	TOKEN_OBJET:  "OBJET",
	TOKEN_INNER:  "INNER",
	TOKEN_LEFT:   "LEFT",
	TOKEN_RIGHT:  "RIGHT",
	TOKEN_JOIN:   "JOIN",
	TOKEN_ON:     "ON",
	TOKEN_GROUP:  "GROUP",
	TOKEN_HAVING: "HAVING",

	TOKEN_LIKE:      "LIKE",
	TOKEN_BETWEEN:   "BETWEEN",
	TOKEN_IS:        "IS",
	TOKEN_NULL:      "NULL",
	TOKEN_UNION:     "UNION",
	TOKEN_ALL:       "ALL",
	TOKEN_WITH:      "WITH",
	TOKEN_RECURSIVE: "RECURSIVE",
	TOKEN_DISTINCT:  "DISTINCT",

	TOKEN_CREATE:   "CREATE",
	TOKEN_DROP:     "DROP",
	TOKEN_DELETE:   "DELETE",
	TOKEN_UPDATE:   "UPDATE",
	TOKEN_SAVE:     "SAVE",
	TOKEN_SET:      "SET",
	TOKEN_VIEW:     "VIEW",
	TOKEN_INDEX:    "INDEX",
	TOKEN_SEQUENCE: "SEQUENCE",

	TOKEN_INTEGER:     "INTEGER",
	TOKEN_BOOLEAN:     "BOOLEAN",
	TOKEN_FLOAT_TYPE:  "FLOAT_TYPE",
	TOKEN_STRING_TYPE: "STRING_TYPE",
	TOKEN_DATE:        "DATE",
	TOKEN_TIME:        "TIME",
	TOKEN_DATETIME:    "DATETIME",
	TOKEN_DURATION:    "DURATION",
	TOKEN_ARRAY:       "ARRAY",

	TOKEN_DIGITS:         "DIGITS",
	TOKEN_LENGTH:         "LENGTH",
	TOKEN_MIN:            "MIN",
	TOKEN_MAX:            "MAX",
	TOKEN_RANGE:          "RANGE",
	TOKEN_INTEGER_DIGITS: "INTEGER_DIGITS",
	TOKEN_DECIMAL_DIGITS: "DECIMAL_DIGITS",

	TOKEN_OPTIONAL: "OPTIONAL",
	TOKEN_SELF:     "SELF",

	TOKEN_YEARS:        "YEARS",
	TOKEN_MONTHS:       "MONTHS",
	TOKEN_DAYS:         "DAYS",
	TOKEN_HOURS:        "HOURS",
	TOKEN_MINUTES:      "MINUTES",
	TOKEN_SECONDS:      "SECONDS",
	TOKEN_MILLISECONDS: "MILLISECONDS",

	TOKEN_TRUE:  "TRUE",
	TOKEN_FALSE: "FALSE",

	TOKEN_AND:    "AND",
	TOKEN_OR:     "OR",
	TOKEN_NOT_OP: "NOT_OP",

	TOKEN_IN: "IN",

	TOKEN_IDENTIFIER:       "IDENTIFIER",
	TOKEN_NUMBER:           "NUMBER",
	TOKEN_FLOAT:            "FLOAT",
	TOKEN_STRING:           "STRING",
	TOKEN_DATE_LITERAL:     "DATE_LITERAL",
	TOKEN_TIME_LITERAL:     "TIME_LITERAL",
	TOKEN_DURATION_LITERAL: "DURATION_LITERAL",

	TOKEN_PLUS:          "PLUS",
	TOKEN_MINUS:         "MINUS",
	TOKEN_MULTIPLY:      "MULTIPLY",
	TOKEN_DIVIDE:        "DIVIDE",
	TOKEN_MODULO:        "MODULO",
	TOKEN_ASSIGN:        "ASSIGN",
	TOKEN_EQUAL:         "EQUAL",
	TOKEN_NOT_EQUAL:     "NOT_EQUAL",
	TOKEN_LESS:          "LESS",
	TOKEN_LESS_EQUAL:    "LESS_EQUAL",
	TOKEN_GREATER:       "GREATER",
	TOKEN_GREATER_EQUAL: "GREATER_EQUAL",

	TOKEN_LPAREN:        "LPAREN",
	TOKEN_RPAREN:        "RPAREN",
	TOKEN_LBRACE:        "LBRACE",
	TOKEN_RBRACE:        "RBRACE",
	TOKEN_LBRACKET:      "LBRACKET",
	TOKEN_RBRACKET:      "RBRACKET",
	TOKEN_SEMICOLON:     "SEMICOLON",
	TOKEN_COMMA:         "COMMA",
	TOKEN_COLON:         "COLON",
	TOKEN_DOT:           "DOT",
	TOKEN_LESS_ANGLE:    "LESS_ANGLE",
	TOKEN_GREATER_ANGLE: "GREATER_ANGLE",
}

// Keywords map des mots-clés
var Keywords = map[string]TokenType{
	"action": TOKEN_ACTION,
	"start":  TOKEN_START,
	"stop":   TOKEN_STOP,
	"end":    TOKEN_END,

	"if":      TOKEN_IF,
	"else":    TOKEN_ELSE,
	"while":   TOKEN_WHILE,
	"for":     TOKEN_FOR,
	"foreach": TOKEN_FOREACH,
	"switch":  TOKEN_SWITCH,
	"case":    TOKEN_CASE,
	"default": TOKEN_DEFAULT,

	"let":      TOKEN_LET,
	"const":    TOKEN_CONST,
	"struct":   TOKEN_STRUCT,
	"function": TOKEN_FUNCTION,
	"return":   TOKEN_RETURN,

	"select": TOKEN_SELECT,
	"from":   TOKEN_FROM,
	"where":  TOKEN_WHERE,
	"order":  TOKEN_ORDER,
	"by":     TOKEN_BY,
	"asc":    TOKEN_ASC,
	"desc":   TOKEN_DESC,
	"limit":  TOKEN_LIMIT,
	"offset": TOKEN_OFFSET,
	"as":     TOKEN_AS,
	"into":   TOKEN_INTO,
	"objet":  TOKEN_OBJET,
	"inner":  TOKEN_INNER,
	"left":   TOKEN_LEFT,
	"right":  TOKEN_RIGHT,
	"join":   TOKEN_JOIN,
	"on":     TOKEN_ON,
	"group":  TOKEN_GROUP,
	"having": TOKEN_HAVING,

	"like":      TOKEN_LIKE,
	"between":   TOKEN_BETWEEN,
	"is":        TOKEN_IS,
	"null":      TOKEN_NULL,
	"union":     TOKEN_UNION,
	"all":       TOKEN_ALL,
	"with":      TOKEN_WITH,
	"recursive": TOKEN_RECURSIVE,
	"distinct":  TOKEN_DISTINCT,

	"create":   TOKEN_CREATE,
	"drop":     TOKEN_DROP,
	"delete":   TOKEN_DELETE,
	"update":   TOKEN_UPDATE,
	"save":     TOKEN_SAVE,
	"set":      TOKEN_SET,
	"view":     TOKEN_VIEW,
	"index":    TOKEN_INDEX,
	"sequence": TOKEN_SEQUENCE,
	"not":      TOKEN_NOT,
	// "default":  TOKEN_DEFAULT,

	"integer":  TOKEN_INTEGER,
	"boolean":  TOKEN_BOOLEAN,
	"float":    TOKEN_FLOAT_TYPE,
	"string":   TOKEN_STRING_TYPE,
	"date":     TOKEN_DATE,
	"time":     TOKEN_TIME,
	"datetime": TOKEN_DATETIME,
	"duration": TOKEN_DURATION,
	"array":    TOKEN_ARRAY,

	"digits":         TOKEN_DIGITS,
	"length":         TOKEN_LENGTH,
	"min":            TOKEN_MIN,
	"max":            TOKEN_MAX,
	"range":          TOKEN_RANGE,
	"integer_digits": TOKEN_INTEGER_DIGITS,
	"decimal_digits": TOKEN_DECIMAL_DIGITS,

	"optional": TOKEN_OPTIONAL,
	"self":     TOKEN_SELF,

	"years":        TOKEN_YEARS,
	"months":       TOKEN_MONTHS,
	"days":         TOKEN_DAYS,
	"hours":        TOKEN_HOURS,
	"minutes":      TOKEN_MINUTES,
	"seconds":      TOKEN_SECONDS,
	"milliseconds": TOKEN_MILLISECONDS,

	"true":  TOKEN_TRUE,
	"false": TOKEN_FALSE,

	"And": TOKEN_AND,
	"Or":  TOKEN_OR,
	"Not": TOKEN_NOT_OP,

	"in": TOKEN_IN,
}
