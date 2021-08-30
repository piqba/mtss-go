package mtssgo

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// Client is an interface that implements Mtss's API
type Client interface {
	// GetMtssJobs returns the corresponding jobs on fetch call, or an error.
	GetMtssJobs(
		ctx context.Context,
	) ([]Mtss, error)
}

// Client
// client represents a Mtss client. If the Debug field is set to an io.Writer
// (for example os.Stdout), then the client will dump API requests and responses
// to it.  To use a non-default HTTP client (for example, for testing, or to set
// a timeout), assign to the HTTPClient field. To set a non-default URL (for
// example, for testing), assign to the URL field.
type client struct {
	url        string
	httpClient *http.Client
	debug      io.Writer
}

// Mtss represents a object for Job from Mtss source
type Mtss struct {
	ID            int     `json:"id"`
	Company       string  `json:"organismo"`
	Position      string  `json:"cargo"`
	Taken         int     `json:"cantidad"`
	Entity        string  `json:"entidad"`
	Province      string  `json:"provincia"`
	Municipality  string  `json:"municipio"`
	Availability  int     `json:"ocupadas"`
	Activity      string  `json:"actividad"`
	Pay           float32 `json:"salario"`
	SchoolLevel   string  `json:"nivelEscolar"`
	Details       string  `json:"observaciones"`
	EntityMail    string  `json:"correo_entidad"`
	EntityAddress string  `json:"direccion_entidad"`
	EntityPhone   string  `json:"telefono_entidad"`
	RegisterDate  string  `json:"fecha_registro"`
	UniqueStamp   string  `json:"unique_stamp"`
	Enabled       bool    `json:"habilitada"`
	Source        string  `json:"source"`
	TypeWork      string  `json:"type_work"`
}

// ToMAP conver this struct to a simple map
func (mt *Mtss) ToMAP() (toHashMap map[string]interface{}, err error) {

	fromStruct, err := json.Marshal(mt)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(fromStruct, &toHashMap); err != nil {
		return toHashMap, err
	}

	return toHashMap, nil
}
