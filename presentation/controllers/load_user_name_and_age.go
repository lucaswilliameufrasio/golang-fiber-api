package controllers

import (
	"fmt"
	presentationprotocols "lucaswilliameufrasio/golang-fiber-api/presentation/protocols"
)

// // LoadUserNameAndAgeController to get user name and age
// func LoadUserNameAndAgeController(c *fiber.Ctx) error {
// 	msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
// 	fmt.Println(c.Params())
// 	return c.Status(200).JSON(fiber.Map{
// 		"info":   msg,
// 		"active": true,
// 	}) // => info: 👴 john is 75 years old
// }

// LoadUserNameAndAgeController to get user name and age
func LoadUserNameAndAgeController(request *presentationprotocols.HTTPRequest) presentationprotocols.HTTPResponse {
	msg := fmt.Sprintf("👴 %s is %s years old", request.Params("name"), request.Params("age"))

	response := map[string]interface{}{
		"info":   msg,
		"active": true,
	}
	// return c.Status(200).JSON(fiber.Map{}) // => info: 👴 john is 75 years old
	return presentationprotocols.HTTPResponse{
		StatusCode: 200,
		Data:       response,
	}
}
