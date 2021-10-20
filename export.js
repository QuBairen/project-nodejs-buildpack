// run on the wikipedia page

let items = document.querySelectorAll("#mw-content-text > div.mw-parser-output > ul > li")

let laws = [];
let lawcount = 0;
items.forEach(function(item){
  var quote = item.textContent.split(":");
  if(quote.length == 2){
    laws[lawcount++] = { name: quote[0], quote: quote[1] };
  }
})
JSON.stringify(laws);


