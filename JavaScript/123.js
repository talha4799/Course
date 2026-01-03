
function validateForm(event) {
    event.preventDefault()
    let email = document.getElementById("email").value;
    let name = document.getElementById("fname").value;
    let phone = document.getElementById("phone").value;
    let password = document.getElementById("password").value;
    let image = document.getElementById("image").value;

    // let valid = true;

    // clear errors
    document.getElementById("emailError").textContent = "";
    document.getElementById("nameError").textContent = "";
    document.getElementById("phoneError").textContent = "";
    document.getElementById("passError").textContent = "";
    document.getElementById("imageError").textContent = "";


    if (email.length < 7) {
        document.getElementById("emailError").textContent = 
        "Must be More than 7 characters";
        valid = false;
    } else {
    console.log(email)

    }
    // Name validation (min 5)
    if (name.length < 5) {
        document.getElementById("nameError").textContent =
            "Name must be at least 5 characters";
        valid = false;        
    } else {
        console.log(name)
    }

    // Phone validation (exact 11 digits)
    if (!/^\d{11}$/.test(phone)) {
        document.getElementById("phoneError").textContent =
            "Phone number must be exactly 11 digits";
        valid = false;
    } else {
    console.log(phone)

    }

    // Password validation (min 6)
    if (password.length < 6) {
        document.getElementById("passError").textContent =
            "Password must be at least 6 characters";
        valid = false;
    } else {
    console.log(password)
    }
    return valid
}
function validateForm(event) {
    event.preventDefault(); // page reload stop

    let name = document.getElementById("fname").value;
    let email = document.getElementById("email").value;
    let phone = document.getElementById("phone").value;
    let password = document.getElementById("password").value;

    // simple validation
    if (name === "" || email === "" || phone === "" || password === "") {
        alert("Please fill all fields");
        return false;
    }

    // output div
    let output = document.getElementById("output");

    // new data add (append)
    output.innerHTML += `
        <hr>
        <p><b>Name:</b> ${name}</p>
        <p><b>Email:</b> ${email}</p>
        <p><b>Phone:</b> ${phone}</p>
        <p><b>Password:</b> ${password}</p>
    `;

    // form clear after submit
    document.querySelector("form").reset();

    return false;
}

