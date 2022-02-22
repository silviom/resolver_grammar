%{
package main
import "fmt"
%}

%union{
String string
Expr expr 
}

%token<String> BEGIN END NAME OPEN CLOSE SPACE NOT_SPACE COMMA

%type <Expr> text stmt stmts resolver arguments argument argument_text

%%

start: 
  stmts {
    fmt.Printf("PARSE RESULT:\n%v\n", $1)
    resolverlex.(*interpreter).parseResult = $1
  } 
  ;

stmts:
  stmts stmt {
    stmts := $1.(*stmtsAST)
    stmts.list = append(stmts.list, $2.(*stmtAST))
    $$ = $1
    ;fmt.Printf(">> stmts | stmts stmt\n")
  }
  | stmt { 
    stmts := &stmtsAST{}
    stmts.list = append(stmts.list, $1.(*stmtAST))
    $$ = stmts
    ;fmt.Printf(">> stmts | stmt\n")
  }
  ;

stmt:
  text stmt {
    stmt := $2.(*stmtAST)
    stmt.before = $1.(string) + stmt.before
    $$ = stmt
    ;fmt.Printf(">> stmt | text stmt\n")
  }
  | stmt text {
    stmt := $1.(*stmtAST)
    stmt.after += $2.(string)
    $$ = stmt
    ;fmt.Printf(">> stmt | stmt text\n")
  }
  | BEGIN resolver END {
    $$ = &stmtAST{
      resolver: $2.(*resolverAST),
    }
    ;fmt.Printf(">> stmt | BEGIN resolver END \n")
  }
  | text {
    $$ = &stmtAST{
      before: $1.(string),
    }
    ;fmt.Printf(">> stmt | text \n")
  }
  ;

text:
    text SPACE      { $$ = $1.(string) + $2 ;fmt.Printf(">> text | text SPACE \n")}
  | text OPEN       { $$ = $1.(string) + $2 ;fmt.Printf(">> text | text OPEN \n")}
  | text CLOSE      { $$ = $1.(string) + $2 ;fmt.Printf(">> text | text CLOSE \n")}
  | text COMMA      { $$ = $1.(string) + $2 ;fmt.Printf(">> text | text COMMA \n")}
  | text NAME       { $$ = $1.(string) + $2 ;fmt.Printf(">> text | text NAME \n")}
  | text NOT_SPACE  { $$ = $1.(string) + $2 ;fmt.Printf(">> text | text NOT_SPACE \n")}
  | SPACE           { $$ = $1 ;fmt.Printf(">> text | SPACE \n")}
  | OPEN            { $$ = $1 ;fmt.Printf(">> text | OPEN \n")}
  | CLOSE           { $$ = $1 ;fmt.Printf(">> text | CLOSE \n")}
  | COMMA           { $$ = $1 ;fmt.Printf(">> text | COMMA \n")}
  | NAME            { $$ = $1 ;fmt.Printf(">> text | NAME \n")}
  | NOT_SPACE       { $$ = $1 ;fmt.Printf(">> text | NOT_SPACE \n")}
  ;

resolver:
  SPACE resolver    {
    $$ = $2
   ;fmt.Printf(">> resolver | SPACE resolver\n")
  }
  | resolver SPACE  {
    $$ = $1
   ;fmt.Printf(">> resolver | resolver SPACE\n")
  }
  | NAME SPACE OPEN arguments CLOSE {
    $$ = &resolverAST{
      name: $1,
      arguments: $4.(*argsAST),
    } 
   ;fmt.Printf(">> resolver |NAME SPACE OPEN arguments CLOSE \n")
  }
  | NAME OPEN arguments CLOSE {
    $$ = &resolverAST{
      name: $1,
      arguments: $3.(*argsAST),
    }
   ;fmt.Printf(">> resolver | NAME OPEN arguments CLOSE\n")
  }
  ;

arguments:
  {
    $$ = &argsAST{}
    ;fmt.Printf(">> args | empty\n")
  }
  | argument {
    args := &argsAST{}
    args.list = append(args.list, $1.(*argAST))
    $$ = args
    ;fmt.Printf(">> args | argument\n")
  }
  | arguments COMMA argument {
    args := $1.(*argsAST)
    args.list = append(args.list, $3.(*argAST))
    $$ = args
   ;fmt.Printf(">> arguments | arguments COMMA argument\n")
  }
  ;
  
argument:
  argument_text argument {
    arg := $2.(*argAST)
    arg.before = $1.(string) + arg.before
    $$ = arg
    ;fmt.Printf(">> argument | argument_text argument\n")
  }
  | argument argument_text {
    arg := $1.(*argAST)
    arg.after += $2.(string)
    $$ = arg
    ;fmt.Printf(">> argument | argument argument_text\n")
  }
  | BEGIN resolver END {
    $$ = &argAST{
      resolver: $2.(*resolverAST),
    }
    ;fmt.Printf(">> argument | BEGIN resolver END \n")
  }
  | argument_text {
    $$ = &argAST{
      before: $1.(string),
    }
    ;fmt.Printf(">> argument | argument_text \n")
  }
  ;

argument_text:
  argument_text SPACE      { $$ = $1.(string) + $2 ;fmt.Printf(">> argument_text | argument_text SPACE\n")}
  | argument_text NAME       { $$ = $1.(string) + $2 ;fmt.Printf(">> argument_text | argument_text NAME\n")}
  | argument_text NOT_SPACE  { $$ = $1.(string) + $2 ;fmt.Printf(">> argument_text | argument_text NOT_SPACE\n")}
  | SPACE      { $$ = $1 ;fmt.Printf(">> argument_text | SPACE\n")}
  | NAME       { $$ = $1 ;fmt.Printf(">> argument_text | NAME\n")}
  | NOT_SPACE  { $$ = $1 ;fmt.Printf(">> argument_text | NOT_SPACE\n")}
  ;
%%