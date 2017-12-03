package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route {
        "UploadSubDevice",
        "PUT",
        "/subscriptions/{username}/{deviceid}.json",
        UploadSubDevice,
    },
    Route {
        "SetDevices",
        "POST",
        "/api/2/devices/{username}/{deviceid}.json",
        YES,
    },
    Route {
        "GetSubscriptsion",
        "GET",
        "/api/2/subscriptions/{username}/{deviceid}.json",
        YES,
    },
    Route {
        "GetEpisodes",
        "GET",
        "/api/2/episodes/{username}.json",
        YES,
    },
    Route {
        "SetEpisodes",
        "POST",
        "/api/2/episodes/{username}.json",
        YES,
    },
    Route {
        "SetEpisodes",
        "POST",
        "/api/2/episodes/{username}.json?since={since:[0-9]}",
        YES,
    },
}
