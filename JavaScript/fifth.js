// // let age = prompt("What is your age")
// if (age > 20) {
//     console.log("user is ", age)
// }
// let fullname =prompt("what are you name")
// console.log("user name is ", fullname)

let dob = prompt("Enter Your Birthday (YYYY-MM-DD)");

let birthDate = new Date(dob);
let today = new Date();

// calculate age
let aged = today.getFullYear() - birthDate.getFullYear();

// fix the age if birthday not reached yet this year
let month = today.getMonth() - birthDate.getMonth();
if (month < 0 || (month === 0 && today.getDate() < birthDate.getDate())) {
    aged--;
}

if (aged >= 15) {
    console.log("User age:", aged, "✔ Allowed");
} else {
    console.log("Too Young ❌ Not Allowed");
}

// for (let i = 0; i <= 3; i++) {
//     if (i === 2) 
//         continue; // Skip iteration when i is 2
//     console.log(i);
// }


// let day;

// let date = new Date().getDay();
// console.log(date)
// switch (date) {
//   case 0:
//     day = "Sunday";
//     break;
//   case 1:
//     day = "Monday";
//     break;
//   case 2:
//      day = "Tuesday";
//     break;
//   case 3:
//     day = "Wednesday";
//     break;
//   case 4:
//     day = "Thursday";
//     break;
//   case 5:
//     day = "Friday";
//     break;
//   case 6:
//     day = "Saturday";
// }

// let daydate = document.getElementById("day").innerHTML = day;
// console.log(daydate)