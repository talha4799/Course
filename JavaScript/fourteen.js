let computers = [
  "Mac",
  "Windows",
  ["Linux", "Ubuntu"],
  "Chrome OS",
  "Android OS",
  "iOS",
  "HarmonyOS"
];

// show array on page load
function showComputers() {
  let comp = computers.flat();
  document.getElementById("comp").innerHTML = comp.join(", ");
}
showComputers();

function addComputer() {
  let input = document.getElementById("computerInput").value.trim();
  let message = document.getElementById("message");

  // ❌ error: empty input
  if (input === "") {
    message.innerHTML = "❌ Please enter a computer / OS name";
    message.style.color = "red";
    return;
  }

  // ❌ error: already exists
  let allComputers = computers.flat().map(c => c.toLowerCase());
  if (allComputers.includes(input.toLowerCase())) {
    message.innerHTML = "⚠️ This OS already exists";
    message.style.color = "orange";
    return;
  }

  // ✅ success
  computers.push(input);
  showComputers();

  message.innerHTML = "✅ Computer added successfully";
  message.style.color = "green";

  document.getElementById("computerInput").value = "";
}


// let cars = ["Saab", "Volvo", "BMW"];
// let carname = cars.at(1)
// console.log(carname)
// let array3 = ["1", "2", "3"]
// array3.push("4")

// console.log(array3)
// array3.pop(2)
// console.log(array3)

// array3[0] = "0"
// console.log(array3.length)

// let names = new Array("Ronnie", "John", "Doe", "Smith");
// console.log(names);

// let array2 = [];
// array2[3] = 678
// array2[2] ="Hello"
// array2[1] = "Ronnie"
// array2[0] = true
// console.log(array2)
// let array = ["One", "Two", "Three", "So On"]
// console.log(array)