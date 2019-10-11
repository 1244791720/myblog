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

// 从地址栏取参数
// getQueryString("参数名1")
function getQueryString(name){
    var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if(r!=null)return  unescape(r[2]); return null;
}