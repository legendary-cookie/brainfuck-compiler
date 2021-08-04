all: brainfuck-compiler examples

brainfuck-compiler:
	go build -ldflags="-X 'main.Version=v1.0.0'"

examples: helloworld rot13

helloworld: brainfuck-compiler examples/helloworld.b
	./brainfuck-compiler -i examples/helloworld.b -o helloworld

rot13: brainfuck-compiler examples/rot13.b
	./brainfuck-compiler -i examples/rot13.b -o rot13

.PHONY: clean
clean:
	find . -maxdepth 1 -type f -perm -100 -delete
	find . -type f -name "*.o" -delete
