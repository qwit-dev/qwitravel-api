package apis

// api base
const BKKBase = "https://futar.bkk.hu/api/query/v1/ws"

// api dialect
const BKKDialect = "otp" // can be "mobile" or "otp"

// stop related
const BKKStopSchedule = BKKBase + "/" + BKKDialect + "/api/where/schedule-for-stop"
const BKKStopArrivalsAndDepartures = BKKBase + "/" + BKKDialect + "/api/where/arrivals-and-departures-for-stop"
