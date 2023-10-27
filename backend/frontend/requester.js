function requestGet(path) {
    var request = new XMLHttpRequest();
    request.open("GET", "http://localhost:8080/" + path, false);
    var username = sessionStorage.getItem("username");
    var password = sessionStorage.getItem("password");
    var hash = btoa(username + ":" + password);
    request.setRequestHeader("Authorization", "Basic " + hash);
    request.send();
    let jsonData = JSON.parse(request.responseText);
    
    let container = document.getElementById("container");
    
    let table = document.createElement("table");
    table.setAttribute("id", "ctable")
    table.setAttribute("display", "table")
    
    let cols = Object.keys(jsonData[0]);
    
    let thead = document.createElement("thead");
    thead.setAttribute("class", "cthead")
    let tr = document.createElement("tr");
    tr.setAttribute("class", "ctr")
    
    cols.forEach((item) => {
        let th = document.createElement("th");
        th.setAttribute("class", "cth")
        th.innerText = item; 
        tr.appendChild(th); 
    });
    let th = document.createElement("th");
        th.setAttribute("class", "cth")
        th.innerText = "Delete"; 
        tr.appendChild(th); 

    thead.appendChild(tr); 
    table.append(tr) 

    jsonData.forEach((item) => {
        let tr = document.createElement("tr");
        tr.setAttribute("class", "ctr")
        let vals = Object.values(item);
        
        vals.forEach((elem) => {
            let td = document.createElement("td");
            td.setAttribute("class", "ctd")
            td.innerText = elem; 
            tr.appendChild(td); 
        });
        let bt = document.createElement("button");
        bt.setAttribute("id", vals[0]);
        bt.setAttribute("onClick", "requestDelete('"+path+"', this.id);");
        bt.innerHTML = "X";
        let td = document.createElement("td");
        td.setAttribute("class", "ctd");
        bt.setAttribute("class", "tbutton");
        td.appendChild(bt);
        tr.appendChild(td);
        table.appendChild(tr);
    });
    container.appendChild(table);
}

function newGenre() {
    var name = document.getElementById("name").value;
    var genre = new Object();
    genre.name = name;

    var request = new XMLHttpRequest();
    request.open("POST", "http://localhost:8080/genres/", false);
    var username = sessionStorage.getItem("username");
    var password = sessionStorage.getItem("password");
    var hash = btoa(username + ":" + password);
    request.setRequestHeader("Authorization", "Basic " + hash);
    request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    request.send(JSON.stringify(genre));

    alert(request.status);
    location.href = "genres.html";
}

function newInterpreter() {
    var name = document.getElementById("name").value;
    var interpreter = new Object();
    interpreter.name = name;

    var request = new XMLHttpRequest();
    request.open("POST", "http://localhost:8080/interpreters/", false);
    var username = sessionStorage.getItem("username");
    var password = sessionStorage.getItem("password");
    var hash = btoa(username + ":" + password);
    request.setRequestHeader("Authorization", "Basic " + hash);
    request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    request.send(JSON.stringify(interpreter));

    alert(request.status);
    location.href = "interpreters.html";
}

function requestDelete(path, id) {
    if (!confirm("Are you sure you want to delete object with id " + id + "?")) {
        return;
    }
    var request = new XMLHttpRequest();
    request.open("DELETE", "http://localhost:8080/" + path + id, false);
    var username = sessionStorage.getItem("username");
    var password = sessionStorage.getItem("password");
    var hash = btoa(username + ":" + password);
    request.setRequestHeader("Authorization", "Basic " + hash);
    request.send();
    alert(request.status);
    location.href = path.slice(0, -1) + ".html";
}