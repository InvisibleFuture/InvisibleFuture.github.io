import Tabs from "./tabs.js"
import Project from './project.js'

let idea = new Project([])
idea.element.classList.add('markdown_idea')

let project = new Project([])
project.element.classList.add('markdown_project')

let task = new Project([])
task.element.classList.add('markdown_task')

let store = new Project([])
store.element.classList.add('markdown_store')

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

// function Text(data) {
//     let element = document.createElement('div')
//     element.innerHTML = data
//     return element
// }
// 
// function Date(data) {
//     let element = document.createElement('div')
//     element.innerHTML = data
//     return element
// }
// 
// function Image(data) {
//     let element = document.createElement('div')
//     element.innerHTML = data
//     return element
// }