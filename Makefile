all:
	go build

test:
	./brainfuck-compiler hello.b hello
