package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	protoGuest "github.com/alexvassiliou/roganlist/proto"
	"github.com/golang/protobuf/proto"
)

func main() {
	// url := flag.String("url", "https://www.jrepodcast.com/guests/", "url to parse")
	// flag.Parse()

	// resp, err := http.Get(*url)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var guestHandler guest.Guests
	// guestHandler = guest.ParseHTML(resp.Body)

	// http.HandleFunc("/", redirect)

	// http.Handle("/guests/", guestHandler)

	// log.Fatal(http.ListenAndServe(":8080", nil))

	err := add("John Rambo", 1)
	if err != nil {
		log.Fatal(err)
	}

	_ = list()

}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/guests/", 301)
}

const dbPath = "mydb.pb"

func add(name string, ratio float32) error {
	guest := &protoGuest.Guest{
		Name:  name,
		Ratio: ratio,
	}
	g, err1 := proto.Marshal(guest)
	if err1 != nil {
		return err1
	}
	f, err2 := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err2 != nil {
		return err2
	}

	if err := gob.NewEncoder(f).Encode(int64(len(g))); err != nil {
		return err
	}

	_, err3 := f.Write(g)
	if err3 != nil {
		return err3
	}

	if err4 := f.Close(); err4 != nil {
		return err4
	}

	return nil
}

func list() error {
	g, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return err
	}

	for {
		if len(g) == 0 {
			return nil
		} else if len(g) < 4 {
			return fmt.Errorf("remain odd %d bytes, what to do", len(g))
		}

		var length int64
		if err := gob.NewDecoder(bytes.NewReader(g[:4])).Decode(&length); err != nil {
			return err
		}
		g = g[4:]

		var guest protoGuest.Guest
		if err := proto.Unmarshal(g[:length], &guest); err != nil {
			return err
		}
		g = g[length:]

		fmt.Println(guest.Name)
		fmt.Println(guest.Ratio)
	}
}
