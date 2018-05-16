package redigogeofence

import (
	"fmt"

	"github.com/go-playground/log"
	"goyo.in/gpstracker/db"
)

var errorc error

type Sel_params struct {
	imei   string    `json:"imei"`
	fncnm  string    `json:"fncnm"`
	fnctyp string    `json:"fnctyp"`
	almtyp string    `json:"almtyp"`
	points []float64 `json:"points"`
	url    string    `json:"url"`
	param  string    `json:"param"`
}

func Start() {

}

func CheckValue(point []float64, speed byte, imei string) {
	// var format string = "2006-01-02 15:04:05 -0700"
	// con := Pool.Get()
	// defer con.Close()

	// var t_arr []Sel_params

	// fmt.Println(point)
	// t1 := time.Now()

	// vals, err := redis.Values(con.Do("NEARBY", imei, "LIMIT", "10", "POINT", point[1], point[0], 1000))
	// if err != nil {
	// 	log.Fatalf("could not NEARBY: %v\n", err)
	// }
	// // the first element is the cursor
	// if len(vals) < 2 {
	// 	log.Fatal("invalid value")
	// }
	// vals, err = redis.Values(vals[1], nil)
	// if err != nil {
	// 	log.Fatal("invalid value")
	// }
	// for _, val := range vals {
	// 	strs, err := redis.Strings(val, nil)
	// 	if err != nil || len(strs) < 2 {
	// 		log.Fatal("invalid value: %v", err)
	// 	}
	// 	//fmt.Printf("%s >> %s\n", strs[0], strs[1])
	// 	ret, _ := redis.String(con.Do("GET", imei, strs[0]+":data"))
	// 	//a := strings.Split(s.String(), ",")
	// 	var data1 []string = strings.Split(ret, ",")

	// 	if len(data1) == 7 {
	// 		data1 = append(data1, time.Now().Format(format))
	// 		_, err = con.Do("SET", imei, strs[0]+":data", "STRING", strings.Join(data1, ","))
	// 		if err != nil {

	// 		}
	// 		dt := Sel_params{
	// 			imei:   imei,
	// 			almtyp: data1[1],
	// 			fncnm:  strs[0],
	// 			fnctyp: data1[0],
	// 			url:    data1[3],
	// 			param:  data1[4],
	// 		}
	// 		t_arr = append(t_arr, dt)

	// 	} else {
	// 		// data1[7] = time.Now().Format(format)
	// 		// _, err = con.Do("SET", imei, strs[0]+":data", "STRING", strings.Join(data1, ","))
	// 		// if err != nil {

	// 		// }
	// 		frmt, err := time.Parse(format, data1[7])
	// 		if err != nil {
	// 			fmt.Println(data1[7], err)
	// 		}

	// 		diff := t1.Sub(frmt)
	// 		fmt.Println("already checkd wait for ", diff, " min", diff.Minutes())
	// 		if diff.Minutes() > 15 {
	// 			fmt.Println("rechecked ")
	// 			data1[7] = time.Now().Format(format)
	// 			_, err = con.Do("SET", imei, strs[0]+":data", "STRING", strings.Join(data1, ","))
	// 			if err != nil {

	// 			}
	// 			dt := Sel_params{
	// 				imei:   imei,
	// 				almtyp: data1[1],
	// 				fncnm:  strs[0],
	// 				fnctyp: data1[0],
	// 				url:    data1[3],
	// 				param:  data1[4],
	// 			}
	// 			t_arr = append(t_arr, dt)

	// 		} else {
	// 			fmt.Println("already checkd wait for ", diff, " min")
	// 		}
	// 	}

	// 	fmt.Println(t_arr)

	// }
	// go CallService(t_arr)
}

func SetValue(point []float64, speed byte, imei string) {
	tile38con, err := db.Pool.Get()
	if err != nil {
		log.WithFields(log.F("func", "Factory.AnalysisLoop.til38Pool.Get")).Fatal(err)
	}
	defer tile38con.Close()
	hookCmd := fmt.Sprintf("SET %s vh FIELD speed %v POINT %f %f",
		imei,
		speed,
		point[1],
		point[0])
	//fmt.Println(hookCmd)
	_, err1 := tile38con.Do(hookCmd)

	if err1 != nil {
		fmt.Println("Could not SET:" + err1.Error())

	} else {
		//fmt.Println(reply)

	}

	//go CallService(t_arr)
}
