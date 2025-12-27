// let x = 5
// z = x**3
// x--
// let x= 10
// x**=2
// x= x+10
// let text = "Hello"
// text+= "Wolrd"
// text = "hello World"
// x= true
// y= x !=10

// text = "12345"
// mathmin = Math.min(...text)
// mathmax = Math.max(...text)

// x = 9
// y = 20
// let z = x||=y
// console.log(z)
// console.log(mathmin)
// console.log(mathmax)
function checkDate() {

    let dob = document.getElementById("birthday").value;

    if (!dob) {
        alert("Please select a date");
        return;
    }

    // User date
    const userDate = new Date(dob);

    // Limit dates
    const minDate = new Date("2024-11-12"); // Too Young
    const maxDate = new Date("2026-11-05"); // Good to go

    if (userDate < minDate) {
        alert("Too Young");
    } 
    else if (userDate <= maxDate) {
        alert("Good To Go");
    } 
    else {
        alert("Date Not Allowed");
    }
}
