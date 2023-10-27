function requestGet(path, delButton) {
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
    if (delButton) {
        let th = document.createElement("th");
        th.setAttribute("class", "cth")
        th.innerText = "Delete"; 
        tr.appendChild(th); 
    }

    thead.appendChild(tr); 
    table.append(tr) 

    jsonData.forEach((item) => {
        let tr = document.createElement("tr");
        tr.setAttribute("class", "ctr")
        let vals = Object.values(item);
        
        vals.forEach((elem) => {
            let td = document.createElement("td");
            td.setAttribute("class", "ctd");
            if (typeof(elem) == "object") {
                let values = Object.values(elem);
                if (typeof(values)[0] == "object") {
                    let stringArray = "";
                    values.forEach((item) => {
                        stringArray += item.Name + ", ";
                    });
                    stringArray = stringArray.slice(0, -2);
                    td.innerText = stringArray;
                }
                else {
                    td.innerText = Object.values(elem)[1];
                }
            }
            else {
                td.innerText = elem; 
            }
            tr.appendChild(td); 
        });
        if (delButton) {
            let td = document.createElement("td");
            let bt = document.createElement("button");
            bt.setAttribute("id", vals[0]);
            bt.setAttribute("onClick", "requestDelete('"+path+"', this.id);");
            bt.innerHTML = "X";
            bt.setAttribute("class", "tbutton");
            td.appendChild(bt);
            td.setAttribute("class", "ctd");
            tr.appendChild(td);
        }
        table.appendChild(tr);
    });
    container.appendChild(table);
}

function requestGetById(path, id) {
    var request = new XMLHttpRequest();
    request.open("GET", "http://localhost:8080/" + path + "?id=" + id, false);
    var username = sessionStorage.getItem("username");
    var password = sessionStorage.getItem("password");
    var hash = btoa(username + ":" + password);
    request.setRequestHeader("Authorization", "Basic " + hash);
    request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    request.send();

    return request.responseText.slice(1, -1);
}

function requestCreate(path, object) {
    var request = new XMLHttpRequest();
    request.open("POST", "http://localhost:8080/" + path, false);
    var username = sessionStorage.getItem("username");
    var password = sessionStorage.getItem("password");
    var hash = btoa(username + ":" + password);
    request.setRequestHeader("Authorization", "Basic " + hash);
    request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    request.send(JSON.stringify(object));

    alert(request.status);
    location.href = path.slice(0, -1) + ".html";
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

function newGenre() {
    var name = document.getElementById("name").value;
    var genre = new Object();
    genre.name = name;

    return genre;
}

function newInterpreter() {
    var name = document.getElementById("name").value;

    var interpreter = new Object();
    interpreter.name = name;

    return interpreter;
}

function newSong() {
    let name = document.getElementById("name").value;
    let url = document.getElementById("url").value;
    let genreId = document.getElementById("genreId").value;
    let interpretersIds = document.getElementsByClassName("manyRelation");
    let interpreters = {};
    for (i = 0; i < interpretersIds.length; i++) {
        interpreters[i] = (JSON.parse(requestGetById("interpreters/", interpretersIds.item(i).value)));
    }

    let song = new Object();
    song.name = name;
    song.url = url;
    song.genre = JSON.parse(requestGetById("genres/", genreId));
    song.interpreters = Object.values(interpreters);

    return song;
}

function newItem() {
    let name = document.getElementById("name").value;
    let description = document.getElementById("description").value;

    let item = new Object();
    item.name = name;
    item.description = description;

    return item;
}

function newTopic() {
    let topic = document.getElementById("topic").value;
    
    let topicobj = new Object();
    topicobj.topic = topic;

    return topicobj;
}

function newSubtopic() {
    let subtopic = document.getElementById("subtopic").value;
    let topicId = document.getElementById("topicId").value;

    let subtopicObj = new Object();
    subtopicObj.subtopic = subtopic;
    subtopicObj.topic = JSON.parse(requestGetById("topics/", topicId));

    return subtopicObj;
}