package models

type BKKStopSchedule struct {
	CurrentTime int64  `json:"currentTime"`
	Version     int    `json:"version"`
	Status      string `json:"status"`
	Code        int    `json:"code"`
	Text        string `json:"text"`
	Data        struct {
		LimitExceeded bool `json:"limitExceeded"`
		Entry         struct {
			StopID        string   `json:"stopId"`
			Date          string   `json:"date"`
			RouteIds      []string `json:"routeIds"`
			NearbyStopIds []string `json:"nearbyStopIds"`
			AlertIds      []string `json:"alertIds"`
			Schedules     []struct {
				RouteID    string   `json:"routeId"`
				AlertIds   []string `json:"alertIds"`
				Directions []struct {
					DirectionID string                 `json:"directionId"`
					Groups      map[string]interface{} `json:"groups"`
					StopTimes   []struct {
						StopID                 string   `json:"stopId"`
						StopHeadsign           string   `json:"stopHeadsign"`
						ArrivalTime            int      `json:"arrivalTime"`
						DepartureTime          int      `json:"departureTime"`
						PredictedArrivalTime   int      `json:"predictedArrivalTime"`
						PredictedDepartureTime int      `json:"predictedDepartureTime"`
						Uncertain              bool     `json:"uncertain"`
						TripID                 string   `json:"tripId"`
						ServiceDate            string   `json:"serviceDate"`
						WheelchairAccessible   bool     `json:"wheelchairAccessible"`
						MayRequireBooking      bool     `json:"mayRequireBooking"`
						AlertIds               []string `json:"alertIds"`
					} `json:"stopTimes"`
				} `json:"directions"`
			} `json:"schedules"`
		} `json:"entry"`
		References struct {
			// Routes map[string]interface{} `json:"routes"`
			// Alerts map[string]interface{} `json:"alerts"`
		} `json:"references"`
		Class string `json:"class"`
	} `json:"data"`
}

type BKKStopArrivalsAndDepartures struct {
	CurrentTime int64  `json:"currentTime"`
	Version     int    `json:"version"`
	Status      string `json:"status"`
	Code        int    `json:"code"`
	Text        string `json:"text"`
	Data        struct {
		LimitExceeded bool `json:"limitExceeded"`
		Entry         struct {
			StopID        string   `json:"stopId"`
			RouteIds      []string `json:"routeIds"`
			AlertIds      []string `json:"alertIds"`
			NearbyStopIds []string `json:"nearbyStopIds"`
			StopTimes     []struct {
				StopID                 string   `json:"stopId"`
				StopHeadsign           string   `json:"stopHeadsign"`
				ArrivalTime            int      `json:"arrivalTime"`
				DepartureTime          int      `json:"departureTime"`
				PredictedArrivalTime   int      `json:"predictedArrivalTime"`
				PredictedDepartureTime int      `json:"predictedDepartureTime"`
				Uncertain              bool     `json:"uncertain"`
				TripID                 string   `json:"tripId"`
				ServiceDate            string   `json:"serviceDate"`
				WheelchairAccessible   bool     `json:"wheelchairAccessible"`
				MayRequireBooking      bool     `json:"mayRequireBooking"`
				AlertIds               []string `json:"alertIds"`
			} `json:"stopTimes"`
		} `json:"entry"`
		References struct {
			// Routes map[string]interface{} `json:"routes"`
			// Trips  map[string]interface{} `json:"trips"`
			// Alerts map[string]interface{} `json:"alerts"`
		} `json:"references"`
		Class string `json:"class"`
	} `json:"data"`
}

type BKKStopRouteDetails struct {
	CurrentTime int64  `json:"currentTime"`
	Version     int    `json:"version"`
	Status      string `json:"status"`
	Code        int    `json:"code"`
	Text        string `json:"text"`
	Data        struct {
		List []struct {
			ID           string `json:"id"`
			ShortName    string `json:"shortName"`
			LongName     string `json:"longName"`
			Description  string `json:"description"`
			Type         string `json:"type"`
			URL          string `json:"url"`
			AgencyID     string `json:"agencyId"`
			BikesAllowed bool   `json:"bikesAllowed"`
			Style        struct {
				Color string `json:"color"`
				Stop  struct {
					Colors []string `json:"colors"`
					Type   string   `json:"type"`
					Image  string   `json:"image"`
				} `json:"stop"`
				Icon struct {
					Type      string `json:"type"`
					Text      string `json:"text"`
					TextColor string `json:"textColor"`
				} `json:"icon"`
				VehicleIcon struct {
					Name string `json:"name"`
				} `json:"vehicleIcon"`
			} `json:"style"`
			SortOrder int `json:"sortOrder"`
		} `json:"list"`
		OutOfRange    bool `json:"outOfRange"`
		LimitExceeded bool `json:"limitExceeded"`
		References    struct {
		} `json:"references"`
		Class string `json:"class"`
	} `json:"data"`
}
