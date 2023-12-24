package main

import (
	"errors"
	"fmt"
)

func commandMapb(c *config) error {
	if c.previous == nil {
		return errors.New("no previous page")
	}
	resp, err := c.client.GetLocationAreas(c.previous)
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