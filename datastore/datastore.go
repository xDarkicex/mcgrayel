package datastore

import mgo "gopkg.in/mgo.v2"

// Session is our database session
var _session *mgo.Session

// Dial dials for dialing mongo server
func Dial() error {
	var err error
	_session, err = mgo.Dial("localhost")
	if err != nil {
		return err
	}
	err = _session.Ping()
	if err != nil {
		return err
	}
	_session.SetMode(mgo.Monotonic, true)
	return err
}

//GetSession ...
func GetSession() *mgo.Session {
	return _session.Clone()
}
