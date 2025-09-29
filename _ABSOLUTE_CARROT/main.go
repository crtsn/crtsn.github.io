package main

import (
	"embed"
	_ "fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
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

func await(awaitable js.Value) ([]js.Value, []js.Value) {
    then := make(chan []js.Value)
    defer close(then)
    thenFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        then <- args
        return nil
    })
    defer thenFunc.Release()

    catch := make(chan []js.Value)
    defer close(catch)
    catchFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        catch <- args
        return nil
    })
    defer catchFunc.Release()

    awaitable.Call("then", thenFunc).Call("catch", catchFunc)

    select {
    case result := <-then:
        return result, nil
    case err := <-catch:
        return nil, err
    }
}

//go:embed images/*
var images embed.FS

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
	FeedMessageToCarrotson(db, "MOM'S SPAGETI")

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

	document := js.Global().Get("document")

	carrot_img, err := images.ReadFile("images/carrot.svg")
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	BlobConstructor := js.Global().Get("Blob")
	blob := BlobConstructor.New([]any{string(carrot_img)}, map[string]any{"type": "image/svg+xml"})
	URL := js.Global().Get("URL")
	url := URL.Call("createObjectURL", blob)
	carrot_svg := document.Call("createElement", "img")
	carrot_svg.Set("src", url)
	carrot_svg.Call("addEventListener", "load",
		js.FuncOf(func(this js.Value, args []js.Value) any {
			URL.Call("revokeObjectURL", url)
			return nil
		}),
		map[string]any{"once": true})
	decode_promise := carrot_svg.Call("decode")
	await(decode_promise)

	happy_mouth_img, err := images.ReadFile("images/happy_mouth.png")
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
	size := len(happy_mouth_img)
	mouth_bytes := js.Global().Get("Uint8Array").New(size)
	js.CopyBytesToJS(mouth_bytes, happy_mouth_img)
	blob = BlobConstructor.New([]any{mouth_bytes}, map[string]any{"type": "image/png"})
	url = URL.Call("createObjectURL", blob)
	happy_mouth_png := document.Call("createElement", "img")
	happy_mouth_png.Set("src", url)
	happy_mouth_png.Call("addEventListener", "load",
		js.FuncOf(func(this js.Value, args []js.Value) any {
			URL.Call("revokeObjectURL", url)
			return nil
		}),
		map[string]any{"once": true})
	decode_promise = happy_mouth_png.Call("decode")
	await(decode_promise)

	canvas := document.Call("querySelector", "#main")
	ctx := canvas.Call("getContext", "2d", map[string]any{"alpha": false})

	const width = 600
	const height = 400
	canvas.Set("width", width)
	canvas.Set("height", height)

	ctx.Set("fillStyle", "#00c3ff")
	ctx.Call("fillRect", 0, 0, width, height*0.7)
	ctx.Set("fillStyle", "#9d582e")
	ctx.Call("fillRect", 0, height*0.7, width, height)

	ctx.Call("drawImage", carrot_svg, 200, 50)
	ctx.Call("drawImage", happy_mouth_png, 340, 145, 345/2.8, 345/2.8)

	rectX := 30
	rectY := 30
	rectWidth := 250
	rectHeight := 170
	ctx.Set("fillStyle", "#ececbc")
	ctx.Call("beginPath")
	ctx.Call("roundRect", rectX, rectY, rectWidth, rectHeight, 40)
	ctx.Call("closePath")
	ctx.Call("fill")
	ctx.Call("beginPath")
	ctx.Call("moveTo", 270, 180)
	ctx.Call("lineTo", 330, 190)
	ctx.Call("lineTo", 270, 150)
	ctx.Call("closePath")
	ctx.Call("fill")

	ctx.Set("fillStyle", "#000")
	ctx.Set("font", "24px sans")
	ctx.Set("textAlign", "center")
	ctx.Set("textBaseline", "middle")
	ctx.Call("fillText", maskDiscordPings(message), rectX+(rectWidth/2),rectY+(rectHeight/2))

	select {}
}
