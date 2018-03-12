package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// CategoriserClient is an interface for http clients.
type CategoriserClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Categoriser containse the client for retrieving locations.
type Categoriser struct {
	client CategoriserClient
}

// SetClient sets the client.
func (c *Categoriser) SetClient(client CategoriserClient) {
	c.client = client
}

// Categorise determines the zombie status of a driver.
func (c *Categoriser) Categorise(d *Driver, ls []Location) {
	if len(ls) > 1 {
		uniq := make(map[Location]bool)
		for _, location := range ls {
			uniq[location] = true
		}

		if len(uniq) == 1 {
			d.Zombie = true
		} else {
			d.Zombie = false
		}
	}
}

// GetLocations fetches the locations for the last 5 minutes.
func (c *Categoriser) GetLocations(d Driver) []Location {
	request := c.newLocationsRequest(d)
	client := c.newClient()

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	locations := []Location{}
	err = decoder.Decode(&locations)
	if err != nil {
		log.Fatal(err)
	}

	return locations
}

func (c *Categoriser) newClient() CategoriserClient {
	if c.client == nil {
		return &http.Client{}
	}
	return c.client
}

func (c *Categoriser) newLocationsRequest(d Driver) *http.Request {
	locationURL := c.newLocationURL(d)
	buffer := strings.NewReader("")
	request, err := http.NewRequest("GET", locationURL.String(), buffer)
	if err != nil {
		log.Fatal(err)
	}

	// request.Host = locationURL.Host
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Zombie/0.42")

	return request
}

func (c *Categoriser) newLocationURL(d Driver) *url.URL {
	stringID := strconv.Itoa(int(d.ID))

	host := getenv("LOCATION_HOST", "location:3001")
	path := "/drivers/" + stringID + "/coordinates"

	return &url.URL{Scheme: "http", Host: host, Path: path}
}

func getenv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return fallback
}
