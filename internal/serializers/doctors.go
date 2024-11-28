package serializers

import "encoding/json"

// Root represents the root of the JSON structure
type Root struct {
	Search          Search          `json:"search"`
	OrderItems      OrderItems      `json:"order_items"`
	Categories      []Category      `json:"categories"`
	Filters         []Filter        `json:"filters"`
	SelectedFilters SelectedFilters `json:"selected_filters"`
	SEOInfo         SEOInfo         `json:"seo_info"`
	ConsultFooter   string          `json:"consult_footer"`
	Footers         []Footer        `json:"footers"`
}

// Search represents the search object
type Search struct {
	QueryID    string     `json:"query_id"`
	Total      int        `json:"total"`
	IsLanding  bool       `json:"is_landing"`
	Pagination Pagination `json:"pagination"`
	Result     []Result   `json:"result"`
}

// Pagination represents pagination details
type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// Result represents a search result
type Result struct {
	ID                    string            `json:"id"`
	ServerID              int               `json:"server_id"`
	Type                  string            `json:"type"`
	Title                 string            `json:"title"`
	Prefix                string            `json:"prefix"`
	Image                 string            `json:"image"`
	View                  string            `json:"view"`
	DisplayExpertise      string            `json:"display_expertise"`
	Satisfaction          int               `json:"satisfaction"`
	RatesCount            int               `json:"rates_count"`
	Centers               []Center          `json:"centers"`
	DisplayAddress        string            `json:"display_address"`
	Badges                json.RawMessage   `json:"badges"`
	IsBulk                bool              `json:"is_bulk"`
	Price                 string            `json:"price"`
	ConsultActiveBooking  bool              `json:"consult_active_booking"`
	PresenceActiveBooking bool              `json:"presence_active_booking"`
	URL                   string            `json:"url"`
	Actions               []Action          `json:"actions"`
	Experience            interface{}       `json:"experience"`
	Position              int               `json:"position"`
	HasPresciption        bool              `json:"has_presciption"`
	Insurances            []string          `json:"insurances"`
	ExperimentDetails     ExperimentDetails `json:"experiment_details"`
	PresenceFreeturn      interface{}       `json:"presence_freeturn"`
	Expertises            []Expertise       `json:"expertises"`
	Gender                int               `json:"gender"`
	Expertise             []string          `json:"expertise"`
	RateInfo              RateInfo          `json:"rate_info"`
	ConsultServices       []ConsultService  `json:"consult_services"`
	DoctorID              string            `json:"doctor_id"`
	FreeturnsInfo         []FreeturnInfo    `json:"freeturns_info"`
	NumberOfVisits        int               `json:"number_of_visits"`
	WaitingTimeInfo       interface{}       `json:"waiting_time_info"`
	Slug                  string            `json:"slug"`
	GraduationDate        interface{}       `json:"graduation_date"`
	Star                  float64           `json:"star"`
	AvailableTimeStatus   int               `json:"available_time_status"`
	Services              []Service         `json:"services"`
	UniversityName        []string          `json:"university_name"`
	DisplayName           string            `json:"display_name"`
	RecordType            string            `json:"record_type"`
	CenterID              []string          `json:"center_id"`
	ConsultFreeturn       int64             `json:"consult_freeturn"`
	Name                  string            `json:"name"`
	CalculatedRate        int               `json:"calculated_rate"`
	MedicalCode           string            `json:"medical_code"`
}

// Center represents a center object
type Center struct {
	ID            string `json:"id"`
	Status        int    `json:"status"`
	UserCenterID  string `json:"user_center_id"`
	ServerID      int    `json:"server_id"`
	Name          string `json:"name"`
	DisplayNumber string `json:"display_number"`
	Address       string `json:"address"`
	ProvinceName  string `json:"province_name"`
	CityName      string `json:"city_name"`
	CenterType    int    `json:"center_type"`
	Map           Map    `json:"map"`
	ActiveBooking bool   `json:"active_booking"`
}

// Map represents latitude and longitude details
type Map struct {
	Lat json.RawMessage `json:"lat"`
	Lon json.RawMessage `json:"lon"`
}

// Action represents an action object
type Action struct {
	Title    string `json:"title"`
	Outline  bool   `json:"outline"`
	TopTitle string `json:"top_title"`
	URL      string `json:"url"`
}

// ExperimentDetails represents experiment details
type ExperimentDetails struct {
	SearchIndex        string `json:"search_index"`
	ConsultSearchIndex string `json:"consult_search_index"`
}

// Expertise represents expertise information
type Expertise struct {
	Degree          Degree           `json:"degree"`
	AliasTitle      string           `json:"alias_title"`
	Expertise       ExpertiseDetails `json:"expertise"`
	ExpertiseGroups []ExpertiseGroup `json:"expertise_groups"`
}

// Degree represents degree information
type Degree struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ExpertiseDetails represents detailed expertise
type ExpertiseDetails struct {
	Name string `json:"name"`
}

// ExpertiseGroup represents an expertise group
type ExpertiseGroup struct {
	Name string `json:"name"`
}

// RateInfo represents rate information
type RateInfo struct {
	WaitingTime        float64 `json:"waiting_time"`
	WaitingTimeCount   int     `json:"waiting_time_count"`
	DoctorEncounter    float64 `json:"doctor_encounter"`
	QualityOfTreatment float64 `json:"quality_of_treatment"`
	Rate               float64 `json:"rate"`
	CommentsCount      int     `json:"comments_count"`
	RatesCount         int     `json:"rates_count"`
	CountDislike       int     `json:"count_dislike"`
}

// ConsultService represents a consultation service
type ConsultService struct {
	ID        string `json:"id"`
	FreePrice int64  `json:"free_price"`
}

// FreeturnInfo represents free turn information
type FreeturnInfo struct {
	CenterID      string `json:"center_id"`
	AvailableTime string `json:"available_time"`
	Freeturn      string `json:"freeturn"`
}

// Service represents a service
type Service struct {
	ID        string     `json:"id"`
	CenterID  string     `json:"center_id"`
	Workhours []Workhour `json:"workhours"`
}

// Workhour represents work hours
type Workhour struct {
	Types []Type `json:"types"`
}

// Type represents a type of work hour
type Type struct {
	Name string `json:"name"`
}

// OrderItems represents the order items
type OrderItems struct {
	Clinic                string `json:"clinic"`
	ClinicDoctorPopular   string `json:"clinic_doctor_popular"`
	ClinicFirstFreeturn   string `json:"clinic_first_freeturn"`
	ClinicLessWaitingTime string `json:"clinic_less_waiting_time"`
	ClinicDoctorVisits    string `json:"clinic_doctor_visits"`
}

// Category represents a category
type Category struct {
	Title         string        `json:"title"`
	Image         string        `json:"image"`
	URL           string        `json:"url"`
	Value         string        `json:"value"`
	SubCategories []SubCategory `json:"sub_categories"`
	Count         int           `json:"count"`
}

// SubCategory represents a sub-category
type SubCategory struct {
	Title string `json:"title"`
	Value string `json:"value"`
	URL   string `json:"url"`
}

// Filter represents a filter object
type Filter struct {
	Title string `json:"title"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Items []Item `json:"items"`
}

// Item represents an item in a filter
type Item struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Count int    `json:"count"`
}

// SelectedFilters represents selected filters
type SelectedFilters struct {
	TurnType string `json:"turn_type"`
	Category string `json:"category"`
}

// SEOInfo represents SEO information
type SEOInfo struct {
	ID              int          `json:"id"`
	Name            interface{}  `json:"name"`
	Title           string       `json:"title"`
	Description     string       `json:"description"`
	Heading         string       `json:"heading"`
	PageDescription string       `json:"page_description"`
	SEOBox          string       `json:"seo_box"`
	Content         string       `json:"content"`
	CanonicalLink   string       `json:"canonical_link"`
	Breadcrumbs     []Breadcrumb `json:"breadcrumbs"`
	JSONLD          []JSONLD     `json:"jsonld"`
}

// Breadcrumb represents a breadcrumb
type Breadcrumb struct {
	Text string `json:"text"`
	Href string `json:"href"`
}

// JSONLD represents JSON-LD information
type JSONLD struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

// Footer represents a footer object
type Footer struct {
	Title string       `json:"title"`
	Items []FooterItem `json:"items"`
}

// FooterItem represents an item in a footer
type FooterItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
