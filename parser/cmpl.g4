grammar cmpl;

policy     : default* rule+ ;
default    : 'DEFAULT ' decision ';';
rule       : decision ( ' WITH ' parameters )? ' WHEN ' conditions ( ' PRIORITY ' priority )? ';' ;

decision   : type assignment ;
item       : key ':' val ;
parameters : parameter (',' parameter)* ;
conditions : condition (lops condition)* ;
condition  : expression (lops expression)* | expression (lops condition)* | LPAREN condition RPAREN | ;
expression : lhs OPS rhs ;

type       : SCOPE | ACTION | FACT ;
lops       : AND | OR ;
assignment : key ':' val ;
parameter  : key ':' val ;
priority   : NUM ;

lhs        : ID ('.' ID)* ;
rhs        : ID | NUM | STR ;
key        : ID ('.' ID)* ;
val        : ID | NUM | STR ;

AND        : 'AND' ;
OR         : 'OR' ;

SCOPE      : 'GRANT $scope AS ' ;
ACTION     : 'EXEC $action AS ' ;
FACT       : 'ADD $fact AS ' ;

ID         : [$]*[a-zA-Z_][a-zA-Z_0-9]* ;
STR        : '"' (~["\\] | '\\' .)* '"' ;
NUM        : [0-9]+ ;
OPS        : '=' | '!=' | '<' | '>' | '<=' | '>=' ;
LPAREN     : '(' ;
RPAREN     : ')' ;

WS         : [ \t\r\n]+ -> skip ;