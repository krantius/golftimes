package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

var RanchoSanJoaquin = Course{
	Name:         "Rancho San Joaquin",
	BookingClass: 3869,
	ScheduleID:   4312,
}

var OsoCreek = Course{
	Name:         "Oso Creek",
	BookingClass: 3655,
	ScheduleID:   4102,
}

var MilesSquare = Course{
	Name:         "Miles Square",
	BookingClass: 3413,
	ScheduleID:   3759,
}

type Course struct {
	Name         string `json:"name"`
	BookingClass int    `json:"booking_class"`
	ScheduleID   int    `json:"schedule_id"`
}

type API struct {
	HTTPClient *http.Client
}

type TeeTimeResponse struct {
	TeesheetID                string `json:"teesheet_id"`
	TeesheetHoles             string `json:"teesheet_holes"`
	Time                      string `json:"time"`
	CourseID                  string `json:"course_id"`
	CourseName                string `json:"course_name"`
	ScheduleName              string `json:"schedule_name"`
	ScheduleID                string `json:"schedule_id"`
	AvailableSpots            int    `json:"available_spots"`
	MinimumPlayers            int    `json:"minimum_players"`
	TradeMinPlayers           int    `json:"trade_min_players"`
	TradeAvailablePlayers     int    `json:"trade_available_players"`
	ForeupTradeDiscountRate   int    `json:"foreup_trade_discount_rate"`
	Holes                     int    `json:"holes"`
	HasSpecial                bool   `json:"has_special"`
	SpecialDiscountPercentage int    `json:"special_discount_percentage"`
	GroupID                   bool   `json:"group_id"`
	RequireCreditCard         int    `json:"require_credit_card"`
	BookingClassID            int    `json:"booking_class_id"`
	BookingFeeRequired        bool   `json:"booking_fee_required"`
	BookingFeePrice           bool   `json:"booking_fee_price"`
	BookingFeePerPerson       bool   `json:"booking_fee_per_person"`
	GreenFeeTaxRate           bool   `json:"green_fee_tax_rate"`
	GreenFeeTax               int    `json:"green_fee_tax"`
	GuestGreenFeeTaxRate      bool   `json:"guest_green_fee_tax_rate"`
	GuestGreenFeeTax          int    `json:"guest_green_fee_tax"`
	CartFeeTaxRate            bool   `json:"cart_fee_tax_rate"`
	CartFeeTax                int    `json:"cart_fee_tax"`
	GuestCartFeeTaxRate       bool   `json:"guest_cart_fee_tax_rate"`
	GuestCartFeeTax           int    `json:"guest_cart_fee_tax"`
	SpecialID                 bool   `json:"special_id"`
	ForeupDiscount            bool   `json:"foreup_discount"`
	PayOnline                 string `json:"pay_online"`
	GreenFee                  int    `json:"green_fee"`
	CartFee                   int    `json:"cart_fee"`
	GuestGreenFee             int    `json:"guest_green_fee"`
	GuestCartFee              int    `json:"guest_cart_fee"`
	RateType                  string `json:"rate_type"`
}

func (a *API) GetTimes(course Course, date string) ([]*TeeTimeResponse, error) {
	req, err := http.NewRequest(http.MethodGet, "https://foreupsoftware.com/index.php/api/booking/times", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("time", "all")
	q.Set("date", date)
	q.Set("holes", "18")
	q.Set("players", "4")
	q.Set("booking_class", strconv.Itoa(course.BookingClass))
	q.Set("schedule_id", strconv.Itoa(course.ScheduleID))
	q.Set("specials_only", "0")
	q.Set("api_key", "no_limits")
	req.URL.RawQuery = q.Encode()

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	times := []*TeeTimeResponse{}
	if err := json.Unmarshal(b, &times); err != nil {
		return nil, err
	}

	return times, nil
}
