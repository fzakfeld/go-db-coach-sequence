package dbcoachsequence

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func NewDbCoachSequenceClient() *DbCoachSequenceClient {
	return &DbCoachSequenceClient{
		BaseURL: "https://ist-wr.noncd.db.de/wagenreihung/1.0",
	}
}

func (c *DbCoachSequenceClient) GetSequence(trainNumber string, timestamp string) (CoachSequence, error) {
	var coachSequence CoachSequence

	url := c.BaseURL + "/" + trainNumber + "/" + timestamp

	resp, err := http.Get(url)

	if err != nil {
		return coachSequence, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return coachSequence, err
	}

	var response CoachSequenceResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return coachSequence, err
	}

	if response.Error.Id != 0 {
		return coachSequence, errors.New(response.Error.Message)
	}

	for _, vehicle := range response.Data.IsFormation.VehicleGroup[0].Vehicles {
		coach := Coach{
			VehicleNumber: vehicle.VehicleNumber,
			CoachNumber:   vehicle.CoachNumber,
		}

		for _, feature := range vehicle.Features {
			if feature.Type == "KLIMA" {
				coach.HasAc = true
			}

			if feature.Type == "PLAETZEBAHNCOMFORT" {
				coach.HasBahnBonusSeats = true
			}

			if feature.Type == "PLAETZESCHWERBEH" {
				coach.HasAccessibleSeats = true
			}

			if feature.Type == "FAMILIE" {
				coach.HasFamilyCompartment = true
			}

			if feature.Type == "PLAETZEROLLSTUHL" {
				coach.HasFamilyCompartment = true
			}

			if feature.Type == "ROLLSTUHLTOILETTE" {
				coach.HasAccessibleToilet = true
			}

			if feature.Type == "RUHE" {
				coach.HasQuietArea = true
			}

			if feature.Type == "PLAETZEFAHRRAD" {
				coach.HasBikeSpace = true
			}

			if feature.Type == "INFO" {
				coach.HasInfo = true
			}
		}

		if vehicle.Category == "REISEZUGWAGENZWEITEKLASSE" {
			coach.TravelClass = 2
		} else if vehicle.Category == "REISEZUGWAGENERSTEKLASSE" {
			coach.TravelClass = 1
		}

		coachSequence.Coaches = append(coachSequence.Coaches, coach)
	}

	return coachSequence, nil
}
