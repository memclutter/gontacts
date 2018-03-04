package components

import "gopkg.in/mgo.v2"

var (
	MongoDB *mgo.Database
)

func MongoInit(mongoUrl string) error {
	dialInfo, err := mgo.ParseURL(mongoUrl)
	if err != nil {
		return err
	}

	ms, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return err
	}

	MongoDB = ms.DB(dialInfo.Database)

	return nil
}

func MongoClose() {
	MongoDB.Session.Close()
}
