var polylineUtils = require('time-aware-polyline');
var dateFormat = require('dateformat');
var a = new polylineUtils.TimeAwareEncoder()
var points = [
  
]



// console.log(a.decodeTimeAwarePolyline(encode))

var MongoClient = require('mongodb').MongoClient;
var url = "mongodb://localhost:27017/";

MongoClient.connect(url, function(err, db) {
  if (err) throw err;
  console.log("Database created!");

  var dbo = db.db("training2");
  var query = { vhid: "351608085548159",actvt:"loc", sertm:{
      "$gte":new Date("2017-10-04T18:30:00+530")
    }};
  dbo.collection("vhtrps").find(query).toArray(function(err, result) {
    if (err) throw err;
    console.log(result);
    for (let i = 0; i < 2; i++) {
        const element = result[i];
        console.log(dateFormat(element.sertm,"yyyy-mm-dd'T'hh:MM:ss+05:30"))
        points.push([ element.loc[1],
            element.loc[0],
            dateFormat(element.sertm,"yyyy-mm-dd'T'hh:MM:ss+05:30")
            ])
        }
   

    db.close();

    var encode1 = a.encodeTimeAwarePolyline(points)
    console.log(encode1)

    console.log(a.decodeTimeAwarePolyline(encode1))
  });


});

