$(document).ready(function () {
    $.get("httP://localhost:8000/dig",
        {
            "time":gettime(),
            "url":geturl(),
            "refer":getrefer(),
            "ua":getuser_agent(),
        });
})
/*
gettime(); //js获取当前时间
getip(); //js获取客户端ip
geturl(); //js获取客户端当前url
getrefer(); //js获取客户端当前页面的上级页面的url
getuser_agent(); //js获取客户端类型
getcookie() //js获取客户端cookie
*/

function gettime(){
    var nowDate = new Date();
    return nowDate.toLocaleString();
}
function geturl(){
    return window.location.href;
}
function getrefer(){
    return document.referrer;
}
function getcookie(){
    return document.cookie;
}
function getuser_agent(){
    return navigator.userAgent;
}