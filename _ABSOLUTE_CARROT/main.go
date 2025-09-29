package main

import (
	_ "embed"
	_ "fmt"
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

func feed_carrot(this js.Value, args []js.Value) any {
	db := this.Get("db")
	food := args[0].String()
	FeedMessageToCarrotson(db, food)

	// JSON := js.Global().Get("JSON")
	// stmt := db.Call("prepare", "SELECT * FROM Carrotson_Branches")
	// for stmt.Call("step").Bool() {
	// 	row := stmt.Call("getAsObject");
	// 	fmt.Println("Here is a row: " + JSON.Call("stringify", row).String())
	// }
	return nil
}

func carrot_generate(this js.Value, args []js.Value) any {
	message, err := CarrotsonGenerate(this.Get("db"), "", 256)
	if err != nil {
		log.Printf("%s\n", err)
		return nil
	}
	return message 
}

func main() {
	window := js.Global().Get("window")
	db := js.Global().Get("db")
	db.Call("run", init_sql)

	window.Set("carrot_generate", js.FuncOf(carrot_generate))
	window.Set("feed_carrot", js.FuncOf(feed_carrot))
	
	FeedMessageToCarrotson(db, "HELLO")
	FeedMessageToCarrotson(db, "HELP")
	FeedMessageToCarrotson(db, "HELL")
	FeedMessageToCarrotson(db, "HELLO KITTY")
	FeedMessageToCarrotson(db, "HELLO WORLD")
	FeedMessageToCarrotson(db, "MOM'S SPAGETI")

	select {}
}
