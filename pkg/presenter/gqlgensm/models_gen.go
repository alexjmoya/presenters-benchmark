// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgensm

import (
	"fmt"
	"io"
	"rfc/presenters/pkg/common"
	"rfc/presenters/pkg/domainHotelCommon"
	"strconv"
)

type AddOn struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AddOns struct {
	Distribute   *string `json:"distribute"`
	Distribution []AddOn `json:"distribution"`
}

type BookableOptionSearch interface {
	IsBookableOptionSearch()
}

type CriteriaSearch struct {
	CheckIn     string         `json:"checkIn"`
	CheckOut    string         `json:"checkOut"`
	Hotels      []string       `json:"hotels"`
	Occupancies []RoomCriteria `json:"occupancies"`
	Language    *string        `json:"language"`
	Currency    *string        `json:"currency"`
	Nationality *string        `json:"nationality"`
	Market      string         `json:"market"`
}

type HotelCriteriaSearchInput struct {
	CheckIn      string                        `json:"checkIn"`
	CheckOut     string                        `json:"checkOut"`
	Hotels       []string                      `json:"hotels"`
	Destinations []string                      `json:"destinations"`
	Occupancies  []domainHotelCommon.Occupancy `json:"occupancies"`
	Language     *string                       `json:"language"`
	Currency     *string                       `json:"currency"`
	Nationality  *string                       `json:"nationality"`
	Market       *string                       `json:"market"`
}

type HotelSearch struct {
	Context         *string                     `json:"context"`
	Stats           *StatsRequest               `json:"stats"`
	AuditData       *common.AuditData           `json:"auditData"`
	RequestCriteria *CriteriaSearch             `json:"requestCriteria"`
	Options         []*domainHotelCommon.Option `json:"options"`
	Errors          []common.AdviseMessage      `json:"errors"`
	Warnings        []common.AdviseMessage      `json:"warnings"`
}

func (HotelSearch) IsResponse() {}

type HotelXQuery struct {
	Search              *HotelSearch  `json:"search"`
	SearchStatusService ServiceStatus `json:"searchStatusService"`
}

type Priceable interface {
	IsPriceable()
}

type Response interface {
	IsResponse()
}

type RoomCriteria struct {
	Paxes []domainHotelCommon.Pax `json:"paxes"`
}

type Search struct {
	Hotel *HotelSearch `json:"hotel"`
}

type ServiceStatus struct {
	Code        *string `json:"code"`
	Type        *string `json:"type"`
	Description *string `json:"description"`
}

type Stat struct {
	Start    string   `json:"start"`
	End      string   `json:"end"`
	Duration *float64 `json:"duration"`
}

type StatAccess struct {
	Name                string            `json:"name"`
	Total               Stat              `json:"total"`
	StaticConfiguration *Stat             `json:"staticConfiguration"`
	Hotels              int               `json:"hotels"`
	Zones               int               `json:"zones"`
	Cities              int               `json:"cities"`
	RequestAccess       *StatPlugin       `json:"requestAccess"`
	ResponseAccess      *StatPlugin       `json:"responseAccess"`
	Transactions        []StatTransaction `json:"transactions"`
	Plugins             []StatPlugin      `json:"plugins"`
}

type StatPlugin struct {
	Name  string `json:"name"`
	Total Stat   `json:"total"`
}

type StatTransaction struct {
	Reference           string `json:"reference"`
	Total               Stat   `json:"total"`
	BuildRequest        Stat   `json:"buildRequest"`
	WorkerCommunication Stat   `json:"workerCommunication"`
	ParseResponse       Stat   `json:"parseResponse"`
}

type StatsRequest struct {
	Total          Stat         `json:"total"`
	Validation     Stat         `json:"validation"`
	Process        Stat         `json:"process"`
	Configuration  Stat         `json:"configuration"`
	Request        Stat         `json:"request"`
	Response       Stat         `json:"response"`
	RequestPlugin  *StatPlugin  `json:"requestPlugin"`
	ResponsePlugin *StatPlugin  `json:"responsePlugin"`
	Hotels         int          `json:"hotels"`
	Zones          int          `json:"zones"`
	Cities         int          `json:"cities"`
	DockerID       string       `json:"dockerID"`
	Accesses       []StatAccess `json:"Accesses"`
}

type PriceType string

const (
	PriceTypeGross  PriceType = "GROSS"
	PriceTypeNet    PriceType = "NET"
	PriceTypeAmount PriceType = "AMOUNT"
)

func (e PriceType) IsValid() bool {
	switch e {
	case PriceTypeGross, PriceTypeNet, PriceTypeAmount:
		return true
	}
	return false
}

func (e PriceType) String() string {
	return string(e)
}

func (e *PriceType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PriceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PriceType", str)
	}
	return nil
}

func (e PriceType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ServiceType string

const (
	ServiceTypeSkiPass ServiceType = "SKI_PASS"
)

func (e ServiceType) IsValid() bool {
	switch e {
	case ServiceTypeSkiPass:
		return true
	}
	return false
}

func (e ServiceType) String() string {
	return string(e)
}

func (e *ServiceType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ServiceType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ServiceType", str)
	}
	return nil
}

func (e ServiceType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
