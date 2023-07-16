package zefix

import (
	"encoding/json"
	"github.com/denysvitali/sparql-client"
	"gorm.io/gorm/clause"
	"os"
)

// Import takes a .json file returned from a SPARQL query to the Zefix API and
// imports it into the database
func (c *Client) Import(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	var result sparql.Result
	dec := json.NewDecoder(f)
	err = dec.Decode(&result)
	if err != nil {
		return err
	}

	var companies []Company

	for _, v := range result.Results.Bindings {
		comp := Company{
			LegalName: v["legal_name"].Value,
			Name:      v["name"].Value,
			Uri:       v["company_uri"].Value,
			Locality:  v["locality"].Value,
			Type:      v["type"].Value,
			Address:   v["addresse"].Value,
		}
		companies = append(companies, comp)
	}

	tx := c.db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(companies, 1000)
	return tx.Error
}
