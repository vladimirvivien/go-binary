package main

import (
	"encoding/binary"
	"fmt"
	"time"
)

/*
0       1       2       3       4       5       6       7
0123456701234567012345670123456701234567012345670123456701234567
+-------+-------+-------+-------+-------+-------+-------+------+
|    SensorID   |   LocationID  |            Timestamp         |
+-------+-------+-------+-------+-------+-------+-------+------+
|      Temp     |
+---------------+
*/

func main() {
	buf := make([]byte, 10)
	ts := uint32(time.Now().Unix())

	// encoding the data into buf
	binary.BigEndian.PutUint16(buf[0:], 0xa20c) // sensorID
	binary.BigEndian.PutUint16(buf[2:], 0x04af) // locationID
	binary.BigEndian.PutUint32(buf[4:], ts)     // timestamp
	binary.BigEndian.PutUint16(buf[8:], 479)    // temp

	fmt.Printf("% x\n", buf)

	// decoding the data from buf
	sensorID := binary.BigEndian.Uint16(buf[0:])
	locID := binary.BigEndian.Uint16(buf[2:])
	tstamp := binary.BigEndian.Uint32(buf[4:])
	temp := binary.BigEndian.Uint16(buf[8:])

	fmt.Printf("sid: %0#x, locID %0#x ts: %0#x, temp:%d\n", sensorID, locID, tstamp, temp)

}
