// +build windows

package main

import (
	"flag"
	"fmt"
	"govisa/visa"
)

var port = flag.Int("port", 1, "GPIB port number")
var cmd = flag.String("cmd", "*IDN?", "VISA command")

func main() {

	flag.Parse()

	var (
		defaultRM uint32
		instr     uint32
	)

	status := visa.ViOpenDefaultRM(&defaultRM)

	addr := fmt.Sprintf("GPIB1::%d::INSTR", *port)
	fmt.Printf("%s\n", addr)
	visa.ViOpen(defaultRM, addr, uint32(0), uint32(0), &instr)
	defer visa.ViClose(instr)

	visa.ViSetAttribute(instr, visa.VI_ATTR_TMO_VALUE, 5000)

	var (
		instrHandle uint32
		numInstrs   uint32
	)

	buf := make([]byte, 1024)
	visa.ViFindRsrc(defaultRM, "GPIB1::?*", &instrHandle, &numInstrs, buf)
	fmt.Printf("%v, %v, %s\n", instrHandle, numInstrs, visa.TrimByteSlice(buf))

	cmd_s := fmt.Sprintf("%s\n", *cmd)
	fmt.Printf("cmd: %s, len: %d\n", cmd_s, len(cmd_s))
	var retCnt uint32
	if status := visa.ViWrite(instr, cmd_s, uint32(len(cmd_s)), &retCnt); int(status) < visa.VI_SUCCESS {
		panic("Failed to write data to instrument.")
	}

	if status := visa.ViRead(instr, buf, uint32(len(buf)), &retCnt); int(status) < visa.VI_SUCCESS {
		panic("Failed to read data from instrument.")
	}
	fmt.Printf("%s\n", visa.TrimByteSlice(buf))

}
