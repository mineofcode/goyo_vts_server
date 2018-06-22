var polylineUtils = require('time-aware-polyline');
var a = new polylineUtils.TimeAwareEncoder()
var points = [
    [19.13626, 72.92506, '2016-07-21T05:43:09+00:00'],
    [19.13597, 72.92495, '2016-07-21T05:43:15+00:00'],
    [19.13553, 72.92469, '2016-07-21T05:43:21+00:00']
]
var encode = a.encodeTimeAwarePolyline(points)

console.log(encode)
console.log(a.decodeTimeAwarePolyline(encode))

var MongoClient = require('mongodb').MongoClient;
var url = "mongodb://localhost:27017/";

MongoClient.connect(url, function(err, db) {
  if (err) throw err;
  console.log("Database created!");

  var dbo = db.db("training2");
  var query = { vhid: "351608085548159",actvt:"loc", sertm:{"$lte":"2017-11-06T00:00:00.000Z"} };
  dbo.collection("vhtrps").find(query).toArray(function(err, result) {
    if (err) throw err;
    console.log(result);
    db.close();
  });


});

