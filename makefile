EXE = africana
SRC = .
LDFLAGS = -ldflags="-s -w"

linux:
	GOOS=linux go build -v -x -o $(EXE)_linux.bin $(LDFLAGS) $(SRC)

windows:
	GOOS=windows go build -v -x -o $(EXE)_win.exe $(LDFLAGS) $(SRC)

macos:
	GOOS=darwin go build -v -x -o $(EXE)_macos.mc $(LDFLAGS) $(SRC)

all: windows macos linux
	echo "done."

clean:
	rm -f $(EXE)_win.exe $(EXE)_macos.mc $(EXE)_linux.bin
