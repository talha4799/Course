// age = 12
// let txt1 = "minor"
// let txt2 = "old age"
// let txt = (age <= 30) ? txt1 : txt2
// console.log(txt)

// function namenew() {
//    let name = document.getElementById("fullname").value.trim()

//    if (!name) {
//        let onerror = "Please Enter Your Name and details"
//        document.getElementById("error").textContent = onerror
//        // alert(onerror)
//    } else {
//        document.getElementById("error").textContent = ""
//    }

//     alert(name)
//    document.getElementById("user-info").textContent = name

// }

// function showName() {
//     const getName = document.getElementById("fullname").value.trim();

//     if (!getName) {
//         alert("Please enter your name.");
//         return;
//     }

//     alert(getName); // show in alert
//     document.getElementById("user_info").textContent = getName; // show in h1
function verifyVote() {
  const s1 = document.getElementById("s1").value;
  const s2 = document.getElementById("s2").value;
  const msg = document.getElementById("msg");

  // Reset message box
  msg.innerHTML = "";
  msg.className = "msg"; // default class

  // Error #1 — No selection
  if (!s1 || !s2) {
    msg.className = "msg error";
    msg.innerHTML = "⚠️ دونوں صوبے منتخب کرنا لازمی ہے!";
    return;
  }

  // Error #2 — Both same
  if (s1 === s2) {
    msg.className = "msg error";
    msg.innerHTML = "⚠️ دونوں صوبے مختلف ہونے چاہئیں!";
    return;
  }


  // Success
  msg.className = "msg success";
  msg.innerHTML = "✔️ ووٹ کامیابی سے ویریفائی ہو گیا!";
}

