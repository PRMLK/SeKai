//     original $rgba-grey-2 rgba(227,227,227,1)
//     hover $rgba-grey-3 rgba(126,126,126,1)
//     focus $rgba-info-0 rgba(  7,187,195,1)

// TODO: debug cookie funtions
function setCookie(key, value, expires, path) {
    let oDate = new Date();
    oDate.setDate(oDate.getDate() + expires);
    document.cookie = key + '=' + value + '; expires=' + oDate + '; path=' + path;
    console.log(key + '=' + value + '; expires=' + oDate + '; path=' + path)
}

function removeCookie(key) {
    setCookie(key, '', -1);
}

function getCookie(key) {
    var cookieArr = document.cookie.split('; ');
    for (var i = 0; i < cookieArr.length; i++) {
        var arr = cookieArr[i].split('=');
        if (arr[0] === key) {
            return arr[1];
        }
    }
    return false;
}