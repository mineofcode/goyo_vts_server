package models

import (
	"fmt"
	"strings"

	"github.com/go-playground/log"
	"github.com/tidwall/tile38/client"
	"gopkg.in/mgo.v2/bson"
	opts "goyo.in/gpstracker/const"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
	patterns "goyo.in/gpstracker/patterns"
)

const (
	ksep = ":"
)

func CreateGeoFence(params datamodel.GeoFenceModelAr) (result []datamodel.GeoFenceResponse, err error) {

	_sn := getDBSession().Copy()
	defer _sn.Close()

	tile38con, err := db.Pool.Get()
	if err != nil {
		log.WithFields(log.F("func", "Factory.AnalysisLoop.til38Pool.Get")).Fatal(err)
	}
	defer tile38con.Close()
	//fmt.Println(vhid, d)
	c := col(_sn, patterns.ColGeofence)

	response := []datamodel.GeoFenceResponse{}
	for _, param := range params.Data {
		resp := datamodel.GeoFenceResponse{Imei: param.Imei, FenceName: param.FenceName}
		info, dberr := c.Upsert(bson.M{"imei": param.Imei, "fncnm": param.FenceName}, bson.M{"$set": param})
		if dberr != nil {
			resp.Status = false
			resp.Msg = "Error while creating geofence"
			//	return resp, dberr
		}

		fmt.Println(info)
		if info.Updated > 0 {
			resp.Status = true
			resp.Msg = "Geofence Updated Successfully"
			//	return resp, nil
			setWebhook(param.FenceName, param.Imei, param.Points[0], param.Points[1], param.Radius, param.Active, tile38con)
		} else if info.Matched > 0 {
			resp.Status = false
			resp.Msg = "No Changes"
			//	return resp, nil
		} else if info.UpsertedId != nil {
			resp.Status = true
			resp.Msg = "Geofence Created Successfully"
			setWebhook(param.FenceName, param.Imei, param.Points[0], param.Points[1], param.Radius, true, tile38con)
			//return resp, nil
		}

		response = append(response, resp)
	}

	return response, nil

}

func setWebhook(id string, imei string, lat float64, lon float64, radious int, isins bool, con *client.Conn) string {

	whokid := strings.Join([]string{imei, id}, ksep)
	//whokid = "ttst"
	hookCmd := ""

	if isins == false {
		hookCmd = fmt.Sprintf("DELHOOK %s",
			whokid)
	} else {

		hookCmd = fmt.Sprintf("SETHOOK %s grpc://%s NEARBY %s FENCE DETECT enter,exit POINT %f %f %v",
			whokid,
			opts.DefaultOpts.WebScok.Addr,
			imei,
			lat, lon, radious)
	}
	reply, err := con.Do(hookCmd)

	if err != nil {
		fmt.Println("Could not SET:" + err.Error())
		return "Could not SET:" + err.Error()
	} else {
		fmt.Println(reply)
	}

	return "Successfully SET"

}

//with redis

func CreateGeoFence_RED(params datamodel.GeoFenceModel) (result interface{}, err error) {

	// if params.FenceType == "circle" {
	// 	con := redigogeofence.Pool.Get()
	// 	defer con.Close()
	// 	s := reflect.ValueOf(params.Points)

	// 	_, err = con.Do("SET", params.Imei, params.FenceName, "POINT", s.Index(0).Interface(), s.Index(1).Interface())

	// 	if err != nil {
	// 		return "Could not SET:" + err.Error(), nil
	// 	}
	// 	var buffer bytes.Buffer
	// 	buffer.WriteString(params.FenceType)                  //
	// 	buffer.WriteString(",")                               //
	// 	buffer.WriteString(params.AlamTyp)                    //
	// 	buffer.WriteString(",")                               //
	// 	buffer.WriteString(strconv.FormatBool(params.Active)) //
	// 	buffer.WriteString(",")                               //
	// 	buffer.WriteString(params.CallBackURL)                //
	// 	buffer.WriteString(",")                               //
	// 	buffer.WriteString(params.Params)                     //
	// 	buffer.WriteString(",")                               //
	// 	buffer.WriteString(params.TimeType)                   //
	// 	buffer.WriteString(",")                               //
	// 	buffer.WriteString(params.FenceTime)                  //
	// 	fmt.Println(buffer.String())
	// 	_, err = con.Do("SET", params.Imei, params.FenceName+":data", "STRING", buffer.String())

	// 	if err != nil {
	// 		return "Could not SET:" + err.Error(), nil
	// 	}

	// 	return "Successfully SET", nil
	// }

	return "Invalid Command", nil
}

// //get fences
// func GetGeoFence_RED(params datamodel.GetGeoFenceParams) (result interface{}, err error) {

// 	con := redigogeofence.Pool.Get()
// 	defer con.Close()

// 	//result, err = con.Do("SCAN", params.Imei)
// 	vals, err := redis.Values(con.Do("SCAN", params.Imei))
// 	if err != nil {
// 		return fmt.Sprintf("could not SCAN: %v\n", err), err
// 	}
// 	// vals, err = redis.Values(vals[1], nil)
// 	// if err != nil {
// 	// 	return "Could not SET:" + err.Error(), nil
// 	// }

// 	return fmt.Sprintf("%s", vals[1]), nil
// }

func GetGeoFence(params interface{}) (result []datamodel.GeoFenceModel, err error) {

	_sn := getDBSession().Copy()
	defer _sn.Close()

	//fmt.Println(vhid, d)

	c := col(_sn, patterns.ColGeofence)
	var datamodel []datamodel.GeoFenceModel
	c.Find(params).All(&datamodel)
	return datamodel, nil
}

//delete fence

func DeleteGeoFence(params datamodel.GetGeoFenceParams) (result interface{}, err error) {

	_sn := getDBSession().Copy()
	defer _sn.Close()

	tile38con, err := db.Pool.Get()
	if err != nil {
		log.WithFields(log.F("func", "Factory.AnalysisLoop.til38Pool.Get")).Fatal(err)
	}
	defer tile38con.Close()

	qury := bson.M{}

	if params.FenceName != "" {
		qury = bson.M{"imei": params.Imei, "fncnm": params.FenceName}
	} else {
		qury = bson.M{"imei": params.Imei}
	}
	c := col(_sn, patterns.ColGeofence)
	info, err1 := c.RemoveAll(qury)

	if err1 != nil {
		return "Could not delete:" + err1.Error(), err
	}

	hookCmd := ""
	if params.FenceName == "" {
		hookCmd = fmt.Sprintf("PDELHOOK %s:*",
			params.Imei)
	} else {
		hookCmd = fmt.Sprintf("PDELHOOK %s:%s",
			params.Imei, params.FenceName)
	}

	reply, err := tile38con.Do(hookCmd)

	if err != nil {
		fmt.Println("Could not SET:" + err.Error())
		//return "Could not SET:" + err.Error()

	} else {
		fmt.Println(reply)

	}

	return fmt.Sprintf("%v Record(s) deleted successfully.", info.Removed), nil
}

func DeleteGeoFence_RED(params datamodel.GetGeoFenceParams) (result interface{}, err error) {

	// con := redigogeofence.Pool.Get()
	// defer con.Close()

	// _, err = con.Do("DEL", params.Imei, params.FenceName)

	// if err != nil {
	// 	return "Could not delete:" + err.Error(), err
	// }

	// _, err = con.Do("DEL", params.Imei, params.FenceName+":data")

	// if err != nil {
	// 	return "Could not delete:" + err.Error(), err
	// }

	return "Successfully deleted", nil
}
