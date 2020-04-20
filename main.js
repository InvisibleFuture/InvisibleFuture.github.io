console.log("wo haha~")

fetch('/main.json', {
	headers: {
		Accept: 'application/json',
		'Content-Type': 'application/json; charset=utf-8',
		Authorization: 'SAPISIDHASH 012ao5airhvdu22esdahk3pn74'
	}
})
.then(response => { return response.json();})
.then(data => { window.console.log(data); })
.catch(e => {console.log(e)})