var current_rotation = 0;
var captchaImage = document.getElementById("captcha-image");
var spinner = document.getElementById("spinner");
var confirmation = document.getElementById("confirmation");
var buttonLeft = document.getElementById("rotate-right");
var buttonRight = document.getElementById("rotate-left");
var buttonSubmit = document.getElementById("submit");
var groupID = document.getElementById("group-id");
var captchaCount = document.getElementById("captcha-count");

function rotate(direction) {
    if (direction == "right") {
        current_rotation += 40;
        document.querySelector("#captcha-image").style.transform = 'rotate(' + current_rotation + 'deg)';
        document.getElementById("rotation-degree").value = current_rotation;
        console.log("right")
    } else {
        current_rotation -= 40;
        document.querySelector("#captcha-image").style.transform = 'rotate(' + current_rotation + 'deg)';
        document.getElementById("rotation-degree").value = current_rotation;
        console.log("left")
    }
}

function resetRotation(){
    current_rotation = 0;
    document.querySelector("#captcha-image").style.transform = 'rotate(' + current_rotation + 'deg)';
    document.getElementById("rotation-degree").value = current_rotation;
}

function showSpinner(yes) {
    if (yes){
        captchaImage.style.display = "none";
        confirmation.style.display = "none";
        spinner.style.display = "block";
        buttonLeft.disabled = true;
        buttonRight.disabled = true;
        buttonSubmit.disabled = true;
    } else {
        captchaImage.style.display = "block";
        confirmation.style.display = "none";
        spinner.style.display = "none";
        buttonLeft.disabled = false;
        buttonRight.disabled = false;
        buttonSubmit.disabled = false;
    }
}

function showConfirmation() {
    captchaImage.style.display = "none";
    spinner.style.display = "none";
    confirmation.style.display = "block";
    buttonLeft.disabled = true;
    buttonRight.disabled = true;
    buttonSubmit.disabled = true;

    groupID.classList.remove("badge-secondary");
    groupID.classList.add("badge-success");

    captchaCount.classList.remove("badge-secondary");
    captchaCount.classList.add("badge-success");
}