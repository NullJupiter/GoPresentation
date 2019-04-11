let data = null;
let titleDom = null;
let bulletPointsDom = null;
let currentFoil = null

document.addEventListener('keyup', (e) => {
    if (e.keyCode == 39) {
        nextFoil();
    } else if (e.keyCode == 37) {
        previousFoil();
    }
});

window.onload = () => {
    currentFoil = 0;
    data = JSON.parse(document.getElementById('data').innerHTML);
    titleDom = document.getElementById('title');
    bulletPointsDom = document.getElementById('bullet-points');

    // Load first foil
    updateContent()
}

function updateContent() {
    document.title = data[currentFoil].title;
    titleDom.innerHTML = data[currentFoil].title;
    let bulletPointsOutput = '';
    for (b of data[currentFoil].bullet_points) {
        bulletPointsOutput += '<li>➜ ' + b + '</li>';
    }
    bulletPointsDom.innerHTML = bulletPointsOutput;    
}

function nextFoil() {
    currentFoil++;
    if (currentFoil > data.length-1) {
        currentFoil = 0;
    }
    updateContent();
}

function previousFoil() {
    currentFoil--;
    if (currentFoil < 0) {
        currentFoil = data.length-1;
    }
    updateContent();
}