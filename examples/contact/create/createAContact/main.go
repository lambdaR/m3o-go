package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/contact"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Contact.Create(&contact.CreateRequest{
		Addresses: []contact.Address{
			contact.Address{
				Label:    "company address",
				Location: "123 street address",
			}},
		Birthday: "1995-01-01",
		Emails: []contact.Email{
			contact.Email{
				Address: "home@example.com",
				Label:   "home",
			}},
		Links: []contact.Link{
			contact.Link{
				Label: "blog",
				Url:   "https://blog.joe.me",
			}},
		Name: "joe",
		Note: "this person is very important",
		Phones: []contact.Phone{
			contact.Phone{
				Label:  "home",
				Number: "010-12345678",
			}},
	})
	fmt.Println(rsp, err)
}
