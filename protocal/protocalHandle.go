package protocalHandler

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"goyo.in/gpstracker/datashare"

	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/crc16"
	"goyo.in/gpstracker/fcm"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/redigogeofence"
	"goyo.in/gpstracker/shared"
	//"github.com/jasonlvhit/gocron"
)

//Start the Function

//Terminal connection persists
func addClient(conn net.Conn, client string, allwspd int, vtsid int,
	PushClients []string, LLoc []float64, LAC int, LACC int, LTime time.Time) shared.ClientsMod {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()
	ipaddress := conn.RemoteAddr().String()
	fmt.Println("Client Added  ", ipaddress)
	shared.Clients[ipaddress] = shared.ClientsMod{Con: conn, Imei: client, Lstm: LTime,
		Allwspd: allwspd, VtsID: vtsid, PClients: PushClients, Loc: LLoc, AC: LAC, Acc: LACC}
	return shared.Clients[ipaddress]
}

func RemoveClient(conn net.Conn) {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()
	ipaddress := conn.RemoteAddr().String()
	delete(shared.Clients, ipaddress)
}

func getClient(conn net.Conn) (client shared.ClientsMod) {
	ipaddress := conn.RemoteAddr().String()
	return shared.Clients[ipaddress]
}

func setClient(conn net.Conn, client shared.ClientsMod) {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()

	ipaddress := conn.RemoteAddr().String()
	shared.Clients[ipaddress] = client
}

func UpdateAllowSpeed(speed int, ip string) {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()
	_clientsMod := shared.Clients[ip]
	if _clientsMod.Imei != "" {
		_clientsMod.Allwspd = speed
		shared.Clients[ip] = _clientsMod
	}
}

func UpdatePushClient(party []string, ip string) {
	shared.Locker.Lock()
	defer shared.Locker.Unlock()
	_clientsMod := shared.Clients[ip]
	if _clientsMod.Imei != "" {
		_clientsMod.PClients = party
		shared.Clients[ip] = _clientsMod
	}

}

//ParseData Parse recceived data
func ParseData(_data []byte, lendata int, connection net.Conn) {
	//fmt.Printf("%02x\n", bytes.Trim(_data, "\x00"))
	//Concox Device
	//Check for valid command
	if !(_data[0] == 0x78 || _data[0] == 0x79) {

		// if _data[0] == 0x5b && _data[len(_data)-1] == 0x5d {
		// 	fmt.Printf("%02x\n", bytes.Trim(_data, "\x00"), " watch detected")
		// } else {
		fmt.Println("invalid command", _data[len(_data)-1])
		// }

	} else if _data[0] == 0x78 { //check for 78 commands
		if _data[3] == 0x01 {
			registerDevice(_data, lendata, connection) // Registration Response
		} else if _data[3] == 0x13 {
			heartBeat(_data, lendata, connection) // HeartBeat Response
		} else if (_data[3] == 0x22) || (_data[3] == 0x12) {
			locationDt(_data, lendata, connection) //location data
		} else if _data[3] == 0x15 {
			commandReply(_data, lendata, connection) //location data
		}
	} else if _data[0] == 0x79 { //check for 79 Alarm Commnd

		switch _data[5] {
		case 0x00: //External power voltage
			{
				//needs to be added
				break
			}
		case 0x04: //Terminal status sync
			{
				almDecode04(_data, lendata, connection)
				//needs to be added
				break
			}
		case 0x5: //door status
			{
				almDecode05(_data, lendata, connection)
				break
			}

		}
	}

	//data1 := string(_data[:lendata])

	//fmt.Print(data1)
	//connection.Write([]byte(data1 + "\n"))
}

var cmd = ""

//register device in terminal
func registerDevice(_data []byte, lendata int, connection net.Conn) {
	reply := []byte{0x78, 0x78, 0x05, 0x01}          //assign reply variable
	serial := _data[12:14]                           //get crc from data
	_crxCRC := append([]byte{0x05, 0x01}, serial...) // create crc string
	_crxCRCF := crc16.GetCrc16(_crxCRC)              // get computed crc in variable
	//fmt.Printf("%02X\n", _crxCRCF)
	_crxCRCF = append(serial, _crxCRCF...) // append final crc and reply data
	reply = append(reply, _crxCRCF...)     // append final crc and reply data
	reply = append(reply, 0x0D, 0x0A)      //EOF
	//get imei number
	_imei := fmt.Sprintf("%x", _data[4:12])[1:16] //getting imei number
	// fmt.Println(_imei)
	// fmt.Println(reply)

	spd := models.GetVehiclesData(_imei)
	if spd.VhId == "" {
		fmt.Println(_imei + " is not registered")
		return
	} else {
		fmt.Println(_imei + " is logged In")
	}

	connection.Write(reply)

	addClient(connection, _imei, spd.AllowSpd, spd.VtsID, spd.PClients, spd.Loc, spd.AC, spd.ACC, spd.SerTm)

	_d := bson.M{
		"acttm": time.Now(),
		"actvt": "login",
		"sertm": time.Now(),
		"imei":  _imei,
		"btrst": "BTRY",
		"flag":  "inprog",
		"appvr": "1.0",
		"vhid":  _imei,
		"ip":    connection.RemoteAddr().String(),
		"speed": 0}
	models.UpdateData(_d, _imei, "reg", nil)
}

//getting heart beat
func heartBeat(_data []byte, lendata int, connection net.Conn) {
	_clnt := getClient(connection)
	if _clnt.Imei == "" {
		return
	}

	reply := []byte{0x78, 0x78, 0x05, 0x13} //assign reply variable
	serial := _data[7:9]                    //get crc from data

	_crxCRC := append([]byte{0x05, 0x13}, serial...) // create crc string
	_crxCRCF := crc16.GetCrc16(_crxCRC)              // get computed crc in variable
	_crxCRCF = append(serial, _crxCRCF...)           // append final crc and reply data
	reply = append(reply, _crxCRCF...)               // append final crc and reply data
	reply = append(reply, 0x0D, 0x0A)                //EOF
	//Client get by ipaddress

	//extract data from received data
	fmt.Println(_data[4:5])
	_prd := fmt.Sprintf("%08b", _data[4:5])
	_prd = _prd[1 : len(_prd)-1]
	btrt := "BTRY"
	//fmt.Println(_prd)
	data := HertBt{
		Acttm:  time.Now(),
		Actvt:  "hrtbt",
		Sertm:  time.Now(),
		Speed:  0,
		Imei:   _clnt.Imei,
		Flag:   "inprog",
		Appvr:  "1.0",
		Vhid:   _clnt.Imei,
		Btr:    batryper(int(_data[5])),
		Btrst:  btrt,
		Alm:    (_prd[2:3] + _prd[3:4] + _prd[4:5]), //100: SOS,011: Low Battery Alarm,010: Power Cut Alarm,001: Shock Alarm,000: Normal
		Gsmsig: networkper(int(_data[6])),           //0x00: no signal,0x01: extremely weak signal,0x02: very weak signal,0x03: good signal,0x04: strong signal
	}

	data.Oe, _ = strconv.Atoi(_prd[0:1])   //1: oil and electricity disconnected, 0: gas oil and electricity
	data.Gp, _ = strconv.Atoi(_prd[1:2])   //1: GPS tracking is on,0: GPS tracking is off
	data.Chrg, _ = strconv.Atoi(_prd[5:6]) //1: Charge On,0: Charge Off
	data.Acc, _ = strconv.Atoi(_prd[6:7])  //1: ACC high,0: ACC Low
	data.Df, _ = strconv.Atoi(_prd[7:8])   //1: Defense Activated,0: Defense Deactivated
	//
	// fmt.Println(data.Acc)

	_clnt.Lstm = data.Sertm
	var otherdata interface{}
	_clnt.Acc = data.Acc
	otherdata = bson.M{
		"actvt":    "loc",
		"sertm":    time.Now(),
		"imei":     _clnt.Imei,
		"alwspeed": _clnt.Allwspd,
		"isp":      false,
		"flag":     "acc",
		"acc":      data.Acc,
		"appvr":    "1.0",
		"loc":      _clnt.Loc,
		"bearing":  0,
		"speed":    0,
		"vhid":     _clnt.Imei,
	}
	if _clnt.Acc != data.Acc {
		go fcm.SendACCAlertTotopic(_clnt.Imei, data.Acc)
	}

	if data.Chrg == 1 {
		data.Btrst = "CHRG"
	}

	setClient(connection, _clnt)
	//need to call mongo db
	models.UpdateData(data, _clnt.Imei, "hrt", otherdata)
	// fmt.Println(fmt.Sprintf("%x", reply))
	//a := fmt.Sprintf("%v", data)
	//send reply to terminal
	// socket := socketios.GetSocketIO()
	// socket.BroadcastTo(_clnt.imei, "msgd", a)
	connection.Write(reply)

	//send tp 3rd party client
	datashare.SendToClient(_clnt)

}

//getting heart beat
func locationDt(_data []byte, lendata int, connection net.Conn) {

	_clnt := getClient(connection)
	if _clnt.Imei == "" {
		return
	}
	_dt := "20" + fmt.Sprintf("%d-%d-%d %d:%d:%d", _data[4], _data[5], _data[6], _data[7], _data[8], _data[9]) //conver to Date
	//fmt.Println(_dt)
	crs := fmt.Sprintf("%x", _data[10:11])            //Quantity of GPS	information	satellites
	_stlt, _ := strconv.ParseInt("0x0"+crs[1:], 0, 8) //satlites
	//extract data from received data
	var _lat float64
	var _lon float64

	_lat = float64(binary.BigEndian.Uint32(_data[11:15])) / (30000 * 60) //Lattitude
	_lon = float64(binary.BigEndian.Uint32(_data[15:19])) / (30000 * 60) //Longitude

	_courus := fmt.Sprintf("%016b", binary.BigEndian.Uint16(_data[20:22]))
	_bearing, _ := strconv.ParseInt(_courus[6:], 2, 64) // get bearing
	point := []float64{toFixed(_lon, 6), toFixed(_lat, 6)}
	//fmt.Println(_clnt.allwspd)

	data := bson.M{
		"gpstm":    _dt,
		"actvt":    "loc",
		"sertm":    time.Now(),
		"imei":     _clnt.Imei,
		"alwspeed": _clnt.Allwspd,
		"isp":      false,
		"flag":     "inprog",
		"appvr":    "1.0",
		"acc":      _clnt.Acc,
		"sat":      _stlt,
		"loc":      point,
		"postyp":   _courus[2:3],
		"bearing":  _bearing,
		"speed":    _data[19],
		"vhid":     _clnt.Imei}

	_clnt.Lstm = time.Now()
	_clnt.Loc = point

	if _clnt.Allwspd > 0 {

		crspeed := int(_data[19])

		//fmt.Println(crspeed)
		if crspeed > _clnt.Allwspd {
			// speed voilence
			//fmt.Println(int(_data[19]))
			go fcm.SendSpeedAlertTotopic(_clnt.Imei, crspeed)
			data["lstspd"] = crspeed
			data["lstspdtm"] = time.Now()
			data["isp"] = true
		}
	}

	setClient(connection, _clnt)
	//need to call mongo db
	models.UpdateData(data, _clnt.Imei, "loc", nil)
	checkGeofence(point, _data[19], _clnt.Imei)

	//send tp 3rd party client
	datashare.SendToClient(_clnt)
}

// send points to check for geofence
func checkGeofence(pont []float64, speed byte, imei string) {
	go redigogeofence.SetValue(pont, speed, imei)
}

///alarm decoding
func almDecode00() {

}

//decode alarm details
func almDecode04(_data []byte, lendata int, connection net.Conn) {
	_clnt := getClient(connection)
	if _clnt.Imei == "" {
		return
	}

	_dataT := bytes.Trim(_data, "\x00")
	// 797900849404414c4d313d37353b414c4d323d44353b414c4d333d35463b535441313d34303b4459443d30313b534f533d393030343339303837342c2c3b43454e5445523d393030343339303837343b46454e43453d46656e63652c4f46462c302c302e3030303030302c302e3030303030302c3330302c494e206f72204f55542c313b00ef5cfa0d0a
	//ALM1
	// s := []byte{0x01}
	// src := _data[7 : len(_data)-6]
	srt := string(_dataT[6 : len(_dataT)-6])
	arrcmt := strings.Split(srt, ";")

	// alm1 := arrcmt[0] // ALM1;
	// alm2 := arrcmt[1] // ALM2;
	// alm3 := arrcmt[2] // ALM3;

	//
	DYDVAL := strings.SplitN(arrcmt[4], "=", -1)
	src := []byte(DYDVAL[1])
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	DYD := fmt.Sprintf("%08b", dst[:n])
	oe := DYD[7:8]

	if DYD[5:6] == "1" || DYD[6:7] == "1" {
		return
	}

	eventData := bson.M{
		"sertm": time.Now(),
		"vhid":  _clnt.Imei,
		"oe":    oe,
		"evt":   "dd",
		"actvt": "evt",
	}

	data := bson.M{
		"sertm": time.Now(),
		"vhid":  _clnt.Imei,
		"oe":    oe,
	}

	models.UpdateData(data, _clnt.Imei, "dd", eventData)

}

//door status pin info
func almDecode05(_data []byte, lendata int, connection net.Conn) {
	_clnt := getClient(connection)
	if _clnt.Imei == "" {
		return
	}

	_prd := fmt.Sprintf("%08b", _data[6:7])
	_prd = _prd[1 : len(_prd)-1]
	// println(_prd[7:8])
	val, _ := strconv.Atoi(_prd[7:8])
	otherdata := bson.M{
		"evt":   "d1",
		"sertm": time.Now(),
		"vhid":  _clnt.Imei,
		"val":   val,
		"actvt": "evt",
	}
	_clnt.AC = val
	_clnt.Lstm = time.Now()

	data := bson.M{
		"sertm": time.Now(),
		"imei":  _clnt.Imei,
		"vhid":  _clnt.Imei,
		"d1":    val,
	}
	setClient(connection, _clnt)
	models.UpdateData(data, _clnt.Imei, "d1", otherdata)

}

//command reply 02 june 2018

func commandReply(_data []byte, lendata int, connection net.Conn) {
	//78 78 2d 15 25 00 00 03 f0 437574206f666620746865206675656c20737570706c793a2053756363657373210002000e149d0d0a
	_dataT := bytes.Trim(_data, "\x00")
	data := binary.BigEndian.Uint32(_data[5:9])
	stringtext := string(_dataT[9 : len(_dataT)-8])
	go fcm.SendEventAlertTotopic(data, stringtext)
}
