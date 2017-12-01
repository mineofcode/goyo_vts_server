package dataprocess

import "math"
import "time"
import "errors"

const EARTH_RADIUS float64 = 6371

func getDistance(origin []float64, destination []float64) float64 {
	// return distance in meters
	lon1 := toRadian(origin[0])
	lat1 := toRadian(origin[1])
	lon2 := toRadian(destination[0])
	lat2 := toRadian(destination[1])

	var deltaLat = lat2 - lat1
	var deltaLon = lon2 - lon1

	var a = math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(deltaLon/2), 2)
	var c = 2 * math.Asin(math.Sqrt(a))
	return c * EARTH_RADIUS * 1000
}

func toRadian(degree float64) float64 {
	return degree * math.Pi / 180
}

func isDifferentSegmentDist(end []float64, start []float64) bool {
	// function to determine whether a polyline
	// segment split should happen
	var distance = getDistance(start, end)
	return distance > 500
}

func isDifferentSegmentTime(end time.Time, start time.Time) bool {
	// function to determine whether a polyline
	// segment split should happen
	return end.Sub(start).Minutes() > 10
}

func updateBearing(oldBearing int, gpxPair [][]float64) int {
	var start = []float64{gpxPair[0][0], gpxPair[0][1]}
	var end = []float64{gpxPair[1][0], gpxPair[1][1]}
	var newBearing = computeHeading(start, end)

	if newBearing != 0 {
		return round(newBearing*100) / 100.0
	} else {
		return oldBearing
	}
}

func computeHeading(start []float64, end []float64) float64 {
	var lat1 = toRadian(start[0])
	var lat2 = toRadian(end[0])
	var lng1 = toRadian(start[1])
	var lng2 = toRadian(end[1])
	return math.Atan2(math.Sin(lng2-lng1)*math.Cos(lat2), math.Cos(lat1)*math.Sin(lat2)-math.Sin(lat1)*math.Cos(lat2)*math.Cos(lng2-lng1)) * 180 / math.Pi
}

///Encoder

var (
	errDimensionalMismatch  = errors.New("dimensional mismatch")
	errInvalidByte          = errors.New("invalid byte")
	errUnterminatedSequence = errors.New("unterminated sequence")
)

func round(x float64) int {
	if x < 0 {
		return int(-math.Floor(-x + 0.5))
	}
	return int(math.Floor(x + 0.5))
}

// A Codec represents an encoder.
type Codec struct {
	Dim   int     // Dimensionality, normally 2
	Scale float64 // Scale, normally 1e5
}

var defaultCodec = Codec{Dim: 2, Scale: 1e5}

// DecodeUint decodes a single unsigned integer from buf. It returns the
// decoded uint, the remaining unconsumed bytes of buf, and any error.
func DecodeUint(buf []byte) (uint, []byte, error) {
	var u, shift uint
	for i, b := range buf {
		switch {
		case 63 <= b && b < 95:
			u += (uint(b) - 63) << shift
			return u, buf[i+1:], nil
		case 95 <= b && b < 127:
			u += (uint(b) - 95) << shift
			shift += 5
		default:
			return 0, nil, errInvalidByte
		}
	}
	return 0, nil, errUnterminatedSequence
}

// DecodeInt decodes a single signed integer from buf. It returns the decoded
// int, the remaining unconsumed bytes of buf, and any error.
func DecodeInt(buf []byte) (int, []byte, error) {
	u, buf, err := DecodeUint(buf)
	if err != nil {
		return 0, nil, err
	}
	if u&1 == 0 {
		return int(u >> 1), buf, nil
	}
	return -int((u + 1) >> 1), buf, nil
}

// EncodeUint appends the encoding of a single unsigned integer u to buf.
func EncodeUint(buf []byte, u uint) []byte {
	for u >= 32 {
		buf = append(buf, byte((u&31)+95))
		u = u >> 5
	}
	buf = append(buf, byte(u+63))
	return buf
}

// EncodeInt appends the encoding of a single signed integer i to buf.
func EncodeInt(buf []byte, i int) []byte {
	var u uint
	if i < 0 {
		u = uint(^(i << 1))
	} else {
		u = uint(i << 1)
	}
	return EncodeUint(buf, u)
}

// DecodeCoord decodes a single coordinate from buf. It returns the coordinate,
// the remaining unconsumed bytes of buf, and any error.
func (c Codec) DecodeCoord(buf []byte) ([]float64, []byte, error) {
	coord := make([]float64, c.Dim)
	for i := range coord {
		var err error
		var j int
		j, buf, err = DecodeInt(buf)
		if err != nil {
			return nil, nil, err
		}
		coord[i] = float64(j) / c.Scale
	}
	return coord, buf, nil
}

// DecodeCoords decodes an array of coordinates from buf. It returns the
// coordinates, the remaining unconsumed bytes of buf, and any error.
func (c Codec) DecodeCoords(buf []byte) ([][]float64, []byte, error) {
	var coord []float64
	var err error
	coord, buf, err = c.DecodeCoord(buf)
	if err != nil {
		return nil, nil, err
	}
	coords := [][]float64{coord}
	for i := 1; len(buf) > 0; i++ {
		coord, buf, err = c.DecodeCoord(buf)
		if err != nil {
			return nil, nil, err
		}
		for j := range coord {
			coord[j] += coords[i-1][j]
		}
		coords = append(coords, coord)
	}
	return coords, nil, nil
}

// DecodeFlatCoords decodes coordinates from buf, appending them to a
// one-dimensional array. It returns the coordinates, the remaining unconsumed
// bytes in buf, and any error.
func (c Codec) DecodeFlatCoords(fcs []float64, buf []byte) ([]float64, []byte, error) {
	if len(fcs)%c.Dim != 0 {
		return nil, nil, errDimensionalMismatch
	}
	last := make([]int, c.Dim)
	for len(buf) > 0 {
		for j := 0; j < c.Dim; j++ {
			var err error
			var k int
			k, buf, err = DecodeInt(buf)
			if err != nil {
				return nil, nil, err
			}
			last[j] += k
			fcs = append(fcs, float64(last[j])/c.Scale)
		}
	}
	return fcs, nil, nil
}

// EncodeCoord encodes a single coordinate to buf.
func (c Codec) EncodeCoord(buf []byte, coord []float64) []byte {
	for _, x := range coord {
		buf = EncodeInt(buf, round(c.Scale*x))
	}
	return buf
}

// EncodeCoords appends the encoding of an array of coordinates coords to buf.
func (c Codec) EncodeCoords(buf []byte, coords [][]float64) []byte {
	last := make([]int, c.Dim)
	for _, coord := range coords {
		for i, x := range coord {
			ex := round(c.Scale * x)
			buf = EncodeInt(buf, ex-last[i])
			last[i] = ex
		}
	}
	return buf
}

// EncodeFlatCoords encodes a one-dimensional array of coordinates to buf.
func (c Codec) EncodeFlatCoords(buf []byte, fcs []float64) ([]byte, error) {
	if len(fcs)%c.Dim != 0 {
		return nil, errDimensionalMismatch
	}
	last := make([]int, c.Dim)
	for i, x := range fcs {
		ex := round(c.Scale * x)
		buf = EncodeInt(buf, ex-last[i%c.Dim])
		last[i%c.Dim] = ex
	}
	return buf, nil
}

// DecodeCoord decodes a single coordinate from buf using the default codec.
func DecodeCoord(buf []byte) ([]float64, []byte, error) {
	return defaultCodec.DecodeCoord(buf)
}

// DecodeCoords decodes an array of coordinates from buf using the default
// codec.
func DecodeCoords(buf []byte) ([][]float64, []byte, error) {
	return defaultCodec.DecodeCoords(buf)
}

// EncodeCoord returns the encoding of an array of coordinates using the
// default codec.
func EncodeCoord(coord []float64) []byte {
	return defaultCodec.EncodeCoord(nil, coord)
}

// EncodeCoords returns the encoding of an array of coordinates using the
// default codec.
func EncodeCoords(coords [][]float64) []byte {
	return defaultCodec.EncodeCoords(nil, coords)
}
