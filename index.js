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
			response.innerHTML = new_message;
		});
		console.log("sql-js inited");
	});
})


function send_message() {
	var new_message = window.carrot_generate(message.value);
	response.innerHTML = new_message;
	if (!ignore_toggle.classList.contains('enabled')) {
		window.feed_carrot(message.value);
		message.value = "";
	}
}

function toggle_ignore()
{
	if (!ignore_toggle.classList.contains('enabled')) {
		ignore_toggle.classList.add("enabled");
	}
	else {
		ignore_toggle.classList.remove("enabled");
	}
	if (ignore_toggle.classList.contains('enabled')) {
		ignore_toggle.title = "Not ignore: carrot will eat your message; (I)"
	}
	else {
		ignore_toggle.title = "Ignore: carrot will not eat your message; (I)"
	}
}

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
		response.innerHTML = new_message;
	}

	message.onkeydown = e => {
		if(e.keyCode == 13){
			send_message();
		}
	}

	ignore_toggle.onclick = e => {
		toggle_ignore();	
	}

	send.onclick = send_message;
	toggle_ignore();	
}
document.addEventListener("DOMContentLoaded", function(event) {
	console.log("DOMContentLoaded");
	response.innerHTML = "...";
});


window.onkeydown = e => {
    if (e.target.type == "text") return;
    switch (e.code) {
        case "KeyR":
            if (e.ctrlKey === false) {
				var new_message = window.carrot_generate();
				response.innerHTML = new_message;
			}
			break
		case "KeyI":
			toggle_ignore();	
			break
    }
}

