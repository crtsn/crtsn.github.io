package main

// GOOS=js GOARCH=wasm go build -o main.wasm main.go

import (
	_ "embed"
	"fmt"
	"net/url"
	"strings"
	"log"
	"regexp"
	"syscall/js"
)

var DiscordPingRegexp = regexp.MustCompile("<@[0-9]+>")

func maskDiscordPings(message string) string {
	return DiscordPingRegexp.ReplaceAllString(message, "@[DISCORD PING REDACTED]")
}

var (
	// TODO: make the CommandPrefix configurable from the database, so we can set it per instance
	CommandPrefix = "[\\$\\!]"
	CommandDef    = "([a-zA-Z0-9\\-_]+)( +(.*))?"
	CommandRegexp = regexp.MustCompile("^ *(" + CommandPrefix + ") *" + CommandDef + "$")
)

type Command struct {
	Prefix string
	Name   string
	Args   string
}

func parseCommand(source string) (Command, bool) {
	matches := CommandRegexp.FindStringSubmatch(source)
	if len(matches) == 0 {
		return Command{}, false
	}
	return Command{
		Prefix: matches[1],
		Name:   matches[2],
		Args:   matches[4],
	}, true
}

//go:embed 02-carrotson.sql
var init_sql string

func main() {
	window := js.Global().Get("window")
	db := js.Global().Get("db")
	db.Call("run", init_sql)
	
	location := window.Get("location").Get("href").String()
	if strings.Contains(location, "?") {
		u, err := url.Parse(location)
		if err != nil {
		    panic(err)
		}
		query, err := url.QueryUnescape(u.RawQuery)
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(query, "/") {
			query = query[1:]
		}
		FeedMessageToCarrotson(db, query)
	}

	FeedMessageToCarrotson(db, "HELLO")
	FeedMessageToCarrotson(db, "HELP")
	FeedMessageToCarrotson(db, "HELL")
	FeedMessageToCarrotson(db, "HELLO KITTY")
	FeedMessageToCarrotson(db, "HELLO WORLD")

	// JSON := js.Global().Get("JSON")
	// Stmt := db.Call("prepare", "SELECT * FROM Carrotson_Branches")
	// For stmt.Call("step").Bool() {
	// 	row := stmt.Call("getAsObject");
	// 	fmt.Println("Here is a row: " + JSON.Call("stringify", row).String())
	// }

	// message, err := CarrotsonGenerate(db, "HEL", 256)
	message, err := CarrotsonGenerate(db, "", 256)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
	fmt.Println(maskDiscordPings(message))

	document := js.Global().Get("document")
	par := document.Call("createElement", "p")
	par.Set("innerHTML", "CARROT SAYS: " + maskDiscordPings(message))
	document.Call("getElementsByTagName", "body").Index(0).Call("appendChild", par)
	
	select {}
}
