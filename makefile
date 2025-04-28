#Africana-Framework builder. Run make

SRC = .
EXE = afrconsole
LDFLAGS = -ldflags="-s -w"

linux:
	GOOS=linux go build -v -x -o $(EXE) $(LDFLAGS) $(SRC)
	echo "[+] building completed for linux distros ..."

macos:
	GOOS=darwin go build -v -x -o $(EXE) $(LDFLAGS) $(SRC)
	echo "[+] building completed for darwin distros ..."

windows:
	GOOS=windows go build -v -x -o $(EXE).exe $(LDFLAGS) $(SRC)
	echo "[+] building completed for windows distros ..."

all: linux macos windows
	echo "[+] building completed for all supported distros ..."

clean:
	rm -f $(EXE).exe $(EXE)
	echo "[+] cleaning completed ..."
