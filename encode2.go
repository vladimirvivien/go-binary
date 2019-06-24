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
	Sensid uint32
	Locid  uint16
	Tstamp uint32
	Temp   int16
}

func main() {
	dataOut := []packet{
		{Sensid: 1, Locid: 1233, Tstamp: 123452123, Temp: 12},
		{Sensid: 2, Locid: 4567, Tstamp: 133452124, Temp: 32},
		{Sensid: 7, Locid: 8910, Tstamp: 143452125, Temp: -12},
	}

	// encode a slice of packets
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, dataOut); err != nil {
		fmt.Println(err)
		return
	}

	// decode all items in slice
	dataIn := make([]packet, 3)
	//bufIn := bytes.NewReader(bufOut.Bytes())
	if err := binary.Read(buf, binary.LittleEndian, dataIn); err != nil {
		fmt.Println("failed to Read:", err)
		return
	}

	fmt.Printf("%v\n", dataIn)

}
