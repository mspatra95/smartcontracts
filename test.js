var express = require("express");
var app = express();
var cors = require('cors');

app.options('*', cors());
app.use(cors());


const Zillow = require("node-zillow")

const zillow = new Zillow('X1-ZWz1894yvb2z2j_9prl8')

const parameters = {
    address: "2114 Bigelow Ave N",
    citystatezip: "Seattle, WA",
}
app.get("/getdetails", async (req, res, next) => {
    console.log("***")
    req.setTimeout(50000000)
    
    

    await  zillow.get('GetDeepSearchResults', parameters)
    .then(function(data) {
        // results = data.response.results.result[0].zestimate[0].amount[0]["_"];
        // console.log('zestimate amount');
       // console.log(results);

        req.data = data

       // next();
        }).catch(function (error) {
            console.error(error)
            
          })
 // console.log(req.data);
  res.json(req.data);
});


app.listen(3000, () => {
    console.log("Server running on port 3000");
})
