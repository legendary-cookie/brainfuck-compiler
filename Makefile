all: release

release:
	go build -ldflags="-X 'main.Version=v1.0.0'"

test: all
	./brainfuck-compiler -i hello.b -o hello
