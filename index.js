const go = new Go();
// wasm_promise = WebAssembly.instantiateStreaming(fetch("/main.wasm"), go.importObject)

// const db_path = "/test.sqlite"
const db_path = "/bee.sqlite"
// const dataPromise = fetch(db_path).then(res => res.arrayBuffer());
// window.db_promise = initSqlJs({ locateFile: file => `/${file}` }).then(function(SQL){
// 	dataPromise.then(buf => {
// 		window.db = new SQL.Database(new Uint8Array(buf));
// 		console.log("db set");
// 		db.create_function("starts_with", (a, b) => a.startsWith(b));
// 		// db.create_function("random", () => { 
// 		// 	return Math.floor(Math.random() * (9223372036854775807 - -9223372036854775807 + 1) + -9223372036854775807)
// 		// });
// 		wasm_promise.then((result) => {
// 			console.log("wasm promise: go run started");
//     		go.run(result.instance);
// 			console.log(window.location)
// 			food = decodeURI(window.location.pathname.slice(1,))
// 			window.feed_carrot(food)
// 
// 			canvas = document.querySelector('#main');
// 			ctx = canvas.getContext('2d', {alpha: false})
// 			
// 			rectX = 30
// 			rectY = 30
// 			rectWidth = 250
// 			rectHeight = 170
// 			
// 			ctx.fillStyle = "#000"
// 			ctx.font = "24px sans"
// 			ctx.textAlign = "center"
// 			var lines = fragmentText(window.carrot_generate(), rectWidth - parseInt(ctx.font,0));
// 		    lines.forEach(function(line, i) {
// 		        ctx.fillText(line, rectX + rectWidth / 2, rectY + (rectHeight - parseInt(ctx.font,0) * lines.length) / 2 + (i + 1) * parseInt(ctx.font,0));
// 		    });
// 		});
// 		console.log("sql-js inited");
// 	});
// })


function fragmentText(text, maxWidth) {
    var words = text.split(' '),
        lines = [],
        line = "";
    if (ctx.measureText(text).width < maxWidth) {
        return [text];
    }
    while (words.length > 0) {
        while (ctx.measureText(words[0]).width >= maxWidth) {
            var tmp = words[0];
            words[0] = tmp.slice(0, -1);
            if (words.length > 1) {
                words[1] = tmp.slice(-1) + words[1];
            } else {
                words.push(tmp.slice(-1));
            }
        }
        if (ctx.measureText(line + words[0]).width < maxWidth) {
            line += words.shift() + " ";
        } else {
            lines.push(line);
            line = "";
        }
        if (words.length === 0) {
            lines.push(line);
        }
    }
    return lines;
}
window.onload = function () {
	console.log("ONLOAD");
	// canvas = document.querySelector('#main');
	// ctx = canvas.getContext('2d', {alpha: false})

	// const width = 600
	// const height = 400

	// ctx.fillStyle = '#00c3ff'
	// ctx.fillRect(0, 0, width, height*0.7)
	// ctx.fillStyle = '#9d582e'
	// ctx.fillRect(0, height*0.7, width, height)
	// 
	// ctx.drawImage(carrot_svg, 200, 50)
	// ctx.drawImage(happy_mouth, 340, 145, 345/2.8, 345/2.8)

	// rectX = 30
	// rectY = 30
	// rectWidth = 250
	// rectHeight = 170
	// ctx.fillStyle = '#ececbc'
	// ctx.beginPath()
	// ctx.roundRect(rectX, rectY, rectWidth, rectHeight, 40)
	// ctx.closePath()
	// ctx.fill()

	// ctx.beginPath()
	// ctx.moveTo(270, 180)
	// ctx.lineTo(330, 190)
	// ctx.lineTo(270, 150)
	// ctx.closePath()
	// ctx.fill()

	today = new Date();
	var cday = new Date(today.getFullYear(), 10, 9);
	if (today.getMonth() == 10 && today.getDate() > 9) {
	    cday.setFullYear(cday.getFullYear() + 1);
	}  
	var one_day = 1000 * 60 * 60 * 24;
	subtitle.innerHTML += Math.ceil((cday.getTime() - today.getTime()) / (one_day)) + " days left"
}

window.onkeydown = e => {
    if (e.target.type == "text") return;
    switch (e.code) {
        case "KeyR":
            if (e.ctrlKey === false) {
				rectX = 30
				rectY = 30
				rectWidth = 250
				rectHeight = 170
				ctx.fillStyle = '#ececbc'
				ctx.beginPath()
				ctx.roundRect(rectX, rectY, rectWidth, rectHeight, 40)
				ctx.closePath()
				ctx.fill()

				ctx.beginPath()
				ctx.moveTo(270, 180)
				ctx.lineTo(330, 190)
				ctx.lineTo(270, 150)
				ctx.closePath()
				ctx.fill()

				ctx.fillStyle = "#000"
				ctx.font = "24px sans"
				ctx.textAlign = "center"
				var lines = fragmentText(window.carrot_generate(), rectWidth - parseInt(ctx.font,0));
		    	lines.forEach(function(line, i) {
		    	    ctx.fillText(line, rectX + rectWidth / 2, rectY + (rectHeight - parseInt(ctx.font,0) * lines.length) / 2 + (i + 1) * parseInt(ctx.font,0));
		    	});
            } break
    }
}
