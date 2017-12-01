package protocalHandler

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/crc16"
	"goyo.in/gpstracker/models"
)

//Clients keep ip address in map key
var Clients = make(map[string]clientsMod)
var locker sync.Mutex

type clientsMod struct {
	con  net.Conn
	imei string
	lstm time.Time
}

//Terminal connection persists
func addClient(conn net.Conn, client string) {
	locker.Lock()
	defer locker.Unlock()
	ip_address := conn.RemoteAddr().String()

	Clients[ip_address] = clientsMod{con: conn, imei: client, lstm: time.Now()}
}
func RemoveClient(conn net.Conn) {
	locker.Lock()
	defer locker.Unlock()

	ip_address := conn.RemoteAddr().String()
	delete(Clients, ip_address)
}

func getClient(conn net.Conn) (client clientsMod) {
	ip_address := conn.RemoteAddr().String()
	_clientsMod := Clients[ip_address]
	return _clientsMod
}

//ParseData Parse recceived data
func ParseData(_data []byte, lendata int, connection net.Conn) {

	if !((_data[0] == 0x78) && (_data[1] == 0x78)) {
		fmt.Println("not equl")
	} else if _data[3] == 0x01 {
		registerDevice(_data, lendata, connection)
	} else if _data[3] == 0x13 {
		heartBeat(_data, lendata, connection)
	} else if (_data[3] == 0x22) || (_data[3] == 0x12) {
		locationDt(_data, lendata, connection)
	}

	//data1 := string(_data[:lendata])

	//fmt.Print(data1)
	//connection.Write([]byte(data1 + "\n"))
}

//register device in terminal
func registerDevice(_data []byte, lendata int, connection net.Conn) {
	reply := []byte{0x78, 0x78, 0x05, 0x01}          //assign reply variable
	serial := _data[12:14]                           //get crc from data
	_crxCRC := append([]byte{0x05, 0x01}, serial...) // create crc string
	_crxCRCF := crc16.GetCrc16(_crxCRC)              // get computed crc in variable
	_crxCRCF = append(serial, _crxCRCF...)           // append final crc and reply data
	reply = append(reply, _crxCRCF...)               // append final crc and reply data
	reply = append(reply, 0x0D, 0x0A)                //EOF
	//get imei number
	_imei := fmt.Sprintf("%x", _data[4:12])[1:16] //getting imei number
	// fmt.Println(_imei)
	// fmt.Println(reply)
	connection.Write(reply)

	addClient(connection, _imei)
	//need to call mongo db
	// fmt.Println(len(Clients))
	// fmt.Println(_imei)
	//send reply to terminal
	// socket := socketios.GetSocketIO()
	// socket.BroadcastTo(_imei, "msgd", fmt.Sprintf("%x", reply))

	_d := bson.M{
		"acttm": time.Now(),
		"actvt": "login",
		"sertm": time.Now(),
		"imei":  _imei,
		"btrst": "BTRY",
		"flag":  "inprog",
		"appvr": "1.0",
		"vhid":  _imei,
		"speed": 0}
	models.UpdateData(_d, _imei, "reg")

}

//getting heart beat
func heartBeat(_data []byte, lendata int, connection net.Conn) {
	_clnt := getClient(connection)
	if _clnt.imei == "" {
		return
	}
	_clnt.lstm = time.Now()

	reply := []byte{0x78, 0x78, 0x05, 0x13} //assign reply variable
	serial := _data[7:9]                    //get crc from data

	_crxCRC := append([]byte{0x05, 0x13}, serial...) // create crc string
	_crxCRCF := crc16.GetCrc16(_crxCRC)              // get computed crc in variable
	_crxCRCF = append(serial, _crxCRCF...)           // append final crc and reply data
	reply = append(reply, _crxCRCF...)               // append final crc and reply data
	reply = append(reply, 0x0D, 0x0A)                //EOF
	//Client get by ipaddress

	//extract data from received data
	_prd := fmt.Sprintf("%08b", _data[4:5])
	_prd = _prd[1 : len(_prd)-1]
	btrt := "BTRY"
	//fmt.Println(_prd)
	data := HertBt{
		Acttm:  time.Now(),
		Actvt:  "hrtbt",
		Sertm:  time.Now(),
		Speed:  0,
		Imei:   _clnt.imei,
		Flag:   "inprog",
		Appvr:  "1.0",
		Vhid:   _clnt.imei,
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

	if data.Chrg == 1 {
		data.Btrst = "CHRG"
	}

	// data := bson.M{
	// 	"acttm":  time.Now(),
	// 	"actvt":  "hrtbt",
	// 	"sertm":  time.Now(),
	// 	"imei":   _clnt.imei,
	// 	"flag":   "inprog",
	// 	"appvr":  "1.0",
	// 	"vhid":   _clnt.imei,
	// 	"btr":    batryper(int(_data[5])),
	// 	"btrst":  btrt,
	// 	"oe":     _prd[0:1],                           //1: oil and electricity disconnected, 0: gas oil and electricity
	// 	"gp":     _prd[1:2],                           //1: GPS tracking is on,0: GPS tracking is off
	// 	"alm":    (_prd[2:3] + _prd[3:4] + _prd[4:5]), //100: SOS,011: Low Battery Alarm,010: Power Cut Alarm,001: Shock Alarm,000: Normal
	// 	"chrg":   _prd[5:6],                           //1: Charge On,0: Charge Off
	// 	"acc":    _prd[6:7],                           //1: ACC high,0: ACC Low
	// 	"df":     _prd[7:8],                           //1: Defense Activated,0: Defense Deactivated,
	// 	"gsmsig": networkper(int(_data[6])),           //0x00: no signal,0x01: extremely weak signal,0x02: very weak signal,0x03: good signal,0x04: strong signal
	// 	"lng":    1}

	//need to call mongo db
	models.UpdateData(data, _clnt.imei, "hrt")
	// fmt.Println(fmt.Sprintf("%x", reply))
	//a := fmt.Sprintf("%v", data)
	//send reply to terminal
	// socket := socketios.GetSocketIO()
	// socket.BroadcastTo(_clnt.imei, "msgd", a)
	connection.Write(reply)
}

//getting heart beat
func locationDt(_data []byte, lendata int, connection net.Conn) {

	//reply := []byte{0x78, 0x78, 0x05, 0x13} //assign reply variable
	//serial := _data[7:9] //get crc from data
	//78 78 1F 12 0B 08 1D 11 2E 10 CC 02 7A C7 EB 0C 46 58 49 00 14 8F 01 CC 00 28 7D 00 1F B8 00 03 80 81 0D 0A
	//_crxCRC := append([]byte{0x05, 0x13}, serial...) // create crc string
	//_crxCRCF := crc16.GetCrc16(_crxCRC)              // get computed crc in variable
	//_crxCRCF = append(_crxCRCF, 0x0D, 0x0A)          // append final crc and reply data
	//reply = append(reply, _crxCRCF...)               //append all data to reply variable
	//Client get by ipaddress
	//datetime

	_clnt := getClient(connection)
	if _clnt.imei == "" {
		return
	}
	_clnt.lstm = time.Now()

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

	data := bson.M{
		"gpstm":   _dt,
		"actvt":   "loc",
		"sertm":   time.Now(),
		"imei":    _clnt.imei,
		"flag":    "inprog",
		"appvr":   "1.0",
		"sat":     _stlt,
		"loc":     []float64{toFixed(_lon, 6), toFixed(_lat, 6)},
		"postyp":  _courus[2:3],
		"bearing": _bearing,
		"speed":   _data[19],
		"vhid":    _clnt.imei}

	//need to call mongo db
	models.UpdateData(data, _clnt.imei, "loc")
	//fmt.Println(fmt.Sprintf("%x", reply))
	//send reply to terminal
	//connection.Write(reply)
}
