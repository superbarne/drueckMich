// Dieser Code wird im Kontext einer fremden Seite (aktiver Tab) ausgeführt.
// Mit der eigenen Anwendung (background.js) kann nur über messages kommuniziert werden:
//
chrome.runtime.onMessage.addListener(
    function (request, sender, sendResponse) {
        if (request.message === "backgroundAnContent_hatDraufGedrueckt") {

            var hrefActiveTab = window.location.href;

            chrome.runtime.sendMessage({
                "message": "contentAnBackground_hierDattHref",
                "href": hrefActiveTab
            });
        }
    }
);