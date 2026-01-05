let numbers = new Set()

let alpha = "alpha"
let beta = "beta"
let gamma = "gamma"
let lambda = "lambda"

numbers.add(alpha)
numbers.add(alpha)
numbers.add(beta)
numbers.add(gamma)
numbers.add(lambda)

console.log(numbers)
console.log(typeof(numbers))
document.getElementById("set").innerHTML = numbers.size