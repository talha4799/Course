let txt = "Please Always keep Your Phone Muted and Silenced";

let z = txt.matchAll(/ed/g);   // matchAll needs a RegExp
let y = Array.from(z);         // correct Array.from usage

console.log(y);

// z = txt.indexOf("e", 17)
// let txt = "    Hello World      t  "
// z = txt.trimStart()
// z = txt.charAt(8)
// z = txt.at(8)
// z = txt[8]
// z = txt.slice( 2, 8)
// z = txt.isWellFormed()

// let heading = "Hello World"
// head = `
//<ul style="list-style-type: circle;">
//<li>
//Hello
//</li>
//</ul
//<h1>${heading}</h1>`
//document.getElementById("heading").innerHTML = head
// let price = 20
// discount = 3

// total = `${(price * (.25 + discount)).toFixed(4)}`
// console.log(total)

// let fname = "John"
// let lname = "Doe"
// let fullname = `My name is ${fname} ${lname}`

// console.log(fullname)