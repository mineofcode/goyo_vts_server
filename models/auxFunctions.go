package models

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2"
	patterns "goyo.in/gpstracker/patterns"
)

const (
	vhclsCollection  = patterns.ColVhcls
	vhtrpsCollection = patterns.ColVhtrps
)

var session *mgo.Session
var errDial error

// acceso a la DB por cada consulta
func accessDB(collection string) *mgo.Collection {

	session, errDial = patterns.GetSessErrMongoDBSession("Dial")

	verifyErr(errDial)
	session.SetMode(mgo.Monotonic, true)
	col := session.DB(patterns.Dbname).C(collection)
	return col
}

// acceso a la DB por cada consulta
func getDBSession() *mgo.Session {
	if session == nil {
		fmt.Println("enter main - connecting to mongo")

		session, errDial = patterns.GetSessErrMongoDBSession("DialWithInfo")

		fmt.Println("enter main - dial")

		verifyErr(errDial)

		fmt.Println("enter main - dial")

		session.SetMode(mgo.Monotonic, true)
		fmt.Println("enter main - dial")

	}
	return session
}

func GetSession() *mgo.Session {
	return getDBSession()
}

func col(sess *mgo.Session, name string) *mgo.Collection {
	return sess.DB(patterns.Dbname).C(name)
}

// verifica el error
func verifyErr(err error) {
	if err != nil {
		fmt.Printf("\nError: %s\n", err)
		panic(err)
	}
}

// genera un password aleatoriamente
func generarPassword(longitud int) (cad string) {
	rand.Seed(time.Now().UTC().UnixNano())
	caracteres := "abcdefghijkmnpqrtuvwxyzABCDEFGHIJKLMNPQRTUVWXYZ2346789"
	contraseña := ""
	for i := 0; i < longitud; i++ {
		ln := rand.Intn(len(caracteres))
		contraseña += string(caracteres[ln])
	}
	return contraseña
}

// verifica si exite un elemento en una lista
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// invierte un string
func Encrypt(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
