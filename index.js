const go = new Go();
wasm_promise = WebAssembly.instantiateStreaming(fetch("/main.wasm"), go.importObject)


// const db_path = "/test.sqlite"
const db_path = "/bee.sqlite"
const dataPromise = fetch(db_path).then(res => res.arrayBuffer());
window.db_promise = initSqlJs({ locateFile: file => `/${file}` }).then(function(SQL){
	dataPromise.then(buf => {
		window.db = new SQL.Database(new Uint8Array(buf));
		console.log("db set");
		db.create_function("starts_with", (a, b) => a.startsWith(b));
		// db.create_function("random", () => { 
		// 	return Math.floor(Math.random() * (9223372036854775807 - -9223372036854775807 + 1) + -9223372036854775807)
		// });
		wasm_promise.then((result) => {
			console.log("wasm promise: go run started");
    		go.run(result.instance);
			console.log(window.location)
			food = decodeURI(window.location.pathname.slice(1,))
			window.feed_carrot(food)

			var new_message = window.carrot_generate();
			message.innerHTML = new_message;
		});
		console.log("sql-js inited");
	});
})



window.onload = function () {
	console.log("ONLOAD");

	today = new Date();
	var cday = new Date(today.getFullYear(), 10, 9);
	if (today.getMonth() == 10 && today.getDate() > 9) {
	    cday.setFullYear(cday.getFullYear() + 1);
	}  
	var one_day = 1000 * 60 * 60 * 24;
	subtitle.innerHTML += Math.ceil((cday.getTime() - today.getTime()) / (one_day)) + " days left"

	refresh.onclick = e => {
		var new_message = window.carrot_generate();
		message.innerHTML = new_message;
	}
}
document.addEventListener("DOMContentLoaded", function(event) {
	console.log("DOMContentLoaded");
	message.innerHTML = "...";
});


window.onkeydown = e => {
    if (e.target.type == "text") return;
    switch (e.code) {
        case "KeyR":
            if (e.ctrlKey === false) {
				var new_message = window.carrot_generate();
				message.innerHTML = new_message;
			} break
    }
}

