package parserdata

import (
	"crypto/md5"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"fmt"

	"../db"
)

//MailData -- pojedinacan detalj iz maila
type MailData struct {
	Signature       string `json:"signature"`
	MessageID       string `json:"messageID"`
	OriginalMessage string `json:"originalMessage"`
	DeviceID        string `json:"deviceID"`
	/*
		deviceType string
			deviceVersion             string
			deviceStatus              string
			deviceDatetime            string
			deviceMainbattery         string
			deviceGsmbattery          string
			redniBrojZapisa           string
			statusZapisa              string
			tipZapisa                 string
			datumVremeZapisa          string
			registarA                 string
			registarB                 string
			registarC                 string
			registarD                 string
			leakAlarm                 string
			tamperAlarm               string
			rezimRadaUlaza            string
			kontrolnaSuma             string
			kodDogadjaja              string
			stanjeUlaza               string
			poslednjiRssi             string
			intervalLogera            string
			oznakaSadrzaja            string
			identificationNr          string
			meterIndex                string
			sndNr                     string
			manufacturer              string
			deviceid                  string
			verzija                   string
			medijum                   string
			konfiguracija             string
			totalRegistar             string
			reverseRegistar           string
			datumVreme                string
			totalRegistarNaDanCitanja string
			datumCitanja              string
			statusGresaka             string
			magneticTamperHour        string
			jacinaSignala             string
	*/
}

//ParseData -- parsiranje pristiglih detalja
func ParseData(w http.ResponseWriter, r *http.Request) {
	var maildata []MailData
	t := time.Now()
	var filename string = t.Format("./log/20060102")

	f, err := os.OpenFile(filename+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &maildata)

	//fmt.Printf("%+v\n", maildata)
	//fmt.Println(len(maildata))
	logger.Println("Broj poruka: " + fmt.Sprint(len(maildata)))

	for i := 0; i < len(maildata); i++ {
		log.Printf("%+v\n", maildata[i])
		logger.Printf("%+v\n", maildata[i])

		if checkSignature(maildata[i].Signature, maildata[i].MessageID, maildata[i].DeviceID) {
			fmt.Println("Ispravan potpis")
			if !db.StoreData() {
				fmt.Println("Error ParseData")
			}
		} else {
			fmt.Println("Neispravan potpis")
		}
	}
}

func checkSignature(sigData string, messageID string, deviceID string) bool {
	var h = []byte(messageID + "|" + deviceID)

	if fmt.Sprintf("%x", md5.Sum(h)) == sigData {
		return true
	}

	return false

}
