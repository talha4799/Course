const person = {
        firstname: "Talha",
        lastname: "Majeed",
        age: "230",
        eyeColor: "Blue"
    };
    console.log(person)
    document.getElementById("before").innerHTML =
        JSON.stringify(person, null,2)
        
// console.log(person)
// delete person.eyeColor
// console.log(person)

// function person(first, last, age, eye) {
//     this.firstname = first,
//     this.lastname = last,
//     this.age = age,
//     this.eyeColor = eye
// }

// person.prototype.country = "Pakistan"
// let talha = new person("Talha", "Majeed",  "120",  "Blue" )
// console.log(talha)
// console.log(talha.country)
// car = {
//     carName: "Lexus",
//     carModel: 761,
//     carPrice: "120,000$",
//     fullCar: function() {
//     return this.carName + " " + this.carModel + " " + this.carPrice
//     }
//     }

// console.log("Car Name:",car.carName,"|","Car Model:", car.carModel,"|","Car Price:", car.carPrice, car.fullCar())

// let sub = a => a
// const multiply = function(a, b, c){  
//     return (a+b)*c  
// }

// z = sub(20)

// function multiply(a,b) {
//     return a * b
// }

// let now = multiply(5, 4)
// console.log(z)