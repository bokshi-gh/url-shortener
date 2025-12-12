let form = document.querySelector("form");
let username = document.getElementById("username");
let password = document.getElementById("password");
let button = document.querySelector("button");
let error = document.querySelector(".error");

button.addEventListener("click", async (e) => {
    e.preventDefault();
    error.style.display = "none";

    if (username.value === "" || password.value === "") {
        error.textContent = "Input fields can't be empty!";
        error.style.display = "block";
        return;
    }

    let options = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username: username.value, password: password.value })
    }

    try {
        let res = await fetch("http://localhost:8080/register", options);
        if (res.ok) {
            console.log("User registered successfully!");
            form.reset();
            window.location.href = "/pages/login.html";
        } else {
            let msg = await res.text();
            error.textContent = msg;
            error.style.display = "block";
        }
    } catch (err) {
        error.textContent = "Network error!";
        error.style.display = "block";
    }
});

