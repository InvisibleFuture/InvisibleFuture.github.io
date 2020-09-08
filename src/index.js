import sentence from "./sentence.js"
import images from "./images.js"

let se = new sentence([
    "一种界面交互体验, 对象不是附属于视域的在展现中始终暗示对象的可预测大小, 在无限大且无边界的对象展现中, 应始终充满/超出视域",
    "清除不必要的对象, 在任何为重点/集中展现的场景. 目的对象本身以外对象主体不应超过2个, 薄弱的逻辑思考能力会对此倍感压力, 以及非重点的心理暗示. 并且也是为在计算原理上更快命中目标",
    "任何对象都应具备可操作的选项, 而非仅仅展示. 明确展现目的和下一步会做的事情, 让事物变得可控 对一句话的回应/删除/涂改/赞许, 对一个角色对象的关注/屏蔽",
    "在对代码渲染中, 代码是作为一个文件片段/对象存在的,它应有单独的操纵方法",
    "奇怪的物, 落在生中的东西"
])


let im = new images([
    "https://www.satori.love/satori.jpeg"
])

document.body.appendChild(se.element)
document.body.appendChild(im.element)