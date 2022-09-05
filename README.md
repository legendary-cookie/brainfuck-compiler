# Brainfuck-compiler
## How this works
The code iterates through the brainfuck file and converts any brainfuck stuff into C-code. The resulting code is written to /tmp/brainfuck.c, 
I then shell out to gcc and compile it.

## Reference (Brainfuck -> C)
The C code will have a ```char array[65535] = {0}; char *ptr = array;``` at the beginning of the ```int main()``` function.

|Brainfuck|C-code|
|---------|------|
|>        |++ptr;|
|<        |--ptr;|
|+        |++*ptr;|
|-        |--*ptr;|
|.        |putchar(*ptr);|
|,        |*ptr = getchar();|
|[        |while (*ptr) {|
|]        |}|
