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


document.body.appendChild(ListView({
	appendChild: [
		Note({text: '奇怪的物, 落在生中的东西'}),
		Note({text: '在对代码渲染中, 代码是作为一个文件片段/对象存在的,它应有单独的操纵方法'}),
		Note({text: '任何对象都应具备可操作的选项, 而非仅仅展示. 明确展现目的和下一步会做的事情, 让事物变得可控 对一句话的回应/删除/涂改/赞许, 对一个角色对象的关注/屏蔽'}),
		Note({text: '清除不必要的对象, 在任何为重点/集中展现的场景. 目的对象本身以外对象主体不应超过2个, 否则会对部分思考能力造成压力. 并且, 也是为了在计算原理上更快命中目标.'}),
		Note({text: '一种界面交互体验, 对象不是附属于视域的在展现中始终暗示对象的可预测大小, 在无限大且无边界的对象展现中, 应始终充满/超出视域'}),
	]
}))

function Note(state) {
	let note = document.createElement('div');
	note.style.cssText = state.style;
	note.innerText = state.text;
	return note;
}

function ListView(state) {
	let listview = document.createElement('div');
	if(state.style) listview.style.cssText = state.style;
	if(state.class) listview.className = state.class;
	state.children && state.children.forEach(item => {
		listview.appendChild(item);
	})
}