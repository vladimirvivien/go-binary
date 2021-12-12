package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
packet representes numeric data to be encoded as shown below:

0       1       2       3       4       5       6       7
0123456701234567012345670123456701234567012345670123456701234567
+-------+-------+-------+-------+-------+-------+-------+------+
|    SensorID   |   LocationID  |            Timestamp         |
+-------+-------+-------+-------+-------+-------+-------+------+
|      Temp     |
+---------------+
*/
type packet struct {
	Sensid uint16
	Locid  uint16
	Tstamp uint32
	Temp   int16
}

func main() {
	dataIn := packet{Sensid: 1, Locid: 1233, Tstamp: 123452123, Temp: 12}

	buf := new(bytes.Buffer)

	// put single encoded data packet into writer buf
	// automatically marshaled to type packet
	if err := binary.Write(buf, binary.BigEndian, dataIn); err != nil {
		fmt.Println(err)
		return
	}

	// get single encoded data packet from reader
	// automatically unmarshal to type packet
	var dataOut packet
	if err := binary.Read(buf, binary.BigEndian, &dataOut); err != nil {
		fmt.Println("failed to Read:", err)
		return
	}

	fmt.Printf("%v", dataOut)

}
