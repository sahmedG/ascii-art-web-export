// This function will make the request to the back-end and will modify the html
function validateForm() {
    let x = document.querySelector("#textarea").value;
    //let y = document.querySelector("#Banner").value;
    if (x != "") {
        loadDoc()
    } else {
        alert("You missed entering the text or choosing a banner");
        return false;
    }
}

function Downloadfile () {
    let text = document.querySelector("#art")
    var fileurl = "/export.txt"
    var requestdata = {AsciiArt: text.innerHTML};
    fetch(fileurl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestdata)
      })
      .then(response => response.blob())
      .then(blob => {
      // Create a temporary anchor element
      var link = document.createElement('a');
      link.href = URL.createObjectURL(blob);
      link.download = 'export.txt'; // Replace with the desired file name and extension
      link.click();

      // Clean up the object URL after the download starts
      setTimeout(() => URL.revokeObjectURL(link.href), 100);
    })
    .catch(error => {
      console.log('Error:', error);
    });
}

function loadDoc() {
    const xhttp = new XMLHttpRequest();
    let text = document.querySelector(".input")
    let banner = document.querySelector('input[name="Banner"]:checked');
    let Color = document.querySelector(".newColor")
    invertColor(Color.value)
    xhttp.onload = function () {
        if (this.status === 500){
            //alert("internal 500 error")
            window.open("http://localhost:8080/internalerror.html","_self")
        } else if (this.status === 400){
            window.open("http://localhost:8080/badrequest.html","_self")
        } else {
        var parsedData = JSON.parse(this.responseText);                 // Parse JSON
        //var formattedText = parsedData.Result.replace(/\n/g, "<br>");   // Replace the lines with <br>
        document.getElementById("art").innerHTML = parsedData.Result;
        document.getElementById("art").value = parsedData.Result;
        document.getElementById("art").style.color = parsedData.ApplyColor;
        document.querySelector("#btndown").disabled = false;
        }
    }
    xhttp.open("POST", "/ascii-art");
    xhttp.setRequestHeader("Content-Type", "application/json; charset=utf-8");
    const body = {
        Text: text.value,
        Banner: banner.value,
        Newcolor: Color.value,
    };
    xhttp.send(JSON.stringify(body));
}

function invertColor(hex) {
    if (hex.indexOf('#') === 0) {
        hex = hex.slice(1);
    }
    // convert 3-digit hex to 6-digits.
    if (hex.length === 3) {
        hex = hex[0] + hex[0] + hex[1] + hex[1] + hex[2] + hex[2];
    }
    if (hex.length !== 6) {
        throw new Error('Invalid HEX color.');
    }
    // invert color components
    var r = (parseInt(hex.slice(0, 2), 16)),
        g = (parseInt(hex.slice(2, 4), 16)),
        b = (parseInt(hex.slice(4, 6), 16));
    // pad each with zeros and return
    //r = padZero(r)
    if (r <= 100 && b <= 100 && g <= 100){
        document.getElementById("result").style.backgroundColor = "#ffffff"
    } else {
        document.getElementById("result").style.backgroundColor = "#302f2f"
    }



}

