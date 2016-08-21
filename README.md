# go-gisbn
It's a standard GOlang library. It fetches book's information by ISBN number based on Google Book API.

Installation
-------------
	go get github.com/eftakhairul/go-gisbn/gisbn

Project Usage
-------------
The API usage is very simple. Just import the go-gisbn package

	import (
		"github.com/eftakhairul/go-gisbn/gisbn"
	)

	//Make instace and call methods
	g := gisbn.New("9780262033848", "AIzaSyDKepjfaVBRcgsnPALw5s2UNyfOk-1FHUU", "ca")
	g.Fetch()
	fmt.Println(g.Title())   //Introduction to Algorithms


Example
-------------	
[example](https://github.com/eftakhairul/go-gisbn/tree/master/example)

### Development

Want to contribute? Great!

1. Fork it ( https://github.com/eftakhairul/go-gisbn/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request


Author
-----------
[Eftakhairul Islam] - eftakhairul [at] gmail [dot] com
[Eftakhairul Islam]:https://eftakhairul.com/


License
-------
MIT


Contribute
----------

This may have bugs and lack of many features. If you want to contribute on this project, you are more than welcome. Please fork the repository from [Github](https://github.com/eftakhairul/go-gisbn).

Please [donate](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=eftakhairul%40gmail%2ecom&lc=CA&item_name=Eftakhairul%20world&item_number=web_product&currency_code=CAD&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted) me to enhance this project.