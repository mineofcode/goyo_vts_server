package protocalHandler

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"

	"goyo.in/gpstracker/crc16"
	"goyo.in/gpstracker/datamodel"
)

var serialmax = 65535
var currentSerial = 0

func SendCommand(client net.Conn, msg datamodel.DeviceCommands) {
	commandWrapper(client, msg)
}

func commandWrapper(client net.Conn, msg datamodel.DeviceCommands) {
	//78 78 0E 80 08 00 00 00 00 73 6F 73 23 00 01 6D 6A 0D 0A

	//convert ascii command into byte array
	strcommand := []byte(msg.Cmd)
	//convert into hex value
	commandlen := len(strcommand)
	fmt.Println(commandlen)

	cmd := make([]byte, 5)
	cmd[0] = 0x78                                 //start command
	cmd[1] = 0x78                                 //start command
	cmd[2] = byte(1 + 1 + 4 + commandlen + 2 + 2) //lenth of packaget
	cmd[3] = 0x80                                 //protocal command
	cmd[4] = byte(hex.EncodedLen(commandlen))     //length of content
	//-------------------------------------------------------
	serverbit := getServerBit(msg.Uniqid)
	cmd = append(cmd, serverbit...) /// Server |
	//  Flag  |
	//  Bit   |
	//....... |
	cmd = append(cmd, strcommand...) ///////Content of command
	//-------------------------------------------------------
	//cmd = append(cmd, []byte{0x00, 0x02}...) ///////Language
	//-------------------------------------------------------
	//add serial number
	serial := getCurrentSerial()

	//	fmt.Println(serial)
	cmd = append(cmd, serial...) ///////Content of command
	//end serial number
	//------------------------------------------------------
	// Compute CRC
	// cmd = append(cmd, getCurrentSerial()...) //8,9
	// cmd = append(cmd, 0x00) //8,9
	// cmd = append(cmd, 0x01) //8,9
	//_crxCRC := []byte{0x80, 0x0E, 0x00, 0x01} // create crc string

	_crxCRCF := crc16.GetCrc16(cmd[2:])
	// get computed crc in variable
	fmt.Printf("%02X\n", _crxCRCF)

	cmd = append(cmd, _crxCRCF...)
	// cmd = append(cmd, 0x6D) // FIXED end
	// cmd = append(cmd, 0x6A) // FIXED end

	//------------------------------------------------------------------------------------------
	cmd = append(cmd, 0x0D) // FIXED end
	cmd = append(cmd, 0x0A) // FIXED end
	fmt.Printf("%02X\n", cmd)
	client.Write(cmd)

}

func getCurrentSerial() []byte {

	if currentSerial == serialmax {
		currentSerial = 0
	}
	currentSerial += 1
	//bs := make([]byte, 2)
	src := []byte(fmt.Sprintf("%04x", currentSerial))

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	// binary. (bs, uint32(currentSerial))
	//fmt.Println([]byte(fmt.Sprintf("%2d", currentSerial)))
	//	fmt.Println(dst)
	return dst[:n]
}

func getServerBit(serverbit uint) []byte {

	//bs := make([]byte, 2)
	src := []byte(fmt.Sprintf("%08x", serverbit))

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(dst)
	return dst[:n]
}
