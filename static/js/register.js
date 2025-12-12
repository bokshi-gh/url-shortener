let form = document.getElementsByTagName("form")[0];
let username = document.getElementsByTagName("input")[0];
let password = document.getElementsByTagName("input")[1];
let button = document.getElementsByTagName("button")[0];

let error = document.getElementsByClassName("error")[0];

button.addEventListener("click", (e) => {
	e.preventDefault();
	
	error.style.display = "none";

	if (username.value === "" || password.value === "") {
		error.textContent = "Input fields can't be empty!";
		error.style.display = "block";
		return;
	}

	let options = {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			username: username.value,
			password: password.value
		})
	}

	registerUser(options)
});

let registerUser = async function(options) {
	let res = await fetch("http://localhost:8080/register", options);

	if (res.ok) {
		console.log("User registered successfully!");

		form.reset();

		window.location.pathname = "pages/login.html";
	}
}


