package dbcoachsequence

type CoachSequenceResponse struct {
	Data  CoachSequenceResponseData  `json:"data"`
	Error CoachSequenceResponseError `json:"error"`
}

type CoachSequenceResponseError struct {
	Id      int    `json:"id"`
	Message string `json:"msg"`
}

type CoachSequenceResponseData struct {
	IsFormation CoachSequenceResponseIsFormation `json:"istformation"`
}

type CoachSequenceResponseIsFormation struct {
	Direction    string                              `json:"fahrtrichtung"`
	VehicleGroup []CoachSequenceResponseVehicleGroup `json:"allFahrzeuggruppe"`
}

type CoachSequenceResponseVehicleGroup struct {
	Vehicles []CoachSequenceResponseVehicle `json:"allFahrzeug"`
}

type CoachSequenceResponseVehicle struct {
	Status        string                                `json:"status"`
	CoachNumber   string                                `json:"wagenordnungsnummer"`
	CoachType     string                                `json:"fahrzeugtyp"`
	Position      string                                `json:"position"`
	Orientation   string                                `json:"orientierung"`
	VehicleNumber string                                `json:"fahrzeugnummer"`
	Category      string                                `json:"kategorie"`
	Features      []CoachSequenceResponseVehicleFeature `json:"allFahrzeugausstattung"`
}

type CoachSequenceResponseVehicleFeature struct {
	Quantity    string `json:"anzahl"`
	Type        string `json:"ausstattungsart"`
	Description string `json:"bezeichnung"`
	Status      string `json:"status"`
}

type DbCoachSequenceClient struct {
	BaseURL string
}

// Friendly structure for general use
type CoachSequence struct {
	Coaches []Coach `json:"coaches"`
}

type Coach struct {
	VehicleNumber            string `json:"vehicle_number"`
	CoachNumber              string `json:"coach_number"`
	HasAc                    bool   `json:"has_ac"`
	HasBahnBonusSeats        bool   `json:"has_bahn_bonus_seats"`
	HasAccessibleSeats       bool   `json:"has_accessible_seats"`
	HasFamilyCompartment     bool   `json:"has_family_compartment"`
	HasWheelchairCompartment bool   `json:"has_wheelchair_compartment"`
	HasAccessibleToilet      bool   `json:"has_accessible_toilet"`
	HasQuietArea             bool   `json:"has_quiet_area"`
	HasBikeSpace             bool   `json:"has_bike_space"`
	HasInfo                  bool   `json:"has_info"`
	TravelClass              int    `json:"travel_class"`
}
