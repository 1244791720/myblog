// 协议
const protocol = "http://"
// 主机ip
const ip = "localhost"
// 端口
const port = "8080"
// 域名
const host  = protocol + ip  + ":" + port + "/"

// 获取全部文章
const getAllArticles = host + "articles"
// 获取分页文章
const getPageArticles = host + "pageArticles"
// 获取单个文章
const getOneArticle = host + "article"
// 获取文章总数
const countArticles = host + "countArticles"
// 修改单个文章内容 参数放在请求体里
const modifyArticleContent = host + "article"
// 修改文章内容表单
const modifyArticleForm = host + "articleForm"
// 删除一篇文章
const deleteArticleApi = host + "deleteOneArticle"
// 登录
const loginApi = host + "login"
// 阅读数加1
const viewNumIncreaseApi = host + "viewNum"
// 获取所有文章类型
const articleTypesApi = host + "articleType"
// 从地址栏取参数
// getQueryString("参数名1")
function getQueryString(name){
    var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if(r!=null)return  unescape(r[2]); return null;
}