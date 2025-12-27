const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];

let year = new Date().getFullYear();
let date = new Date().getMonth();
let month = months[date];
let day = new Date().getDate();
let time = new Date().getHours();
document.getElementById("year").innerHTML = year;
document.getElementById("month").innerHTML = month;
document.getElementById("day").innerHTML = day;
document.getElementById("time").innerHTML = time;
// "use strict";
// const obj = {};
// Object.defineProperty(obj, "x", {value:0, writable:true, configurable:false});

// console.log(obj.x = 3.14);  

// 'use strict'

// let x = 67
// console.log(x)