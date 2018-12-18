// Diese Browser Extension beauftragt content.js, die URL einer fremden Seite zu ermitteln.
// background.js und content.js kommunizieren mit messages:
//
document.addEventListener('DOMContentLoaded', function () {

    // Handler für browser-action button:
    chrome.browserAction.onClicked.addListener(function (tab) {

        // Send a message to the active tab:
        chrome.tabs.query({
            active: true,
            currentWindow: true
        }, function ([activeTab]) {
            chrome.tabs.query({}, function (tabs) {
                const host = 'http://localhost:3000'
                let headers = new Headers();
                let username = 'barne';
                let password = 'barne';
                let data = {
                    url: activeTab.url
                }
                headers.append('Authorization', 'Basic ' + btoa(username + ":" + password));
                fetch(host + "/bookmark", {
                        method: 'POST',
                        headers,
                        body: JSON.stringify(data)
                    })
                    .then(() => {
                        if (tabs.filter(item => item.url.includes(host)).length == 0)
                            chrome.tabs.create({
                                url: host + '/app'
                            })
                    })
            })
        });

    });
});