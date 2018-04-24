package mesq

type IncomMessageType struct {
	Datetime string `json:"datetime"`
	Events   struct {
		Event []EventType `json:"event"`
	} `json:"events"`
	ID       string `json:"id"`
	SystemID string `json:"systemId"`
	Token    string `json:"token"`
}

type EventType struct {
	MesID       string `json:"-"`
	Datetime    string `json:"datetime"`
	Description string `json:"description"`
	Filters     struct {
		Filter []struct {
			Persons struct {
				Person []struct {
					SSOID string `json:"SSOID"`
				} `json:"person"`
			} `json:"persons"`
		} `json:"filter"`
	} `json:"filters"`
	ID      string `json:"id"`
	Message struct {
		Parameters struct {
			Parameter []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"parameter"`
		} `json:"parameters"`
	} `json:"message"`
	StreamID int `json:"streamId"`
	TypeID   int `json:"typeId"`
}
