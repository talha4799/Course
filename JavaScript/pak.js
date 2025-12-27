function verifyVote() {
  const s1 = document.getElementById("s1").value;
  const s2 = document.getElementById("s2").value;
  const msg = document.getElementById("msg");

  msg.innerHTML = "";
  msg.className = "msg";

  // Empty check
  if (!s1 || !s2) {
    msg.className = "msg error";
    msg.innerHTML = "⚠️ دونوں صوبے منتخب کرنا ضروری ہے!";
    return;
  }

  // VALID combination → Punjab + KPK
  if (s1 === "Punjab" && s2 === "KPK") {
    msg.className = "msg success";
    msg.innerHTML = "✔️ Vote Verified Successfully!";
    return;
  }

  // Any other combination → Not allowed
  msg.className = "msg error";
  msg.innerHTML = "❌ یہ صوبائی کمبینیشن الاؤ نہیں ہے!";
}