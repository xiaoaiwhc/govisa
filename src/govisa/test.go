// +build windows

package main

import (
	"bufio"
	"flag"
	"fmt"
	"govisa/visa"
	"io"
	"os"
	"strings"
	"time"
)

var board_id = flag.Uint("id", 0, "The Board ID")
var board_port = flag.Uint("port", 1, "GPIB port number")
var visa_cmd = flag.String("cmd", "*IDN?", "VISA command")
var file_path = flag.String("path", "", "The file path that includes GPIB commands.")

func skipLine(str string) bool {
	//blank line
	if len(str) == 0 {
		return true
	}
	//comment
	return strings.HasPrefix(str, "//")
}

func trimStr(buf []byte) string {
	return strings.TrimSpace(string(visa.TrimByteSlice(buf)))
}

func getCmdFromFile(filePath string) (buf []string) {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		line = strings.TrimSpace(line)
		if skipLine(line) {
			continue
		}

		buf = append(buf, line)
	}

	return
}

func executeCmd(addr string, cmd []string) {
	var (
		defaultRM uint32
		instr     uint32

		retCnt uint32

		resp []byte
	)

	resp = make([]byte, 1024)

	visa.ViOpenDefaultRM(&defaultRM)
	visa.ViOpen(defaultRM, addr, uint32(0), uint32(0), &instr)
	defer visa.ViClose(instr)

	for _, line := range cmd {
		c := strings.SplitN(line, ",", 2)
		c[1] = c[1] + "\n"
		switch c[0] {
		case "W":
			visa.ViWrite(instr, c[1], uint32(len(c[1])), &retCnt)
			fmt.Printf(">> %s", c[1])
		case "R":
			visa.ViRead(instr, resp, uint32(len(resp)), &retCnt)
			fmt.Printf("<< %s\n", trimStr(resp))
		case "Q":
//			fmt.Printf(">> %s", c[1])
//			result := new(string)
//			visa.ViQueryf(instr, c[1], "%s\n", result)
//			fmt.Printf("<< %s", *result)

			visa.ViWrite(instr, c[1], uint32(len(c[1])), &retCnt)
			fmt.Printf(">> %s", c[1])
			time.Sleep(500 * time.Millisecond)
			visa.ViRead(instr, resp, uint32(len(resp)), &retCnt)
			fmt.Printf("<< %s\n", trimStr(resp))
		default:
			fmt.Printf("unknow command: %s\n", line)
		}
	}
}
	
func main() {
	flag.Parse()

	gpibCmd := getCmdFromFile(*file_path)
	
	addr := fmt.Sprintf("GPIB%d::%d::INSTR", *board_id, *board_port)

	executeCmd(addr, gpibCmd)
}
