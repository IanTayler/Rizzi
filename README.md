This is an interpreter for (right now, a fragment of) the Rizzi programming language written in Go. 

Syntactically, this is how the fragment works. Terminals are between '':

	PROGRAM			-> MAIN '[' ID ']' STATEMENTLIST END

	STATEMENTLIST   -> STATEMENT STATEMENTLIST

	STATEMENT		-> EXPR | ORDER | ASSIGNMENT

	ASSIGNMENT		-> ID '=' EXPR

	ORDER			-> 'if' EXPR 'then' STATEMENTLIST 'end' |

					'for' EXPR 'do' STATEMENTLIST 'end'
		   
	EXPR			-> mathemathical expressions with +, *, -, integer /,

					rem, neg, succ, pred, exp, numbers and variables.
		   
	ID				-> variable names

	MAIN			-> 'main' | 'm'

	END				-> 'end' | 'e'

Statements should be on separate lines. The parser works fine when that's the case but I'm not entirely sure why (which is pretty bad).

A program returns its last statement. Running the interpreter like "rizzi FILE -a ARG" passes ARG as an argument. That argument is saved in a variable with the name between '[' and ']'.

There's some syntactic bugs which restrict where we can put IDs. Right now they can be circumvented by writing '0 + var' instead of 'var' when a problem arises, but that's going to be fixed eventually.
There's also syntactic bugs in the way we deal with parentheses.

#TODO (in order. FIFO):
- Fix syntactic bugs on variables and parentheses.
- Add a 'print' order.
- Add multiple variables to programs.
- Add a type system with:
  Short-term:
	i.   Integers and natural numbers (i.e. unsigned integers) as distinct types.
	ii.  Type inference.
	iii. Pairs with elements of arbitrary type.
  Mid- and long-term:
	iv.  Lists (perhaps we can use pairs and pointers to pairs internally)
	v.   Bytes
	vi.  Structs
  Perhaps:
	vi.  UTF-8 characters and UTF-8 strings.
- Add function declarations and functions as first-class types.
- Lexical scope.
- The Go API.
