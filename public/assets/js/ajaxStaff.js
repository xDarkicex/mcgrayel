document.addEventListener("DOMContentLoaded", function(e){
  var xhr = new XMLHttpRequest();
  xhr.onload = function() {
     if (xhr.status != 200) {
       return
     }
     var data = JSON.parse(xhr.responseText)
     var len = data.length
     for (i = 0; i < len; i++) {
      var container = document.querySelectorAll('.col-lg')[0]
      var info = document.createElement('div')
      if (data[i].Name) {
        var title = document.createElement('h4')
        title.className = "card-title"
        title.innerHTML = data[i].Name
        title.style = "font-weight: bold;"
        info.appendChild(title)
      }
      if (data[i].Title) {
        var jobTitle = document.createElement('p')
        jobTitle.innerHTML = data[i].Title
        info.appendChild(jobTitle)
      }
      if (data[i].Phone) {
        var Phone = document.createElement('p')
        Phone.innerHTML = data[i].Phone
        info.appendChild(Phone)
      }
      if (data[i].Email) {
        var Email = document.createElement('p')
        Email.innerHTML = data[i].Email
        info.appendChild(Email)
      }
       info.className = "card row center"
      container.appendChild(info)
     }
 }
 xhr.open("GET", "/api/staff");
 xhr.send()
})

