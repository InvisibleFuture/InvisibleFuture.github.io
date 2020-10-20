/** 
 * 输入:
 * 标签名, 显示回调, 隐藏回调
 * 
 * 输出:
 * element
 * 
 * DEMO:
 * document.body.appendChild(new Tabs([
 *     { name: 'Idea', show: idea.show.bind(idea), hide: idea.hide.bind(idea) },
 *     { name: 'Project', show: project.show.bind(project), hide: project.hide.bind(project) },
 * ], 1).element)
 * 
**/
export default class {
    constructor(list = [], active = 0) {
        this.style('tabs', `
            .tabs {
                display: flex;
                align-items: flex-end;
                justify-content: flex-end;
                user-select:none;
            }
            .tabs div {
                cursor: pointer;
                padding: 0 .8em;
                height: 2em;
                line-height: 2em;
                border-radius: .25em;
                transition: all .75s;
            }
            .tabs div:hover {
                color: #32c787;
            }
            .tabs .active {
                color: #32c787;
                font-size: 1.5em;
                font-weight: 600;
            }
        `)
        // 主元素赋予 classname
        this.element = document.createElement('div')
        this.element.className = 'tabs'


        // 外部调用选中目标
        this.active = active
        // 追加子元素
        list.forEach((item, i) => {
            let div = document.createElement('div')
            div.innerHTML = item.name
            div.onclick = () => {
                item.show()
                div.classList.add('active')
                if (this.active_element && this.active_item && this.active_element != div) {
                    this.active_element.classList.remove('active')
                    this.active_item.hide()
                }
                this.active_item = item
                this.active_element = div
            }
            if (i === this.active) {
                this.active_item = item
                this.active_element = div
                this.active_element.classList.add('active')
                item.show()
            }
            this.element.appendChild(div)
        })
    }
    // 向 <head> 写入 <style> 并防止重复
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