#Africana-Framework builder. Run make

SRC = .
EXE = afrconsole
LDFLAGS = -ldflags="-s -w"

linux:
	GOOS=linux go build -v -x -o $(EXE) $(LDFLAGS) $(SRC)

macos:
	GOOS=darwin go build -v -x -o $(EXE).mac $(LDFLAGS) $(SRC)

windows:
	GOOS=windows go build -v -x -o $(EXE).exe $(LDFLAGS) $(SRC)

all: linux macos windows

clean:
	rm -f $(EXE).exe $(EXE).mac $(EXE)
