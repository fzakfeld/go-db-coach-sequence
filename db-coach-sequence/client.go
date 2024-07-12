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

	for _, vehicleGroup := range response.Data.IsFormation.VehicleGroup {
		train := Train{}

		for _, vehicle := range vehicleGroup.Vehicles {
			coach := Coach{
				CoachType:           vehicle.CoachType,
				CoachSequenceNumber: vehicle.CoachNumber,
			}

			classNumber := vehicle.VehicleNumber[:len(vehicle.VehicleNumber)-4][len(vehicle.VehicleNumber)-7:]
			coachNumber := vehicle.VehicleNumber[:len(vehicle.VehicleNumber)-1][len(vehicle.VehicleNumber)-4:]

			coach.CoachNumber = classNumber + " " + coachNumber

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
			} else if vehicle.Category == "STEUERWAGENZWEITEKLASSE" {
				coach.TravelClass = 2
			} else if vehicle.Category == "REISEZUGWAGENERSTEKLASSE" {
				coach.TravelClass = 1
			} else if vehicle.Category == "STEUERWAGENERSTEKLASSE" {
				coach.TravelClass = 1
			} else if vehicle.Category == "HALBSPEISEWAGENERSTEKLASSE" {
				coach.TravelClass = 1
				coach.HasRestaurant = true
			}

			train.Coaches = append(train.Coaches, coach)

			if train.Class == "" {
				classes := []string{"415", "412", "411", "408", "407", "406", "403", "402", "401"}

				for _, c := range classes {
					if classNumber == c {
						train.Class = classNumber
						break
					}
				}
			}
		}

		specialLiveries := map[string]string{
			"ICE0304": "railbow",
			"ICE9457": "deutschland",
		}
		if livery, ok := specialLiveries[vehicleGroup.Name]; ok {
			train.Livery = livery
		}

		coachSequence.Trains = append(coachSequence.Trains, train)
	}

	return coachSequence, nil
}
