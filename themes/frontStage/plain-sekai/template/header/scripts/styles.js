$(".sekai-input-warp").hover(function handlerIn() {
    let label = $(this).children(".sekai-input-label");
    if (label.length < 0) {
        return false;
    } else {
        if (label.attr("focus") === "false"){
            label.css("color", "rgba(126,126,126,1)");
        }
    }
}, function handlerOut() {
    let label = $(this).children(".sekai-input-label");
    if (label.length < 0) {
        console.log(label.attr("focus") === "true");
        return false;
    } else {
        if (label.attr("focus") === "false") {
            label.css("color", "rgba(171,171,171,1)");
        }
    }
});
$(".sekai-input").focus(function () {
    let label = $(this).parent().children(".sekai-input-label");
    console.log(label.content);
    if (label.length < 0) {
        return false;
    } else {
        label.css("color", "rgba(  7,187,195,1)");
        label.attr("focus", true);
    }
});
$(".sekai-input").focusout(function () {
    let label = $(this).parent().children(".sekai-input-label");
    if (label.length < 0) {
        return false;
    } else {
        label.css("color", "rgba(171,171,171,1)");
        label.attr("focus", false);
    }
});
$.onload(function () {
    $("html").addClass("noscroll");
});