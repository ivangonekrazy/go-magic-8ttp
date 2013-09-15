package main

import (
    "fmt"
    "log"
    "net/http"
    "math/rand"
)

var httpStatuses = map[int]string {
    200: "OK",
    202: "Accepted",
    203: "Non-authoritative information",
    300: "Multiple choices",
    302: "Found",
    401: "Unauthorized",
    403: "Forbidden",
    408: "Request timeout",
    417: "Expectation failed",
    500: "Internal server error",
    501: "Not implemented",
    502: "Bad gateway",
    503: "Service unavailable",
}

type MagicResp struct {
    code    int
    fortune string
}

var responsesArr = []MagicResp {
    {200, "It is certain."},
    {200, "It is decidedly so."},
    {200, "Without a doubt."},
    {200, "Yes, definitely."},
    {200, "You may rely on it."},
    {203, "As I see it, yes."},
    {203, "Most likely."},
    {202, "Outlook good."},
    {200, "Yes."},
    {302, "Signs point to yes."},
    {300, "Reply hazy, try again."},
    {408, "Ask again later."},
    {401, "Better not tell you now."},
    {503, "Cannot predict now."},
    {408, "Concentrate and ask again."},
    {417, "Don't count on it."},
    {403, "My reply is no."},
    {502, "My sources say no."},
    {501, "Outlook not so good."},
    {500, "Very doubtful."},
}

func handleMagic8ttp(w http.ResponseWriter, r *http.Request) {
    var rnd = rand.Intn(len(responsesArr))
    var magic = responsesArr[rnd]
    var respString = fmt.Sprintf("\"%s\" (%d - %s)\n", magic.fortune, magic.code, httpStatuses[magic.code])

    log.Println(r.RemoteAddr + " " + respString)

    w.WriteHeader(magic.code)
    fmt.Fprintf(w, respString)
}

func main() {
    log.Println("Starting Magic 8ttp Ball service...")

    http.HandleFunc("/", handleMagic8ttp)
    http.ListenAndServe(":8080", nil)
}
