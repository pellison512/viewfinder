 let changeColor = document.getElementById('changeColor');

  chrome.storage.sync.get('color', function(data) {
    changeColor.style.backgroundColor = data.color;
    changeColor.setAttribute('value', data.color);
  });

  changeColor.onclick = async function(element) {
    var bkg = chrome.extension.getBackgroundPage();
    var pageTitle = "";
    chrome.tabs.query({active: true, currentWindow: true}, async function(tabs){
      pageTitle = tabs[0].title;
      bkg.console.log('pageTitle:  '+ pageTitle);
      getTabData(pageTitle);
    });
    bkg.console.log('changing color');
    bkg.console.log('page title: \''+ pageTitle+'\'');
    let color = element.target.value;
    chrome.tabs.query({active: true, currentWindow: true}, function(tabs) {
      chrome.tabs.executeScript(
          tabs[0].id,
          {code: 'document.body.style.backgroundColor = "' + color + '";'});
    });
    

  };

  function getTabData(tabTitle) {
    var bkg = chrome.extension.getBackgroundPage();

    fetch('http://www.localhost:8090/windows/'+tabTitle).then(r => r.text()).then(result => {
      // Result now contains the response text, do what you want...
      bkg.console.log(result);
    });
  }