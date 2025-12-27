function details (event) {
    event.preventDefault()
    let fname = document.getElementById("fname").value
    let email = document.getElementById("email").value
    let number = document.getElementById("phone").value
    let password = document.getElementById("password").value
    let salary = document.getElementById("salary").value
    
    document.getElementById("name-output").textContent = fname
    document.getElementById("email-output").textContent = email
    document.getElementById("phone-output").textContent = number
    document.getElementById("password-output").textContent = password
    document.getElementById("salary-output").textContent = salary
}

// let inputs = document.querySelectorAll("input");
// let output = document.getElementById("output");

// inputs.forEach(function (input) {
//     input.addEventListener("input", function () {
//         output.innerHTML = `
//             Name: ${fname.value}<br>
//             Email: ${email.value}<br>
//             Phone: ${phone.value}<br>
//             Password: ${password.value}<br>
//             Salary: ${salary.value}
//         `;
//     });
// });

// function details(email, fullname, phoneNumber, password, username,salary, isverified) {
//     this.email = email;
//     this.fullname = fullname;
//     this.phoneNumber = phoneNumber;
//     this.password = password;
//     this.username = username;
//     this.salary = salary;
//     this.isverified = isverified;
// }

// // let talha = new details("talha@gmail.com", "Talha", +923434234564, 12345678, "talha", 500000, true) 
// let talha = new details(username, "talha", password, 3456)
// document.getElementById("email").innerHTML = talha.email
// document.getElementById("name").innerHTML = talha.fullname
// document.getElementById("phone").innerHTML = talha.phoneNumber
// document.getElementById("username").innerHTML = talha.username
// document.getElementById("password").innerHTML = talha.password
// document.getElementById("salary").innerHTML = talha.salary
// document.getElementById("isverified").innerHTML = talha.isverified
// // txt = "@"
// z = txt.padStart(7, "waleed568")
// console.log(z)
// console.log(typeof(z))


// txt1 = "hello";
// txt2 = "world"
// z = txt1.localeCompare(txt1)
// console.log(z)
// let text = "true world"
// let pr = text.endsWith("true");
// console.log(pr)
// console.log(typeof(pr))