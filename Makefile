BIN_PATH=./bin

run:
	go run .

build: clean
	go build -o $(BIN_PATH)/agg .

clean:
	rm -f $(BIN_PATH)/agg
