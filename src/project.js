export default class {
    constructor(list = []) {
        this.element = document.createElement('div')
        this.hide()
        list.forEach(item => this.new_child(item))
    }
    show() {
        this.element.style.display = 'block'
    }
    hide() {
        this.element.style.display = 'none'
    }
    new_child(item) {
        let div = document.createElement('div')
        div.innerHTML = item
        this.element.appendChild(div)
    }
}