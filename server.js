var express    = require('express'),
    htmlEscape = require('html-escape'),
    bodyParser = require('body-parser');

var app = express();

var serveHelloWorld = function(req, res){
  res.send('Hello world!');
}

app.use(/\/sloths$/, function(req, res){
  res.send('Sloths rule!');
});

app.use('/images', express.static(__dirname+'/public/images'));

app.use('/tea/:flavor', function(req, res){
  var tea = req.params.flavor + ' tea';
  var html = '<img src="/images/sloth.jpg" /><br />'+
             '<h2>I could use some ' + htmlEscape(tea) + '!</h2>';
  res.send(html);
});

app.use(/\/(coffee)+/, function(req, res){
  var html = '<img src="/images/lemur.jpg" /><br />'+
             '<h2>Lemurs = sloths that had too much coffee!</h2>';
  res.send(html);
});

app.get('/coffee-shop', function(req, res){
  res.send('<form action="/order" method="POST">'+
             'Your name <input type="text" name="name" /><br/ >'+
             'Your favorite beverage <input type="text" name="beverage" /><br/ >'+
             '<input type="submit" value="Submit"/>'+
           '</form>');
});

app.post('/order', bodyParser.urlencoded({extended:true}), function(req, res){
  var name     = htmlEscape(req.body.name),
      beverage = htmlEscape(req.body.beverage);
  res.send('<h1>One '+ beverage + ' coming right up ' + name + '!</h1>');
});

app.use('/', serveHelloWorld);
app.listen(1123);
