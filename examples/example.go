package main

import (
	"encoding/json"

	dbcoachsequence "github.com/fzakfeld/go-db-coach-sequence/db-coach-sequence"
)

func main() {
	coachSequenceClient := dbcoachsequence.NewDbCoachSequenceClient()

	coachSequence, err := coachSequenceClient.GetSequence("373", "202407121800")

	if err != nil {
		panic(err)
	}

	foo, _ := json.Marshal(&coachSequence)
	println(string(foo))

}
