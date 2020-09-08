export default class {
    constructor(data) {
        this.element = document.createElement("div")
        data.forEach(item => {
            let el = new sentence(item)
            this.element.appendChild(el.element)
        })
    }
}

class sentence {
    constructor(data) {
        this.element = document.createElement("div")
        this.element.style.cssText = `
            width: 100%;
            max-width: 600px;
            margin: 16px auto;
            padding: 16px;
            border: thin solid rgba(0,0,0,.12);
        `
        let header = document.createElement("div")
        let content = document.createElement("div")
        this.element.appendChild(header)
        this.element.appendChild(content)

        content.innerHTML = data

        let time = document.createElement("span")
        time.innerHTML = data.time
    }
    Render() {}
}