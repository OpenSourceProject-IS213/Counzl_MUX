package new_user

import (
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
	"bytes"
	"encoding/gob"
	"homemade/converter"
)

type Users_MUX struct {
	Username     string
	Date_Created string
}

const (
	u_bucket_MUX = "./database/allUsers.db"
	u_highest_ID = "./database/highest_ID.db"
)

func newUser(brukernavn string) string {
	user := &Users_MUX{brukernavn, converter.GetTime()}

	u := string(user.serialize())

	var id = findHighestID()

	db, err := buntdb.Open(u_bucket_MUX)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(id, u, nil)
		return err
	})
	return id
}

/*
	Finn høyeste ID, inkrementer med 1 og returnér ID for ny bruker
*/
func findHighestID() string {
	var highestID string
	db, err := buntdb.Open(u_highest_ID)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *buntdb.Tx) error {
		val, _ := tx.Get("only_value")
		if val == "" {
			val = "1"
			highestID = converter.StringToInt_plus1(val)
		} else {
			highestID = converter.StringToInt_plus1(val)
		}
			_, _, err := tx.Set("only_value", highestID, nil)
		return err
	})
	return highestID
}

func PrintAllUsersByID() {
	var user *Users_MUX
	db, err := buntdb.Open(u_bucket_MUX)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *buntdb.Tx) error {
		var i = "1"

		for {
			val, err := tx.Get(string(i))
			if err == nil {
				user = deserialize([]byte(val))
				fmt.Printf("ID: %s, Brukernavn: %s, Opprettet: %s", i, user.Username, user.Date_Created)
				fmt.Println()
				i = converter.StringToInt_plus1(i)
			} else {
				break
			}
		}
		if err != nil {
			return err
		}
		return nil
	})
}

func PrintAllUsers() {
	db, err := buntdb.Open(u_bucket_MUX)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	erro := db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			if key != "" {
				val := deserialize([]byte(value))
				fmt.Printf("ID: %s, Brukernavn: %s, Opprettet: %s\n", key, val.Username, val.Date_Created)
				return true
			} else {
				return false
			}
		})

		return err
	})
	fmt.Printf("Error: '%s', hvis error er <nil> er det ikke flere brukere i DB", erro)
}

func (u *Users_MUX) serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(u)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func deserialize(d []byte) *Users_MUX {
	var users Users_MUX

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&users)
	if err != nil {
		log.Panic(err)
	}
	return &users
}
