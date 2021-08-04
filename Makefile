all: brainfuck-compiler examples

brainfuck-compiler:
	go build -ldflags="-X 'main.Version=v1.0.0'"

examples: helloworld rot13 caesar-cipher

helloworld: examples/helloworld.b
	./brainfuck-compiler -i $< -o $@

rot13: examples/rot13.b
	./brainfuck-compiler -i $< -o $@

fibonacci: examples/fibonacci.b
	./brainfuck-compiler -i $< -o $@

caesar-cipher: examples/caesar-cipher.b
	./brainfuck-compiler -i $< -o $@

.PHONY: clean
clean:
	find . -maxdepth 1 -type f -perm -100 -delete
	find . -type f -name "*.o" -delete
