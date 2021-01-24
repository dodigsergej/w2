package parserdata

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fmt"

	"../db"
)

//MailData -- pojedinacan detalj iz maila
type MailData struct {
	signature       string //`json:"signature"`
	messageID       string //`json:"messageID"`
	originalMessage string //`json:"originalMessage"`
	deviceID        string //`json:"deviceID"`
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
	var maildata MailData

	reqBody, err := ioutil.ReadAll(r.Body)
	//fmt.Println(r.Body)
	/*
		err := json.NewDecoder(r.Body).Decode(&maildata)*/
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &maildata)
	//json.NewEncoder(w).Encode(&maildata)
	fmt.Fprintf(w, "Data: %+v", maildata)

	//for i := 0; i < len(maildata); i++ {
	if checkSignature(maildata.signature, maildata.messageID, maildata.deviceID) {
		fmt.Println("devID 1:" + maildata.deviceID)
		fmt.Println("Ispravan potpis")
		if !db.StoreData() {
			fmt.Println("Error ParseData")
		}
	} else {
		fmt.Println("devID 2:" + maildata.deviceID)
		fmt.Println("Nekorektna provera potpisa")
	}

	//}
}

func checkSignature(sigData string, messageID string, deviceID string) bool {
	var h = []byte(messageID + "|" + deviceID)
	if fmt.Sprintf("%x", h) == sigData {
		return true
	}

	return false

}
