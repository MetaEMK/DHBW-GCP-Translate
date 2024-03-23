function addListener() {
    document.body.addEventListener('htmx:beforeSwap', function(evt) {
        document.body.addEventListener('htmx:responseError', function (evt) {
            let errorMessage = evt.detail.xhr.responseText || 'Unbekannter Fehler';
            document.getElementById('error-message').textContent = 'Ein Fehler ist aufgetreten: ' + errorMessage;
            document.getElementById('error-message').style.display = 'block';
        });
    });

    document.body.addEventListener('htmx:configRequest', function(evt) {
        document.getElementById('error-message').style.display = 'none';
    });
}

htmx.onLoad(function(target) {
    console.log("Hello World");
    addListener();
});
