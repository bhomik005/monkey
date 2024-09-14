## Monkey (An Interpreter) - A study project
### Referring to this amazing book written by Thorsten Ball - https://interpreterbook.com/

### Rough Notes - 

Interpreters - The programs that take other programs as their input and produce something.  
There is a huge difference between the interpreters and compilers.
 - Interpreters take source code and evaluate it without producing some visible, intermediate result that can later be 
      executed.
 - Compilers take source code and convert that into another language which underlying system can understand.

Interpreters (tree walking) - parse the source code, build an abstract syntax tree (AST) and then evaluate it.

Monkey programming language will have the following features - 
 
- C-like syntax
- variable bindings
- integers and expressions
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data structure

### MONKEY PROGRAMMING LANGUAGE SYNTAX 
1. Binding values to name in Monkey ( var declaration )  
   ```
   let age = 1;  
   let name = "Monkey";  
   let result = 10 * (20 / 2);
   ```
2. Arrays and hashes  
   ```
   let myArray = [1, 2, 3, 4];  
   let thorsten = {"name": "Thorsten", "age": 28};
   ```
3. Accessing the elements in arrays and hashes is done with index expressions  
   ```
   myArray[0]         // => 1   
   thorsten["name]    // => "Thorsten"
   ```
4. The let statements can also be used to bind functions to names. Here is a small functions that adds two numbers  
   ```
   let add = fn(a, b) { return a + b; }; 
   ```
5. The implicit return statements are also possible which means we can leave out the return if we want to  
   ```
   let add = fn(a, b) { a + b; };
   ```
6. Calling a function in monkey  
   ```
   add(1, 2);
   ```
   
7. A more complex function ( recursive function calls in it )  
   ```
   let fibonacci = fn(x) {
     if (x == 0) {
       0
     } else {
       if (x == 1) {
         1
       } else {
         fibonacci(x - 1) + fibonacci(x - 2);
       }
     }
   };
   ```
8. Monkey also supports special type of functions called higher order functions. The function that takes other functions
   as arguments
   ```
   let twice = fn(f,x) {
     return f(f(x));
   };
   
   let addTwo = fn(x) {
     return x + 2;
   };
   
   twice(addTwo, 2); // => 6 (calling addTwo function twice)
   ```

We are going to build following parts / components 
1. Lexer
2. Parser
3. Abstract Syntax Tree
4. The internal object system
5. The evaluator


  