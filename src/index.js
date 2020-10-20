import Tabs from "./tabs.js"
import Idea from './idea.js'
import Project from './project.js'

//let idea = new Idea([
//    [
//        Text("TEXT DEMO"),
//        Image("https://satori.love/"),
//        Date("2020年10月20日"),
//    ],
//    [
//        Text(`一种界面交互体验,
//            对象不是附属于视域的在展现中始终暗示对象的可预测大小,
//            在无限大且无边界的对象展现中, 应始终充满/超出视域
//        `),
//        Date("2020年10月20日"),
//    ],
//    [
//        Text("清除不必要的对象, 在任何为重点/集中展现的场景. 目的对象本身以外对象主体不应超过2个, 薄弱的逻辑思考能力会对此倍感压力, 以及非重点的心理暗示. 并且也是为在计算原理上更快命中目标"),
//    ]
//])

let idea = new Project([])
task.element.classList.add('markdown_idea')

let project = new Project([
    "任何对象都应具备可操作的选项, 而非仅仅展示. 明确展现目的和下一步会做的事情, 让事物变得可控 对一句话的回应/删除/涂改/赞许, 对一个角色对象的关注/屏蔽",
    "在对代码渲染中, 代码是作为一个文件片段/对象存在的,它应有单独的操纵方法",
])

let task = new Project([])
task.element.classList.add('markdown_task')

let store = new Project([
    "妹は、世界で最もかわいいです",
    "奇怪的物, 落在生中的东西"
])

document.body.appendChild(new Tabs([
    { name: 'Idea', show: idea.show.bind(idea), hide: idea.hide.bind(idea) },
    { name: 'Project', show: project.show.bind(project), hide: project.hide.bind(project) },
    { name: 'Task', show: task.show.bind(task), hide: task.hide.bind(task) },
    { name: 'Store', show: store.show.bind(store), hide: store.hide.bind(store) },
], 0).element)

document.body.appendChild(idea.element)
document.body.appendChild(project.element)
document.body.appendChild(task.element)
document.body.appendChild(store.element)

function Text(data) {
    let element = document.createElement('div')
    element.innerHTML = data
    return element
}

function Date(data) {
    let element = document.createElement('div')
    element.innerHTML = data
    return element
}

function Image(data) {
    let element = document.createElement('div')
    element.innerHTML = data
    return element
}