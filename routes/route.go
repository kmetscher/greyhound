package routes

import "greyhound/controllers"

type Method string
const (
    Get Method = "GET"
    Head Method = "HEAD" 
    Post Method = "POST"
    Put Method = "PUT"
    Patch Method = "PATCH"
    Delete Method = "DELETE"
)

type Route struct {
    Method Method
    Path string
    Controller controllers.Controller
}
