package token

type TokenType int

const (
	// Mots-clés
	ACTION TokenType = iota + 1000
	START
	STOP
	END
	IF
	ELSE
	WHILE
	FOR
	FOREACH
	SWITCH
	CASE
	DEFAULT
	LET
	CONST
	STRUCT
	FUNCTION
	RETURN
	SELECT
	FROM
	WHERE
	ORDER
	BY
	ASC
	DESC
	LIMIT
	OFFSET
	AS
	INTO
	OBJET
	INNER
	LEFT
	RIGHT
	JOIN
	ON
	GROUP
	HAVING
	LIKE
	BETWEEN
	IS
	NULL
	UNION
	ALL
	WITH
	RECURSIVE
	DISTINCT
	CREATE
	DROP
	DELETE
	UPDATE
	SAVE
	SET
	VIEW
	INDEX
	SEQUENCE
	NOT
	DEFAULT_KW
	INTEGER
	BOOLEAN
	FLOAT_TYPE
	STRING_TYPE
	DATE
	TIME
	DATETIME
	DURATION
	ARRAY
	DIGITS
	LENGTH
	MIN
	MAX
	RANGE
	INTEGER_DIGITS
	DECIMAL_DIGITS
	OPTIONAL
	SELF
	YEARS
	MONTHS
	DAYS
	HOURS
	MINUTES
	SECONDS
	MILLISECONDS
	TRUE
	FALSE
	AND
	OR
	NOT_OP
	IN
	NEWLINE
	EOF

	// Identifiants et littéraux
	IDENTIFIER
	NUMBER
	FLOAT
	STRING
	DATE_LITERAL
	TIME_LITERAL
	DURATION_LITERAL

	// Opérateurs et ponctuation
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	MODULO
	ASSIGN
	EQUAL
	NOT_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET
	SEMICOLON
	COMMA
	COLON
	DOT
)

type Token struct {
	Type  TokenType
	Value interface{}
	Line  int
	Pos   int
}

var tokenNames = map[TokenType]string{
	ACTION:         "ACTION",
	START:          "START",
	STOP:           "STOP",
	END:            "END",
	IF:             "IF",
	ELSE:           "ELSE",
	WHILE:          "WHILE",
	FOR:            "FOR",
	FOREACH:        "FOREACH",
	SWITCH:         "SWITCH",
	CASE:           "CASE",
	DEFAULT:        "DEFAULT",
	LET:            "LET",
	CONST:          "CONST",
	STRUCT:         "STRUCT",
	FUNCTION:       "FUNCTION",
	RETURN:         "RETURN",
	SELECT:         "SELECT",
	FROM:           "FROM",
	WHERE:          "WHERE",
	ORDER:          "ORDER",
	BY:             "BY",
	ASC:            "ASC",
	DESC:           "DESC",
	LIMIT:          "LIMIT",
	OFFSET:         "OFFSET",
	AS:             "AS",
	INTO:           "INTO",
	OBJET:          "OBJET",
	INNER:          "INNER",
	LEFT:           "LEFT",
	RIGHT:          "RIGHT",
	JOIN:           "JOIN",
	ON:             "ON",
	GROUP:          "GROUP",
	HAVING:         "HAVING",
	LIKE:           "LIKE",
	BETWEEN:        "BETWEEN",
	IS:             "IS",
	NULL:           "NULL",
	UNION:          "UNION",
	ALL:            "ALL",
	WITH:           "WITH",
	RECURSIVE:      "RECURSIVE",
	DISTINCT:       "DISTINCT",
	CREATE:         "CREATE",
	DROP:           "DROP",
	DELETE:         "DELETE",
	UPDATE:         "UPDATE",
	SAVE:           "SAVE",
	SET:            "SET",
	VIEW:           "VIEW",
	INDEX:          "INDEX",
	SEQUENCE:       "SEQUENCE",
	NOT:            "NOT",
	DEFAULT_KW:     "DEFAULT",
	INTEGER:        "INTEGER",
	BOOLEAN:        "BOOLEAN",
	FLOAT_TYPE:     "FLOAT_TYPE",
	STRING_TYPE:    "STRING_TYPE",
	DATE:           "DATE",
	TIME:           "TIME",
	DATETIME:       "DATETIME",
	DURATION:       "DURATION",
	ARRAY:          "ARRAY",
	DIGITS:         "DIGITS",
	LENGTH:         "LENGTH",
	MIN:            "MIN",
	MAX:            "MAX",
	RANGE:          "RANGE",
	INTEGER_DIGITS: "INTEGER_DIGITS",
	DECIMAL_DIGITS: "DECIMAL_DIGITS",
	OPTIONAL:       "OPTIONAL",
	SELF:           "SELF",
	YEARS:          "YEARS",
	MONTHS:         "MONTHS",
	DAYS:           "DAYS",
	HOURS:          "HOURS",
	MINUTES:        "MINUTES",
	SECONDS:        "SECONDS",
	MILLISECONDS:   "MILLISECONDS",
	TRUE:           "TRUE",
	FALSE:          "FALSE",
	AND:            "AND",
	OR:             "OR",
	NOT_OP:         "NOT_OP",
	IN:             "IN",
	NEWLINE:        "NEWLINE",
	EOF:            "EOF",

	IDENTIFIER:       "IDENTIFIER",
	NUMBER:           "NUMBER",
	FLOAT:            "FLOAT",
	STRING:           "STRING",
	DATE_LITERAL:     "DATE_LITERAL",
	TIME_LITERAL:     "TIME_LITERAL",
	DURATION_LITERAL: "DURATION_LITERAL",

	PLUS:          "+",
	MINUS:         "-",
	MULTIPLY:      "*",
	DIVIDE:        "/",
	MODULO:        "%",
	ASSIGN:        "=",
	EQUAL:         "==",
	NOT_EQUAL:     "!=",
	LESS:          "<",
	LESS_EQUAL:    "<=",
	GREATER:       ">",
	GREATER_EQUAL: ">=",
	LPAREN:        "(",
	RPAREN:        ")",
	LBRACE:        "{",
	RBRACE:        "}",
	LBRACKET:      "[",
	RBRACKET:      "]",
	SEMICOLON:     ";",
	COMMA:         ",",
	COLON:         ":",
	DOT:           ".",
}

func (t TokenType) String() string {
	if name, ok := tokenNames[t]; ok {
		return name
	}
	return "UNKNOWN"
}

func (t Token) String() string {
	return tokenNames[t.Type]
}
