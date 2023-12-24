package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config) error {
	if c.next == nil && c.previous != nil {
		return errors.New("no next page")
	}
	resp, err := c.client.GetLocationAreas(c.next)
	if err != nil {
		return err
	}
	c.next = resp.Next
	c.previous = resp.Previous
	for _, v := range resp.Results {
		fmt.Println(v.Name)
	}
	return nil
}