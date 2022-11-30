// Copyright © 2017 Julien Pivotto <roidelapluie@inuits.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Companies stores the companies we have in cache or fetched from the server.
var Companies = &CompaniesList{}

// Company is a company as seen in the ninetofiver api
type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

// CompaniesList is a list of companies in the ninetofiver api
type CompaniesList struct {
	Companies []Company `json:"results"`
}

func (cs *CompaniesList) apiURL() string {
	return "v2/companies"
}

// Slug is used to represent the model in cli
func (cs *CompaniesList) Slug() string {
	return "companies"
}

func (cs *CompaniesList) augment(c *Client) error {
	return nil
}

func (cs *CompaniesList) isEmpty() bool {
	return len(cs.Companies) == 0
}

// HasPorcelain returns false because companies do not support PorcelainPrettyPrint
func (cs *CompaniesList) HasPorcelain() bool {
	return false
}

// GetColumns returns an empty list because companies are not displayed as a table
func (cs *CompaniesList) GetColumns() []string {
	return *new([]string)
}

// GetDefaultColumns returns an empty list because companies are not displayed as a table
func (cs *CompaniesList) GetDefaultColumns() []string {
	return *new([]string)
}

// PorcelainPrettyPrint does nothing for this model
func (cs *CompaniesList) PorcelainPrettyPrint() {
	return
}

// Get returns the Company from the server
func (c *Company) Get(client Client) error {
	cs := &CompaniesList{}
	resp, err := client.GetRequest(fmt.Sprintf("%s/v2/companies/?name=%s", client.Endpoint, c.Name))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(cs)
	if err != nil {
		return err
	}
	if len(cs.Companies) != 1 {
		return fmt.Errorf("Expected 1 company, got %d", len(cs.Companies))
	}
	*c = cs.Companies[0]
	return nil
}

// Slug returns the short name for the company model
func (c *Company) Slug() string {
	return "company"
}

func (c *Company) apiURL() string {
	return "v2/companies"
}

// SetID sets the ID of a company
func (c *Company) SetID(i int) {
	c.ID = i
}

// Augment populates extra fields for a company
func (c *Company) Augment(cl *Client) {
	return
}

// GetID returns the ID of a company
func (c *Company) GetID() int {
	return c.ID
}

// DeleteArg returns the arg needed to delete the company
func (c *Company) DeleteArg() string {
	return strconv.Itoa(c.ID)
}

// GetByID returns the Company from the server
func (c *Company) GetByID(client Client) error {
	resp, err := client.GetRequest(fmt.Sprintf("%s/v2/companies/%d/", client.Endpoint, c.ID))

	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(c)
	if err != nil {
		return err
	}
	return nil
}

// PrettyPrint prints companies in a nice way to the console
func (cs *CompaniesList) PrettyPrint(columns []string) {
	for _, c := range cs.Companies {
		c.PrettyPrint()
	}
}

// PrettyPrint prints company in a nice way to the console
func (c *Company) PrettyPrint() {
	fmt.Printf("%s [%s]\n", c.Name, c.Country)
}

func (cs *CompaniesList) extraFetchParameters(c Client, args []string) string {
	return ""
}

func init() {
	cache.register(Companies)
	Models.register(Companies, &Company{})
}
