export default class {
    constructor(list = []) {
        //this.style('ideax', ``)

        this.element = document.createElement('div')
        this.element.className = 'idea'
        this.hide()
        list.forEach(item => this.new_child(item))
        this.xstyle('idea', `
            .idea {
                display: none;
                transition: all .75s;
            }
            .idea.show {
                display: block;
                max-width: 800px;
                margin: 1em auto;
                padding: 1em;
            }
            .idea>div {
                margin:2em 1em;
            }
        `)
    }
    show() {
        this.element.classList.add('show')
    }
    hide() {
        this.element.classList.remove('show')
        // 折叠元素
        // 压缩元素
        // 模糊元素
        // 隐藏元素
        // 聚点元素
    }
    new_child(list) {
        let div = document.createElement('div')
        list.forEach(item => div.appendChild(item))
        this.element.appendChild(div)
    }
    // 向 <head> 写入 <style> 并防止重复
    xstyle(id, content) {
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