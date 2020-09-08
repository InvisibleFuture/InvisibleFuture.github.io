export default class {
    constructor(data) {
        this.element = document.createElement("div")
        data.forEach(item => {
            let el = new image(item)
            this.element.appendChild(el.element)
        })
    }
}

class image {
    constructor(data) {
        this.element = document.createElement("img")
        this.element.style.cssText = `
            width: 480px;
            margin: 16px;
            padding: 16px;
            border: thin solid rgba(0,0,0,.12);
        `
        this.element.src = data
    }
}