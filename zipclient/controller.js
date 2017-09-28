document.getElementsByName("username")[0].addEventListener('change', handleNameChange);

function handleNameChange() {
    var value = this.value;
    fetch('http://localhost:4000/hello?name=' + value).then(function(response) {
        response.text().then(function(object) {
            console.log(object);
            document.getElementById('display').innerHTML = object;
        });

    });
}

var timer = setInterval(handleMemoryCheck, 1000);

function handleMemoryCheck() {
    fetch('http://localhost:4000/memory').then(function(response) {
        response.json().then(function(object) {
            document.getElementById('memory').innerHTML = 'Memory Alloc: ' + object['Alloc'];
        });

    });
}