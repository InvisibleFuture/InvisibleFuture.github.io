export default class {
    constructor(list = []) {
        this.style('style_project', `
            .project {
                display: none;
                transition: all .75s;
            }
            .project.show {
                display: block;
                max-width: 800px;
                margin: 1em auto;
                padding: 1em;
            }
            .project>div {
                margin:2em 1em;
            }
            /** markdown css **/
            .project p {
                margin: 2em 1em;
            }
        `)
        this.element = document.createElement('div')
        this.element.className = 'project'
        this.hide()
        list.forEach(item => this.new_child(item))
    }
    show() {
        this.element.classList.add('show')
    }
    hide() {
        this.element.classList.remove('show')
    }
    new_child(item) {
        let div = document.createElement('div')
        div.innerHTML = item
        this.element.appendChild(div)
    }
    style(id, content) {
        let styles = document.head.getElementsByTagName("style")
        for (let i = styles.length; i--; i > 0) {
            if (styles[i].id === id) return
        }
        let style = document.createElement('style')
        style.innerHTML = content
        style.id = id
        document.head.appendChild(style)
    }
}